package common

import (
	db "gabrielricci/stocks/db"
)

type Env struct {
	Repository db.Datastore
}
