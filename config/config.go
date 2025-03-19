package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	TgToken  string `mapstructure:"tgToken"`
	LLM      string `mapstructure:"llm"`
	Model    string `mapstructure:"model"`
	LLMToken string `mapstructure:"llmToken"`
	Preamble string `mapstructure:"preamble"`
}

var Cfg Config

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
