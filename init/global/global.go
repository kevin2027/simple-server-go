package global

import (
	"fmt"
	"log"
	"simple-server-go/init/config"
	"simple-server-go/init/flag"

	"github.com/spf13/viper"
)

var Config *config.Configuration

func init() {
	// if os.Getenv("RELEASE") == "true" {
	// 	log.SetFormatter(&log.JSONFormatter{})
	// 	log.SetReportCaller(true)
	// 	log.SetOutput(io.MultiWriter(logRotation(), os.Stdout))
	// }
	viper.SetConfigFile(flag.Flagconf)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file: %s : %v", flag.Flagconf, err)
		return
	}
	log.Printf("config file is %s", flag.Flagconf)
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Panicf("unmarshal err: %s : %v", flag.Flagconf, err)
	}
}
