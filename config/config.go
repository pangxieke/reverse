package config

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type OpenApiConfig struct {
	Host string
	Port uint
}

type MgoConfig struct {
	Urls string
	DB   string
}

type ServerConfig struct {
	Port    uint
	LogFile string
}

var (
	Env     string
	OpenApi OpenApiConfig
	Mgo     MgoConfig
	Server  ServerConfig
)

func Init(configPaths ...string) (err error) {
	if err := setup(configPaths...); err != nil {
		return err
	}
	if err := initOpenApi(); err != nil {
		return err
	}
	if err := initMgo(); err != nil {
		return err
	}
	if err := initServer(); err != nil {
		return err
	}
	return
}

func setup(paths ...string) (err error) {
	Env = os.Getenv("GO_ENV")
	if "" == Env {
		Env = "dev"
	}
	godotenv.Load(".env." + Env)
	godotenv.Load()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err = viper.ReadInConfig()
	if err != nil {
		logs.Error("Failed to read config file (but environment config still affected), err = %+v\n", err)
		err = nil
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return
}

func initOpenApi() (err error) {
	OpenApi.Host = viper.GetString("openapi.host")
	OpenApi.Port = viper.GetUint("openapi.port")

	if OpenApi.Host == "" {
		return errors.New("openapi.host should not be empty")
	}
	if OpenApi.Port == 0 {
		return errors.New("openapi.port should not be empty")
	}
	return
}

func initMgo() (err error) {
	Mgo.Urls = viper.GetString("mgo.urls")
	Mgo.DB = viper.GetString("mgo.db")

	if Mgo.Urls == "" {
		return errors.New("mgo.host should not be empty")
	}
	if Mgo.DB == "" {
		return errors.New("mgo.db should not be empty")
	}
	return
}

func initServer() (err error) {
	Server.LogFile = viper.GetString("server.log")
	Server.Port = viper.GetUint("server.port")

	if Server.LogFile == "" {
		return errors.New("server.log should not be empty")
	}
	//if Server.Port == 0 {
	//	return errors.New("server.port should not be empty")
	//}

	return
}
