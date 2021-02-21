package upper

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

const goalTable = "goals"

type client struct {
	connectionURL *postgresql.ConnectionURL
	Session       db.Session
}
