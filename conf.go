package puppy

import "time"

type Conf struct {
	Id      int64 // 框架给每个 service 分配的唯一id
	Name    string
	Group   int32  // 用户定义 业务所属群组
	Version int32  // 用户定义 业务版本
	Type    int32  // 用户定义 业务类型
	Root    string // 用户定义 群组中的主节点
	// 断线重连
	BufferSize        int           // 缓冲大小
	MaxReconnectCount int           // 最大重连次数 default 5 times
	ReconnectTimeout  time.Duration // default 5 second
	Node              string
	addr              string
	Http              bool
}

func DefaultConf() Conf {
	return Conf{
		MaxReconnectCount: DEFAULT_MAXRECONNECTCOUNT,
		ReconnectTimeout:  DEFAULT_RECONNECTTIMEOUT,
		BufferSize:        DEFAULT_BUFFERSIZE,
		addr:              ":8080",
	}
}
