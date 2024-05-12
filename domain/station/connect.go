package station

import (
	"sync/atomic"

	"github.com/gogf/gf/v2/errors/gerror"
)

type ConnectionV3 struct {
	// 是否主动关闭连接基站
	IsShutDown atomic.Bool
	// 是否配置自动连接
	IsAutoConnect atomic.Bool
	// 实际链接基站状态
	LinkStatus atomic.Bool // 0：未连接 1：正在链接 2：已连接
}

// 对外展示和基站的连接状态
func (c *ConnectionV3) Status() bool {
	return !c.IsShutDown.Load() && c.LinkStatus.Load()
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
	if !c.LinkStatus.Load() { // 未链接
		// todo connect
		//  no error
		c.LinkStatus.Store(true)
	}
	c.IsShutDown.Store(false)
	return nil
}
