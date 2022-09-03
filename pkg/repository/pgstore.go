package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type PG struct {
	log *logrus.Entry
	db  *sqlx.DB
	dsn string
}

func NewPG(log *logrus.Logger, dsn string) (*PG, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("err connecting to pg: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("err pinging pg after initing connection: %w", err)
	}
	pg := PG{
		log: log.WithField("component", "pgstore"),
		db:  db,
		dsn: dsn,
	}
	return &pg, nil
}

func (pg *PG) StoreDate(lastVisit time.Time) error {
	query := `INSERT INTO last_visit (id, last_visit) VALUES (0, $1)
			  ON CONFLICT (id) DO UPDATE SET last_visit = $1`
	_, err := pg.db.Exec(query, lastVisit)
	if err != nil {
		return fmt.Errorf("err inserting last_visit: %w", err)
	}
	return nil
}

func (pg *PG) GetDate() (time.Time, error) {
	query := `SELECT last_visit FROM last_visit WHERE id = 0`
	var lastVisit time.Time
	if err := pg.db.Get(&lastVisit, query); err != nil {
		return time.Time{}, fmt.Errorf("err getting last_visit: %w", err)
	}
	return lastVisit, nil
}
