package app

import (
	"context"
	"log"
	"time"

	"go-temp/adapters"
	"go-temp/domain/station"

	"github.com/gogf/gf/v2/os/gctx"
)

type BS struct {
	repoV3    station.RepoV3
	connectV3 *station.ConnectionV3
}

func NewBS() *BS {
	ctx := gctx.New()
	var err error
	bs := &BS{
		repoV3: adapters.NewStationRepoV3(),
	}
	bs.connectV3, err = bs.repoV3.GetConnect(ctx)
	if err != nil {
		panic(err)
	}

	if bs.connectV3.LinkStatus.Load() == station.Connecting.Int32() {
		panic("invalid link status:" + station.Connecting.String())
	}
	if bs.connectV3.IsAutoConnect.Load() {
		bs.Connect(ctx)
	}
	bs.autoReconnect(ctx)
	// todo 刷新真实链接状态
	return bs
}

func (bs *BS) autoReconnect(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			// 未手动断开连接,或者真实链接断了
			if !bs.connectV3.IsShutDown.Load() || bs.connectV3.LinkStatus.Load() == station.NotConnected.Int32() {
				bs.Connect(ctx)
			}
		case <-ctx.Done():
			ticker.Stop()
			log.Println("exit auto reconnect ok")
			return
		}
	}
}

func (bs *BS) Connect(ctx context.Context) {
	bs.repoV3.SaveConnect(ctx, func(ctx context.Context) (*station.ConnectionV3, error) {
		if err := bs.connectV3.Connect(); err != nil {
			return nil, err
		}
		return bs.connectV3, nil
	})
}
