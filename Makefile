NAME = demo-go
PACKAGE = github.com/everywan/demo-server-go
MAIN = $(PACKAGE)/entry

DEFAULT_TAG = demo-go:latest
DEFAULT_BUILD_TAG = 1.10-alpine
REMOTE_IMAGE = ccr.ccs.tencentyun.com/everywan/demo-go

BUILD_FLAGS= -ldflags "-s -w" -mod vendor -v -o  $(NAME) entry/main.go

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

linux:
	@go mod vendor
	@GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS)

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
