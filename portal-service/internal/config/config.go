package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const (
	CONFIG_FILE_TYPE = "json"
	CONFIG_FILE_NAME = "config.json"
	LOG_JSON_FORMAT = "json"
	LOG_CONSOLE_FORMAT = "console"
	LOG_COMPOSE_FORMAT = "compose"
)

type myFormatter struct {
	log.TextFormatter
}

type Envs struct {
	AppEnv           string `mapstructure:"APP_ENV"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	FrontPort        string `mapstructure:"FRONT_PORT"`
	AuthenticatePort string `mapstructure:"AUTHENTICATE_PORT"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
	LogFormat string  `mapstructure:"LOG_FORMAT"`
}

var Env Envs

func Initialize() error {
	Env = Envs{}
	// viper.SetConfigFile(".env")
	viper.SetConfigFile(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_FILE_TYPE)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file json : ", err)
		return err
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return err
	}

	if Env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	//log.Printf("setting Broker-port: %s \t Authenticate-port: %s", env.BrokerPort, env.AuthenticatePort)

	configureLogger()

	return nil
}

func configureLogger() {

	logLevel, err := log.ParseLevel(Env.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.Printf("Log level: %s", logLevel.String())
	log.SetLevel(logLevel)
	log.Infof("LOG FORMAT: %s", Env.LogFormat)

	switch Env.LogFormat {
	case LOG_JSON_FORMAT:
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat:        "02-01-2006 15:04:05",
		})
	case LOG_CONSOLE_FORMAT:
		log.SetFormatter(
			&myFormatter{log.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "02-01-2006 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: false,
			}})
	case LOG_COMPOSE_FORMAT:
		log.SetFormatter(
			&myFormatter{
				log.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "02-01-2006 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: false,

			}})
	default:
		log.SetFormatter(
		&log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "02-01-2006 15:04:05",
			ForceColors:            false,
			DisableLevelTruncation: false,
		})


	}
	//if Env.LogFormat == LOG_JSON_FORMAT {
	//	log.SetFormatter(&log.JSONFormatter{
	//		TimestampFormat:        "02-01-2006 15:04:05",
	//	})
	//} else {
	//	log.SetFormatter(
	//		&myFormatter{log.TextFormatter{
	//			FullTimestamp:          true,
	//			TimestampFormat:        "02-01-2006 15:04:05",
	//			ForceColors:            true,
	//			DisableLevelTruncation: true,
	//		}})
	//}

	// TODO: log - add formatters
	// TODO: log - add hooks
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 32 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("\u001B[%dm%s\u001B[0m - [%s] - %s\n %v",
		levelColor, strings.ToUpper(entry.Level.String()), entry.Time.Format(f.TimestampFormat), entry.Message, entry.Data)), nil
}