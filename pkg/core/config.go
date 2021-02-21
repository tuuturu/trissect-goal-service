package core

import (
	"os"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, is.UTFNumeric),
		validation.Field(&c.DSN, validation.Required),
	)
}

func LoadConfig() (cfg Config) {
	cfg.Port = Get("PORT", "3000")
	cfg.DSN = parseDSN(Get("DSN", ""))

	return cfg
}

func Get(key, defaultValue string) (value string) {
	value = os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
