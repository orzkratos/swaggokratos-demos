package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	pb "github.com/orzkratos/demokratos/demo2kratos/api/myhomepage"
	"github.com/yyle88/mutexmap"
)

type MyhomepageService struct {
	pb.UnimplementedMyhomepageServer
	mp *mutexmap.Map[string, *MyInfo]
}

func NewMyhomepageService() *MyhomepageService {
	mp := mutexmap.New[string, *MyInfo]()
	mp.Set("kratos", &MyInfo{
		Name:                  "kratos",
		MyFavoriteDevLanguage: "golang",
		EnglishName:           "kratos",
		ChineseName:           "名字是啥",
	})
	return &MyhomepageService{
		mp: mp,
	}
}

type MyInfo struct {
	Name                  string
	MyFavoriteDevLanguage string
	EnglishName           string
	ChineseName           string
}

func (s *MyhomepageService) GetMyName(ctx context.Context, req *pb.GetMyNameRequest) (*pb.GetMyNameReply, error) {
	myInfo, ok := s.mp.Get(req.AccountName)
	if !ok {
		return nil, errors.New(404, "NOT_FOUND", "wrong")
	}
	return &pb.GetMyNameReply{
		Name: myInfo.Name,
	}, nil
}

func (s *MyhomepageService) GetMyInfo(ctx context.Context, req *pb.GetMyInfoRequest) (*pb.GetMyInfoReply, error) {
	myInfo, ok := s.mp.Get(req.AccountName)
	if !ok {
		return nil, errors.New(404, "NOT_FOUND", "wrong")
	}
	return &pb.GetMyInfoReply{
		Name:                  myInfo.Name,
		MyFavoriteDevLanguage: myInfo.MyFavoriteDevLanguage,
		EnglishName:           myInfo.EnglishName,
		ChineseName:           myInfo.ChineseName,
	}, nil
}

func (s *MyhomepageService) SetMyInfo(ctx context.Context, req *pb.SetMyInfoRequest) (*pb.SetMyInfoReply, error) {
	if req.Name == "" {
		return nil, errors.New(400, "BAD_PARAM", "wrong")
	}
	s.mp.Set(req.Name, &MyInfo{
		Name:                  req.Name,
		MyFavoriteDevLanguage: req.MyFavoriteDevLanguage,
		EnglishName:           req.EnglishName,
		ChineseName:           req.ChineseName,
	})
	return &pb.SetMyInfoReply{
		Message: "success",
	}, nil
}
