package app

import (
	"context"
	pb "imantask/internal/genproto/ppb"
	"imantask/internal/post-grud/domain"
)

type PostService interface {
	GetPostByID(ctx context.Context, req *pb.ID) (*pb.PostResponse, error)
	GetPage(ctx context.Context,req *pb.Page)(*pb.PostResponseList,error)
}

type postService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) PostService {
	return &postService{
		repo: repo,
	}
}

func (p *postService) GetPostByID(ctx context.Context, req *pb.ID) (*pb.PostResponse, error) {
	post, err := p.repo.GetByID(int(req.ID))
	if err != nil {
		return &pb.PostResponse{}, err
	}
	return &post, nil
}

func (p *postService) GetPage(ctx context.Context ,req *pb.Page) (*pb.PostResponseList,error){
	pageSize := req.PageSize
	offset := (req.PageNumber -1) *pageSize

	posts, err := p.repo.GetPage(int(offset),int(pageSize))
	if err != nil {
		return &pb.PostResponseList{}, err
	}
	return &posts,nil
}