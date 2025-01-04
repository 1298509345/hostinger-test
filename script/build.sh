# 项目名
name=`cat version | awk '{print $1}'`

# 构建 go项目
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go mod tidy
go build -o name
