package domain

import "context"

type Comment struct {
	ID        int64  `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"` // خوانده شده از جدول users با JOIN
	TitleEn   string `json:"title_en"`
	TitleFa   string `json:"title_fa"`
	Body      string `json:"body"` // بر اساس زبان درخواست مقداردهی می‌شود (body_fa یا body_en)
	Rating    int    `json:"rating"`
	CreatedAt string `json:"created_at"`
}

type CommentQuery struct {
	ProductID int
	Lang      string
	Page      int
	Limit     int
	Offset    int
}

type CreateCommentRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	TitleEn   string `json:"title_en" binding:"required"`
	TitleFa   string `json:"title_fa" binding:"required"`
	Body      string `json:"body" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
}

type CommentRepository interface {
	GetByProductID(ctx context.Context, query CommentQuery) ([]Comment, error)
	Create(ctx context.Context, comment *Comment) error
}
