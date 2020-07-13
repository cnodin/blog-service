package apiv1

import (
	"fmt"
	"github.com/cnodin/blog-service/global"
	"github.com/gin-gonic/gin"
)

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	
}

func (t Tag) List(c *gin.Context) {
	global.Logger.Infof("%s: %s", "tag", "tag.List")
	global.Logger.Info(fmt.Sprintf("%s:%s", "tag", "tag.List"))
}

func (t Tag) Create(c *gin.Context) {
	
}

func (t Tag) Update(c *gin.Context) {
	
}

func (t Tag) Delete(c *gin.Context)  {
	
}