# golang-protocol
golang protoc json swagger 协议相关数据结构的转换工具集合


安装：

进入该项目根目录 运行如下命令：

    go build -ldflags "-s -w" -o gg main.go version.go

gg 表示运行程序名称 这里你可以自己定义即可。

把这个可运行的二进制程序 放到你的 go/bin 中即可。前提是你已经设置好了对应的go env


1. 复制需要转换的go struct 完整的结构体：

```golang
 type User struct {
	Name string
	Age int
}

```

2. 运行：

    gg pb


3. 粘贴到你的protoc文件中即可：
    
        message user {
            string name = 1;
            int32 age = 2;
        }
