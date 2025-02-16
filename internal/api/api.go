package api

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	tomb "gopkg.in/tomb.v2"
	"gorm.io/gorm"

	"github.com/kormiltsev/be/config"
	itmsrv "github.com/kormiltsev/be/internal/api/gen/http/items/server"
	versrv "github.com/kormiltsev/be/internal/api/gen/http/version/server"
	itemsEnd "github.com/kormiltsev/be/internal/api/gen/items"
	versionEnd "github.com/kormiltsev/be/internal/api/gen/version"
	"github.com/kormiltsev/be/internal/controllers/items"
	"github.com/kormiltsev/be/internal/controllers/version"
	"github.com/kormiltsev/be/internal/dal"

	itmSvc "github.com/kormiltsev/be/internal/service/items"
)

func HandleHTTPServer(t *tomb.Tomb, logger *slog.Logger, netListener net.Listener, db *gorm.DB) {

	d := dal.New(db)
	if err := d.Migrate(); err != nil {
		logger.Error("Database migration error, abandoned", slog.String("error", err.Error()))
		return
	}

	// Services
	itemsService := itmSvc.New(t, logger, d) // + event server?

	logger = logger.WithGroup("HTTP")
	logger.Info("start http server", "addr", netListener.Addr())

	// adapter := middleware.NewLogger(log.New(os.Stderr, "[ HTTP ] ", log.Ltime))
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	// init controllers
	versionController := version.NewController()
	itemsController := items.NewController(itemsService, logger.With(slog.String("service", "item")))

	// init endpoints
	versionEndpoints := versionEnd.NewEndpoints(versionController)
	itemEndpoints := itemsEnd.NewEndpoints(itemsController)

	eh := errorHandler(logger)
	// init servers
	itemSrv := itmsrv.New(itemEndpoints, mux, dec, enc, eh, nil)
	versionSrv := versrv.New(versionEndpoints, mux, dec, enc, eh, nil)

	servers := goahttp.Servers{
		itemSrv,
		versionSrv,
	}

	if config.DebugMode {
		servers.Use(httpmdlwr.Debug(mux, os.Stdout))
	}

	// mount servers
	itmsrv.Mount(mux, itemSrv)
	versrv.Mount(mux, versionSrv)

	// mux = auth.Middleware()(mux)

	srv := &http.Server{Handler: mux}

	t.Go(func() error {
		return srv.Serve(netListener)
	})

	// shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	t.Go(func() error {
		select {
		case <-sig:
			logger.Info("signal received, shutdown server")
		case <-t.Dying():
			logger.Info("shutdown server")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
		return nil
	})
}

func errorHandler(logger *slog.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Error("[%s] ERROR: %s", id, err.Error())
	}
}
