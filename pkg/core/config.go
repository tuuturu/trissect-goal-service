package core

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"os"
)

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, is.UTFNumeric),
	)
}

func LoadConfig() *Config {
	return &Config{
		Port: Get("PORT", "3000"),
	}
}

func Get(key, defaultValue string) (value string) {
	value = os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
