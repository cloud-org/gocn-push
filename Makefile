# 构建脚本

# set-env copy-config 在这里被依赖 在 build-master 和 build-worker 也被依赖，但是不会执行两次
.PHONY: deploy
deploy: set-env copy-config build-linux upx-linux

.PHONY: build
build: set-env copy-config
	go build -v -o bin/gocn cmd/main.go
	@echo "build gocn success"

.PHONY: build-linux
build-linux: set-env copy-config
	GOOS=linux GOARCH=amd64 go build -v -o bin/gocn-linux cmd/main.go
	@echo "build gocn-linux success"

.PHONY: copy-config
copy-config:
	rm -rf bin && mkdir -p bin && cp config/*.yaml bin/
	@echo "copy config success"

.PHONY: set-env
set-env:
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	@echo "set env success"

.PHONY: docker-build
docker-build:
	docker build -f Dockerfile -t gocn-push:${version} .

# 删除无用的 none 镜像 先删除可能跑的容器，后删除镜像
# 想使用真正的 $ 需要用 $$
# TODO: 删除不存在的镜像/被容器占用的镜像 会有问题
.PHONY: remove-none-images
remove-none-images:
	docker images | awk '$$1=="<none>"' | awk '{print $$3}' | xargs docker rmi

# NOTICE: 需要确保有安装 upx
.PHONY: upx
upx:
	upx -v bin/gocn

.PHONY: upx-linux
upx-linux:
	upx -v bin/gocn-linux