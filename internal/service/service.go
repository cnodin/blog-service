package service

import (
	"context"
	"github.com/cnodin/blog-service/global"
	"github.com/cnodin/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
	}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
