运行顺序:

终端切换到 PoW 目录中

1. 生成 mod 文件
```shell
 go mod init pow
```
2. 执行 run 命令
```shell
 go run .\main.go .\block.go .\blockchain.go .\pow.go
```