package app

import (
	"context"

	"go-temp/adapters"
	"go-temp/domain/station"
)

type BS struct {
	repo station.Repo
}

func NewBS() *BS {
	return &BS{
		repo: adapters.NewStationRepo(),
	}
}

func (bs *BS) Connect(ctx context.Context) {
	bs.repo.UpdateConnectV3(ctx, func(ctx context.Context, tr *station.ConnectionV3) (*station.ConnectionV3, error) {
		if err := tr.Connect(); err != nil {
			return nil, err
		}
		return tr, nil
	})
}
