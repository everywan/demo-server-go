NAME = demo-go
PACKAGE = github.com/everywan/demo-server-go
MAIN = $(PACKAGE)/entry

// app 构建基础信息
APP_VERSION     := $(shell git describe --abbrev=0 --tags)
BUILD_VERSION   := $(shell git log -1 --oneline | base64)
BUILD_TIME      := $(shell date "+%FT%T%z")
GIT_REVISION    := $(shell git rev-parse --short HEAD)
GIT_BRANCH      := $(shell git name-rev --name-only HEAD)
GO_VERSION      := $(shell go version)
GOOS            := $(shell go env GOOS)

BUILD_FLAGS= -ldflags "
		-X 'github.com/everywan/commons/utils.AppName=${APP_NAME}'             \
		-X 'github.com/everywan/commons/utils.AppVersion=${APP_VERSION}'       \
		-X 'github.com/everywan/commons/utils.BuildVersion=${BUILD_VERSION}'   \
		-X 'github.com/everywan/commons/utils.BuildTime=${BUILD_TIME}'         \
		-X 'github.com/everywan/commons/utils.GitRevision=${GIT_REVISION}'     \
		-X 'github.com/everywan/commons/utils.GitBranch=${GIT_BRANCH}'         \
		-X 'github.com/everywan/commons/utils.GoVersion=${GO_VERSION}'         \
		-s -w
	" -mod vendor -v -o  $(NAME) ${MAIN}

// docker 构建信息
DEFAULT_TAG = demo-go:latest
DEFAULT_BUILD_TAG = 1.10-alpine
REMOTE_IMAGE = ccr.ccs.tencentyun.com/everywan/demo-go
REMOTE_TAG = "$(shell git tag -l --sort=-v:refname|head -1)"

ifeq "$(MODE)" "dev"
	REMOTE_TAG = staging
endif
ifeq "$(REMOTE_TAG)" ""
	REMOTE_TAG = latest
endif
REMOTE_IMAGE_TAG = "$(REMOTE_IMAGE):$(REMOTE_TAG)"

ifeq "$(BUILD_TAG)" ""
	BUILD_TAG = $(DEFAULT_BUILD_TAG)
endif

// 颜色输出定义
CL_RED  = "\033[0;31m"
CL_BLUE = "\033[0;34m"
CL_GREEN = "\033[0;32m"
CL_ORANGE = "\033[0;33m"
CL_NONE = "\033[0m"

define color_out
	@echo $(1)$(2)$(CL_NONE)
endef

docker-build:
	@go mod vendor
	$(call color_out,$(CL_BLUE),"Building binary in docker ...")
	@docker run --rm -v "$(PWD)":/go/src/$(PACKAGE) \
		-w /go/src/$(PACKAGE) \
		golang:$(BUILD_TAG) \
		go build -v -o $(NAME) $(MAIN)
	$(call color_out,$(CL_GREEN),"Building binary ok")

docker: docker-build
	$(call color_out,$(CL_BLUE),"Building docker image ...")
	@docker build -t $(DEFAULT_TAG) .
	$(call color_out,$(CL_GREEN),"Building docker image ok")

push: docker
	@docker tag $(DEFAULT_TAG) $(REMOTE_IMAGE_TAG)
	$(call color_out,$(CL_BLUE),"Pushing image $(REMOTE_IMAGE_TAG) ...")
	@docker push $(REMOTE_IMAGE_TAG)
	$(call color_out,$(CL_ORANGE),"Done")

build:
	@go mod vendor
	@go build $(BUILD_FLAGS)

idl:
	# If build proto failed, make sure you have protoc installed and:
	# go install google.golang.org/protobuf/cmd/protoc-gen-go
	# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	# Building proto for Golang
	# todo 根据 proto name 循环遍历生成包, 然后 push 到指定仓库
	@protoc \
		--go_out=$(PWD)/idl/proto/record \
		--go-grpc_out=require_unimplemented_servers=false:$(PWD)/idl/proto/record \
 		idl/proto/record.proto
	$(call color_out,$(CL_ORANGE),"Done")

mock:
	# go get github.com/golang/mock/gomock
	# Source Mode
	@mockgen -package=mocks -destination internal/tests/mocks/demo.go -source=demo.go
	# Reflect Mode. 当 Interface 有 embedded interface 时反射模式好用
	# @mockgen -package=mocks -destination internal/tests/mocks/demo.go . DemoService
	# 简化写法
	# @mockgen -package=mocks -destination internal/tests/mocks/demo.go \
	#		github.com/xgxw/toddler-go DemoService

.PHONY: all idl
all:
	build
