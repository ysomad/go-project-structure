package main

import (
	"flag"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ysomad/go-project-structure/internal/app/server"
	"github.com/ysomad/go-project-structure/internal/config"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

func main() {
	var (
		confpath string
		conf     config.Server
		migrate  bool
	)

	flag.StringVar(&confpath, "config", "./configs/server/local.toml", "config path")
	flag.StringVar(&conf.Metadata.Version, "version", "local-0", "service version")
	flag.StringVar(&conf.Metadata.InstanceID, "instance", "0", "service instance id")
	flag.Int64Var(&conf.Metadata.BuildTime, "buildtime", time.Now().Unix(), "build timestamp")
	flag.BoolVar(&migrate, "migrate", false, "run up migrations")
	flag.Parse()

	if err := cleanenv.ReadConfig(confpath, &conf); err != nil {
		slogx.Fatal("read config: " + err.Error())
	}

	if err := server.Run(conf, migrate); err != nil {
		slogx.Fatal(err.Error())
	}
}
