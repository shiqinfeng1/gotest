package station

import (
	"sync/atomic"

	"github.com/gogf/gf/v2/errors/gerror"
)

type status int32

const (
	NotConnected status = 0
	Connecting   status = 1
	Connected    status = 2
)

func (s status) Int32() int32 {
	return int32(s)
}

func (s status) String() string {
	return [...]string{
		"notConnected",
		"is connecting",
		"connected",
	}[int(s)]
}

type ConnectionV3 struct {
	// 是否主动关闭连接基站
	IsShutDown atomic.Bool
	// 是否配置自动连接
	IsAutoConnect atomic.Bool
	// 实际链接基站状态
	LinkStatus atomic.Int32 // 0：未连接 1：正在链接 2：已连接
}

// 对外展示和基站的连接状态
func (c *ConnectionV3) Status() bool {
	return !c.IsShutDown.Load() && c.LinkStatus.Load() == Connected.Int32()
}

// 主动断开基站，只记录断开操作，如果物理上已连接，那么不断开物理连接
func (c *ConnectionV3) Disconnect() error {
	if c.IsShutDown.Load() {
		return gerror.New("already shutdown")
	}
	c.IsShutDown.Store(true)
	return nil
}

// 连接基站，如果物理上没连接，那么启动连接
func (c *ConnectionV3) Connect() error {
	// 正在链接
	if c.LinkStatus.Load() == Connecting.Int32() {
		return gerror.New(Connecting.String())
	}
	// 基站已连接
	if c.LinkStatus.Load() == Connected.Int32() {
		// 重复连接
		if !c.IsShutDown.Load() {
			return gerror.New(Connected.String())
		}
		c.IsShutDown.Store(false)
		return nil
	}
	// 未连接
	if c.LinkStatus.Load() == NotConnected.Int32() { // 未链接
		c.LinkStatus.Store(Connecting.Int32())
		// todo connect
		//  no error
		var err error
		if err != nil {
			c.LinkStatus.Store(NotConnected.Int32())
			return err
		}
		c.LinkStatus.Store(Connected.Int32())
		c.IsShutDown.Store(false)
	}
	return nil
}
