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
type Repo interface {
	UpdateConnectV3(ctx context.Context, updateFn func(ctx context.Context, tr *ConnectionV3) (*ConnectionV3, error)) error
}
