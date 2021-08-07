package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/xun1009/xun_go_blog/api/account"
)

type AccountService struct {
	pb.UnimplementedAccountServer
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AccountReply, error) {
	return &pb.AccountReply{
		Code: 0,
		Message: "登入成功",
		Data: &pb.AccountInfo{
			OauthId: "123232",
			GithubId: "xun1009",
			Name: "liuxun",
			Type: 2,
			Phone: "15019203716",
			ImgUrl: "http://23232.com",
			Email:"liuxun@qq.com",
			Introduce:"我很任性",
			Avatar:"just do it",
			Location: "宝悦公寓",
			Password: "liuxun123",
			CreatedAt: &timestamp.Timestamp{
				Seconds: 1628220570,
			},
		},
	}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.AccountReply, error) {
	return &pb.AccountReply{}, nil
}
func (s *AccountService) RegisterAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AccountReply, error) {
	return &pb.AccountReply{}, nil
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.AccountReply, error) {
	return &pb.AccountReply{}, nil
}
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.AccountCommonReplay, error) {
	return &pb.AccountCommonReplay{}, nil
}
func (s *AccountService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReplay, error) {
	return &pb.ListAccountReplay{}, nil
}
