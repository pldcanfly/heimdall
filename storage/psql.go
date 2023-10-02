package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewPostgresStore() (*PostgresStore, error) {
	conStr := "postgres://postgres:mysecretpassword@localhost:5432/postgres"
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, conStr)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresStore{
		pool: pool,
		ctx:  ctx,
	}, nil
}

func (pg *PostgresStore) Close() {
	pg.pool.Close()
}

func (pg *PostgresStore) GetWatchers() ([]WatcherRecord, error) {
	var watchers []WatcherRecord

	rows, err := pg.pool.Query(pg.ctx, "SELECT id, type, endpoint FROM heimdall.watcher")
	if err != nil {
		return nil, fmt.Errorf("GetWatchers: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var watchertype, endpoint string

		if err := rows.Scan(&id, &watchertype, &endpoint); err != nil {
			return nil, fmt.Errorf("GetWatchers: %w", err)
		}
		watchers = append(watchers, WatcherRecord{
			ID:      id,
			Watcher: watchertype,
			URL:     endpoint,
		})
	}

	return watchers, nil
}

func (pg *PostgresStore) InsertWatcher(url string, watcher string) (WatcherRecord, error) {
	return WatcherRecord{}, nil
}
func (pg *PostgresStore) GetLastResponses(watcher int, len int) ([]ResponseRecord, error) {
	var responses []ResponseRecord

	rows, err := pg.pool.Query(pg.ctx, "SELECT watcher, status, responsetime FROM heimdall.logs WHERE watcher = $1 ORDER BY date DESC LIMIT $2", watcher, len)
	if err != nil {
		return nil, fmt.Errorf("GetLastResponses: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var watcher int
		var status bool
		var responsetime time.Duration

		if err := rows.Scan(&watcher, &status, &responsetime); err != nil {
			return nil, fmt.Errorf("GetLastResponses: %w", err)
		}
		responses = append(responses, ResponseRecord{
			Watcher:     watcher,
			Online:      status,
			ReponseTime: responsetime,
		})
	}

	return responses, nil
}
func (pg *PostgresStore) InsertResponse(watcher int, online bool, responsetime time.Duration) error {
	_, err := pg.pool.Exec(pg.ctx, "INSERT INTO heimdall.logs (responsetime, status, watcher) VALUES ($1, $2, $3)", responsetime, online, watcher)
	return err
}
