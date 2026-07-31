package services

import (
	"context"
	"errors"
	"nike_store_api/internal/domain"
)

type CommentService struct {
	commentRepository domain.CommentRepository
}

func NewCommentService(repo domain.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: repo,
	}
}

func (s *CommentService) GetByProductID(ctx context.Context, query domain.CommentQuery) ([]domain.Comment, error) {
	if query.ProductID <= 0 {
		return nil, errors.New("product id is valid")
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	if query.Limit <= 0 || query.Limit > 50 {
		query.Limit = 10
	}

	query.Offset = (query.Page - 1) * query.Limit

	if query.Lang != "en" {
		query.Lang = "fa"
	}

	return s.commentRepository.GetByProductID(ctx, query)
}
func (s *CommentService) Create(ctx context.Context, comment *domain.Comment) error {

	if comment.ProductID <= 0 {
		return errors.New("product id is invalid")
	}

	if comment.UserID <= 0 {
		return errors.New("user id is invalid")
	}

	if comment.TitleEn == "" && comment.TitleFa == "" {
		return errors.New("comment title can not be empty")
	}

	if comment.Body == "" {
		return errors.New("comment text can not be empty")
	}

	if comment.Rating < 1 || comment.Rating > 5 {
		return errors.New("comment rate must be between 1 until 5")
	}

	return s.commentRepository.Create(ctx, comment)
}
