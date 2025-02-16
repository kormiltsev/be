package run

import (
	"context"
	"fmt"
	standardlog "log"
	"log/slog"
	"net"
	"os"
	"strings"
	"time"

	cli "github.com/urfave/cli/v2"
	tomb "gopkg.in/tomb.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/kormiltsev/be/config"
	"github.com/kormiltsev/be/internal/api"
)

var Command = &cli.Command{
	Name:  "run",
	Usage: "Run server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "socket",
			Usage:   "REST API `socket` either as 'tcp://<address>:<port>' or 'unix://<path>' string",
			EnvVars: []string{"RENAMEME_SOCKET_REST_API"},
			Value:   "tcp://127.0.0.1:8080",
		},
		&cli.StringFlag{
			Name:    "db-sqlite-file",
			Usage:   "SQLite database file address",
			EnvVars: []string{"RENAMEME_DB_SQLITE_FILE"},
			Value:   "/tmp/renameme.db",
		},
	},
	OnUsageError: func(c *cli.Context, err error, isSubCommand bool) error {
		return cli.ShowCommandHelp(c, "run")
	},
	Action: func(c *cli.Context) error {

		var programLevel = new(slog.LevelVar)
		// lifetime changing:
		if config.DebugMode {
			programLevel.Set(slog.LevelDebug)
		}
		slog.Info("logger intialiaed", slog.String("LEVEL", programLevel.String()))
		// log = log.WithGroup("application")

		loghandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     programLevel,
		})
		slogger := slog.New(loghandler)
		slog.SetDefault(slogger)

		t, _ := tomb.WithContext(context.Background())

		netListener, err := makeNetListener(c.String("socket"))
		if err != nil {
			return err
		}

		dbfileaddress := c.String("db-sqlite-file")

		db, err := gorm.Open(sqlite.Open(dbfileaddress), &gorm.Config{
			Logger: gormlogger.New(
				standardlog.New(os.Stderr, "\r\n", standardlog.LstdFlags), // io writer
				gormlogger.Config{
					SlowThreshold:             time.Second,     // Slow SQL threshold
					LogLevel:                  gormlogger.Warn, // Log level // make it Silent,Error, Warn or Info
					IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
					ParameterizedQueries:      true,            // Don't include params in the SQL log
					Colorful:                  false,           // Disable color
				},
			),
		})
		if err != nil {
			slogger.Error("connect to database", slog.String("dbfileaddress", dbfileaddress), slog.String("error", err.Error()))
			return fmt.Errorf("connect to database with connection string %q: %v", dbfileaddress, err)
		}

		defer func(db *gorm.DB) {
			sqlDB, err := db.DB()
			if err != nil {
				slogger.Error("Failed to get database instance", "error:", err.Error())
			}
			defer sqlDB.Close()
		}(db)

		api.HandleHTTPServer(t, slogger, netListener, db)

		return t.Wait()
	},
}

func makeNetListener(socket string) (netListener net.Listener, err error) {
	if strings.HasPrefix(socket, "unix://") {
		f := strings.TrimPrefix(socket, "unix://")
		if _, err = os.Stat(f); err == nil {
			err = os.Remove(f)
			if err != nil {
				return nil, err
			}
		}
		if netListener, err = net.Listen("unix", f); err == nil {
			err = os.Chmod(f, 0600)
		}
	} else {
		socket = strings.TrimPrefix(socket, "tcp://")
		netListener, err = net.Listen("tcp", socket)
	}

	return
}
