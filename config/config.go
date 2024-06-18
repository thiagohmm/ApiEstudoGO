package config

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBHOST      string `mapstructure:"DB_HOST"`
	DBPORT      string `mapstructure:"DB_PORT"`
	DBUSER      string `mapstructure:"DB_USER"`
	DBPASS      string `mapstructure:"DB_PASS"`
	DBNAME      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	JWTExpireIn int    `mapstructure:"JWT_EXPIN"`
	TokenAuth   *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Config, error) {
	var c Config

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	c.TokenAuth = jwtauth.New("HS256", []byte(c.JWTSecret), nil)

	return &c, nil
}
