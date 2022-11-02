package global

import (
	"github.com/zhangyusheng/blog-service/pkg/logger"
	"github.com/zhangyusheng/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
