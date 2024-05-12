package station

type Connection interface {
	// 连接状态
	Status() bool
	// 主动断开基站
	Connect() error
	// 连接基站
	Disconnect() error
}
