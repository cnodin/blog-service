package apiv1

import (
	"github.com/cnodin/blog-service/global"
	"github.com/cnodin/blog-service/internal/service"
	"github.com/cnodin/blog-service/pkg/app"
	"github.com/cnodin/blog-service/pkg/convert"
	"github.com/cnodin/blog-service/pkg/errcode"
	"github.com/cnodin/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
