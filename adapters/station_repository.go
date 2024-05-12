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
func NewStationRepoV3() station.RepoV3 {
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

func (r stationRepo) UpdateConnect(
	ctx context.Context,
	updateFn func(ctx context.Context, tr *station.ConnectionV3) (*station.ConnectionV3, error),
) error {
	return nil
}

func (r stationRepo) SaveConnect(ctx context.Context, saveFn func(ctx context.Context) (*station.ConnectionV3, error)) error {
	return nil
}

func (r stationRepo) GetConnect(ctx context.Context) (*station.ConnectionV3, error) {
	return nil, nil
}
