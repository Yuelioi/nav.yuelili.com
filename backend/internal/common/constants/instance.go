package constants

import (
	"backend/internal/common/config"
	"backend/internal/svc"
	"log"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/conf"
)

var ConfInst config.Config      // 配置实例
var SVCInst *svc.ServiceContext // 上下实例

// 清理路径
func clean() {
	ConfInst.Resource.Announces = filepath.Clean(ConfInst.Resource.Announces)
	ConfInst.Resource.Comments = filepath.Clean(ConfInst.Resource.Comments)
	ConfInst.Resource.Collections = filepath.Clean(ConfInst.Resource.Collections)
	ConfInst.Resource.Icons = filepath.Clean(ConfInst.Resource.Icons)
}

func init() {
	// var configFile = flag.String("f", "etc/config.yaml", "the config file")
	// flag.Parse()

	confData, err := os.ReadFile("etc/main.yaml")
	if err != nil {
		log.Fatal(err)
	}

	conf.LoadFromYamlBytes(confData, &ConfInst)

	// conf.MustLoad(*configFile, &ConfInst)

	clean()

	SVCInst = svc.NewServiceContext(ConfInst)
}
