# Puppy 服务器框架
一个学习的过程，从 0 开始搭建一套完整的服务器框架，支持 http/tcp。

# 模块
[] Gate
[] Socket
[] Http protocol
[] Gob protocol(for rpc)
[] Router
[] Middleware
[] Service discovery
[] NodeManager
[] Rprc
[] Log

# 目标
达成一个高可用的rpc框架,使用co模型来处理数据流,能自由插拔中间件。    
能随时添加/删除节点而不影响整体业务。      
没有主节点，每个节点都包含所有的节点拓扑，节点能够达成一个蜂窝状的rpc架构。     

# 安装
```golang
go get github.com/eavesmy/puppy
```

# 使用
```golang
package main

import "github.com/eavesmy/puppy"
import "fmt"

func main(){

    node := puppy.New(puppy.Conf{
        Root:    "ip:port,ip2:port",
        Name:    "user.name",
        Group:   1,
        Version: 1,
        Type:    1,
    })

    node.Rpc("user.name", func(ctx *puppy.Context) string {

        // in your program, arg could be typed struct.
            
        // call remote service method 'name'
        return ctx.Call("name")
    })
    
    // Handle http request
    node.Post("user.name", func(ctx *puppy.Context) error {
        return nil
    })

    if err := node.Run(":8080"); err != nil {
        fmt.Println(err)
    }
}

```
