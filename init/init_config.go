package init

import (
	"strings"
	"time"

	"github.com/azcov/evermos-flash-sale/pkg/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config for
type Config struct {
	API struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"api"`
	Database struct {
		Pg struct {
			Host                  string        `mapstructure:"host"`
			Port                  int           `mapstructure:"port"`
			DBName                string        `mapstructure:"dbname"`
			User                  string        `mapstructure:"user"`
			Password              string        `mapstructure:"password"`
			SSLMode               string        `mapstructure:"sslmode"`
			MaxOpenConnection     int           `mapstructure:"max_open_connection"`
			MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
			MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime"`
		} `mapstructure:"pg"`
	} `mapstructure:"database"`
}

// setupMainConfig loads app config to viper
func setupMainConfig() (config *Config) {
	zap.S().Info("Executing init/config")

	conf := false

	if util.IsFileorDirExist("app_config.json") {
		conf = true
		zap.S().Info("Local app_config.json file is found, now assigning it with default config")
		viper.SetConfigFile("app_config.json")
		err := viper.ReadInConfig()
		if err != nil {
			zap.S().Info("err: ", err)
		}
	}

	if !conf {
		zap.S().Fatal("Config is required")
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&config)
	if err != nil {
		zap.S().Fatal("err: ", err)
	}

	zap.S().Info("Config APP_ENV: ", util.GetEnv())

	if !util.IsFileorDirExist("app_config.json") && !util.IsProductionEnv() {
		// open a goroutine to watch remote changes forever
		go func() {
			for {
				time.Sleep(time.Second * 5)

				err := viper.WatchRemoteConfig()
				if err != nil {
					zap.S().Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				viper.Unmarshal(&config)
			}
		}()
	}

	return
}
