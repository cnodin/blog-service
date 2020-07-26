package global

import (
	"github.com/cnodin/blog-service/pkg/logger"
	"github.com/cnodin/blog-service/pkg/setting"
)

var (
	ServerSetting	*setting.ServerSettingS
	AppSetting	*setting.AppSettingS
	DatabaseSetting	*setting.DatabaseSettingS
	Logger *logger.Logger
	JWTSetting	*setting.JWTSettngS
)
