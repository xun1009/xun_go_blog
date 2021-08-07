package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Account struct {
	OauthID string
	GithubId string
	Name string
	Type int
	Phone string
	ImgUrl string
	Email string
	Introduce string
	Avatar string
	Location string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountRepo interface {
	Login(context.Context, *Account) error
	GetAccount(context.Context, *Account) error
	RegisterAccount(context.Context, *Account) error
	UpdateAccount(context.Context, *Account) error
	DeleteAccount(context.Context, *Account) error
	ListAccount(context.Context, *Account) error
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUsecase) Login(ctx context.Context, a *Account) error {
	return uc.repo.Login(ctx, a)
}

func (uc *AccountUsecase) GetAccount(ctx context.Context, a *Account) error {
	return uc.repo.GetAccount(ctx, a)
}

func (uc *AccountUsecase) RegisterAccount(ctx context.Context, a *Account) error {
	return uc.repo.RegisterAccount(ctx, a)
}

func (uc *AccountUsecase) UpdateAccount(ctx context.Context, a *Account) error {
	return uc.repo.UpdateAccount(ctx, a)
}

func (uc *AccountUsecase) DeleteAccount(ctx context.Context, a *Account) error {
	return uc.repo.DeleteAccount(ctx, a)
}

func (uc *AccountUsecase) ListAccount(ctx context.Context, a *Account) error {
	return uc.repo.ListAccount(ctx, a)
}
