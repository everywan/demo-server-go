package utils

import (
	"encoding/base64"
	"log"
	"strings"
)

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

// Print your build info
/*
Need add build flags
	go build -ldflags "
		-X 'github.com/everywan/commons/utils.AppName=${APP_NAME}'             \
		-X 'github.com/everywan/commons/utils.AppVersion=${APP_VERSION}'       \
		-X 'github.com/everywan/commons/utils.BuildVersion=${BUILD_VERSION}'   \
		-X 'github.com/everywan/commons/utils.BuildTime=${BUILD_TIME}'         \
		-X 'github.com/everywan/commons/utils.GitRevision=${GIT_REVISION}'     \
		-X 'github.com/everywan/commons/utils.GitBranch=${GIT_BRANCH}'         \
		-X 'github.com/everywan/commons/utils.GoVersion=${GO_VERSION}'         \
		-s -w
	" -mod vendor -v -o  $(NAME) ${MAIN}
*/
func PrintBuildInfo() {
	v, _ := base64.StdEncoding.DecodeString(BuildVersion)
	BuildVersion = strings.TrimSpace(string(v))

	log.Println("App Name:", AppName)
	log.Println("App Version:", AppVersion)
	log.Println("Build version:", BuildVersion)
	log.Println("Build time:", BuildTime)
	log.Println("Git revision:", GitRevision)
	log.Println("Git branch:", GitBranch)
	log.Println("Golang Version:", GoVersion)
}
