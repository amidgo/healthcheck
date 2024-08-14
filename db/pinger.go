package dbpinger

import (
	"context"
	"database/sql"
)

type Pinger struct {
	db *sql.DB
}

func New(db *sql.DB) *Pinger {
	return &Pinger{
		db: db,
	}
}

func (p *Pinger) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *Pinger) Name() string {
	return "db"
}
