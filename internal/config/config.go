package config

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()
	localDbConnString := "postgres://user:password@postgres/banking-system?sslmode=disable"
	viper.SetDefault("database_uri", localDbConnString)
	viper.SetDefault("port", 9191)
}
