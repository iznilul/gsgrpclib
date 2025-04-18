package client

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"time"
)

const DEV = "dev"
const TEST = "test"
const PROD = "prod"

type Configuration struct {
	ServerConfig ServerConfig
}
type ServerConfig struct {
	RpcServerPort      string
	RpcWecomHost       string
	RpcBusinessHost    string
	RpcBookingHost    string
	RpcTrackHost       string
	Tls                bool
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerNameOverride string
}

var Mode string
var Config *Configuration

func init() {
	if args := os.Args; len(args) > 1 && args[1] == "prod" {
		Mode = PROD
	} else if args := os.Args; len(args) > 1 && args[1] == "test" {
		Mode = TEST
	} else {
		Mode = DEV
	}
	conf := new(Configuration)
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config." + Mode)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrap(err, "failed on reading config file"))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(errors.Wrap(err, "failed on unmarshal config file"))
	}
	Config = conf
}

func SetTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var timeout time.Duration
	if Mode == "dev" {
		timeout = 114514 * time.Second
	} else {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	return ctx, cancel
}
