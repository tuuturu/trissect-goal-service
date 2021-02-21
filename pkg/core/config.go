package core

import (
	"errors"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (c Config) Validate() error {
	if c.DiscoveryURL == nil {
		return errors.New("discovery url is required")
	}

	err := is.URL.Validate(c.DiscoveryURL.String())
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, is.UTFNumeric),
		validation.Field(&c.DSN, validation.Required),
	)
}

func LoadConfig() (cfg Config) {
	cfg.Port = Get("PORT", "3000")
	cfg.DSN = parseDSN(Get("DSN", ""))
	cfg.DiscoveryURL, _ = url.Parse(Get("DISCOVERY_URL", ""))

	switch Get("LOG_LEVEL", "info") {
	case "debug":
		cfg.LogLevel = logrus.DebugLevel
	default:
		cfg.LogLevel = logrus.InfoLevel
	}

	return cfg
}

func Get(key, defaultValue string) (value string) {
	value = os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
