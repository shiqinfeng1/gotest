package station

import (
	"context"
	"net"
)

type Addr struct {
	MasterIPv4 net.IP
	MasterIPv6 net.IP
	BackupIPv4 net.IP
	BackupIPv6 net.IP
	NetID      string
}
type RepoV3 interface {
	UpdateConnect(ctx context.Context, updateFn func(ctx context.Context, tr *ConnectionV3) (*ConnectionV3, error)) error
	SaveConnect(ctx context.Context, saveFn func(ctx context.Context) (*ConnectionV3, error)) error
	GetConnect(ctx context.Context) (*ConnectionV3, error)
}
