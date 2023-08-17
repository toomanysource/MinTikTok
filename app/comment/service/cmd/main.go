package main

import (
	"Atreus/app/comment/service/internal/conf"
	"Atreus/pkg/logX"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"os"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "comment"
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", CheckConfigExist(), "config path, eg: -conf config.yaml")
}

func CheckConfigExist() (path string) {
	path = "../configs/config.dev.yaml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "../configs/config.yaml"
	}
	return
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	l := logX.NewDefaultLogger()
	f, err := l.FilePath("../../../../logs/comment/" + l.SetTimeFileName("", false))
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(f, os.Stdout)
	l.SetOutput(writer)
	l.SetLevel(log.LevelDebug)
	logger := log.With(l,
		"service", Name,
		"caller", log.DefaultCaller,
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	app, cleanup, err := wireApp(bc.Server, bc.Client, bc.Data, bc.Jwt, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
