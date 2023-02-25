package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config interface {
	GetCredential() (string, string)
	SetToken(string) error
	Token() string
}

type viperConfig struct {
	Credential struct {
		Key    string
		Secret string
		Token  string
	}
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

func (v *viperConfig) Token() string {
	return v.Credential.Token
}

func (v *viperConfig) SetToken(token string) error {
	v.Credential.Token = token
	viper.Set("credential.token", token)
	return viper.WriteConfig()
}
