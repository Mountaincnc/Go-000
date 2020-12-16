package service

import (
	"context"
	v1 "week04.homework.szs/api/article/v1"
	"week04.homework.szs/internal/biz"
	"week04.homework.szs/internal/data"
)

type ArticleSvc struct {
	a biz.ArticleRepo
	v1.UnimplementedArticleServer
}

func NewArticleSvc(a biz.ArticleRepo) *ArticleSvc {
	return &ArticleSvc{a: a}
}

func (s *ArticleSvc) ContributeArticle(ctx context.Context, req *v1.ArticleRequest) (*v1.ArticleReply, error) {
	art := &data.Article{
		ID:    req.Id,
		Title: req.Title,
	}

	id := s.a.SaveArticle(art)

	return &v1.ArticleReply{Id: id}, nil
}