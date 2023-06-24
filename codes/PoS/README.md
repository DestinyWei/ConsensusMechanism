运行顺序:

终端切换到 PoS 目录中

1. 生成 mod 文件
```shell
 go mod init main
```
2. 执行 run 命令
```shell
 go run .\main.go .\blockchain.go .\coins.go .\miners.go .\pos.go
```