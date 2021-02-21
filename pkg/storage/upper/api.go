package upper

import (
	"github.com/tuuturu/trissect-goal-service/pkg/core"
	"github.com/upper/db/v4/adapter/postgresql"
)

func NewClient(dsn core.DSN) core.StorageClient {
	return &client{
		connectionURL: &postgresql.ConnectionURL{
			User:     dsn.Username,
			Password: dsn.Password,
			Host:     dsn.URI,
			Database: dsn.DatabaseName,
		},
	}
}
