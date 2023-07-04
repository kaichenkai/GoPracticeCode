### 使用GO MODULE管理项目
$env:GO111MODULE = "on"
macos环境:
go env -w GO111MODULE=on

### 设置GO的代理，下载包
$env:GOPROXY = "http://goproxy.cn"
macos环境:
go env -w GOPROXY=http://goproxy.cn