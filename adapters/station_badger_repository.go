package adapters

import (
	"context"
	"log"

	"go-temp/domain/station"

	"github.com/dgraph-io/badger/v3"
)

type stationRepo struct {
	db *badger.DB
}

// NewTrainingRepo .
func NewStationRepo() station.Repo {
	opts := badger.
		DefaultOptions("./").
		WithLogger(nil)

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal("could not open key-value store")
	}

	return &stationRepo{
		db: db,
	}
}

func (r stationRepo) UpdateConnectV3(
	ctx context.Context,
	updateFn func(ctx context.Context, tr *station.ConnectionV3) (*station.ConnectionV3, error),
) error {
	return nil
}
