package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config interface {
	GetCredential() (string, string)
	SetToken(string)
	GetToken() string
}

type viperConfig struct {
	Credential struct {
		Key    string
		Secret string
		Token  string
	}
	Token string
}

func NewConfig(configFile string) Config {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	config := viperConfig{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}
	return &config
}

func (v *viperConfig) GetCredential() (string, string) {
	return v.Credential.Key, v.Credential.Secret
}

func (v *viperConfig) GetToken() string {
	return v.Token
}

func (v *viperConfig) SetToken(token string) {
	v.Token = token
	viper.Set("token", token)
	viper.WriteConfig()
}
