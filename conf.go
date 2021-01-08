package puppy

import "time"

const (
	DEFAULT_MAXRECONNECTCOUNT int           = 5
	DEFAULT_RECONNECTTIMEOUT  time.Duration = time.Second * 5
	DEFAULT_BUFFERSIZE        int           = 512
)

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

	IngoreCase bool
}

func DefaultConf() Conf {
	return Conf{
		MaxReconnectCount: DEFAULT_MAXRECONNECTCOUNT,
		ReconnectTimeout:  DEFAULT_RECONNECTTIMEOUT,
		BufferSize:        DEFAULT_BUFFERSIZE,
		addr:              ":8080",
	}
}

type RouterConf struct {
	Root                  string
	IgnoreCase            bool
	FixedPathRedirect     bool
	TrailingSlashRedirect bool
}

func DefaultRouterConf() RouterConf {
	return RouterConf{
		Root:                  "/",
		IgnoreCase:            true,
		FixedPathRedirect:     true,
		TrailingSlashRedirect: true,
	}
}
