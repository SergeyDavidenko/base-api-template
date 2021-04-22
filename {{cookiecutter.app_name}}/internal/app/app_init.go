package app

import (
	"os/signal"
	"syscall"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/api"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/models"
	{% if cookiecutter.use_postgresql == "y" %}
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/storage"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/storage/postgresql"
	{% endif %}
	{% if cookiecutter.use_redis == "y" %}
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/storage/redis"
	{% endif %}
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// ServerInit init vars, config and log level
func ServerInit(configPath string) {
	var err error
	signal.Notify(api.Quit, syscall.SIGINT, syscall.SIGTERM)
	config.Conf, err = config.LoadConf(configPath)
	if err != nil {
		log.Error(err)
	}
	level, err := log.ParseLevel(config.Conf.Log.Level)
	if err != nil {
		log.Error("Cannot parse log level")
		log.SetLevel(log.InfoLevel)
	}
	log.Debug("Set log level ", level)
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	// Support metrics
	metrics := models.NewMetrics()
	prometheus.MustRegister(metrics)
	{% if cookiecutter.use_postgresql == "y" %}
	storage.StorageDB = postgresql.New()
	errInitDB := storage.StorageDB.Init()
	if errInitDB != nil {
		log.Fatal(errInitDB)
	}
	{% endif %}
	{% if cookiecutter.use_redis == "y" %}
	storage.CacheRedis = redis.New()
	errInitCache := storage.CacheRedis.Init()
	if errInitCache != nil {
		log.Fatal(errInitCache)
	}
	{% endif %}
}