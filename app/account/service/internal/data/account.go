package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/biz"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) Login(ctx context.Context, g *biz.Account) error {
	return nil
}

func (r *accountRepo) GetAccount(ctx context.Context, g *biz.Account) error {
	return nil
}

func (r *accountRepo) RegisterAccount(ctx context.Context, g *biz.Account) error {
	return nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, g *biz.Account) error {
	return nil
}

func (r *accountRepo) DeleteAccount(ctx context.Context, g *biz.Account) error {
	return nil
}

func (r *accountRepo) ListAccount(ctx context.Context, g *biz.Account) error {
	return nil
}
