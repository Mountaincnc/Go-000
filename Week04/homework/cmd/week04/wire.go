// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"week04.homework.szs/internal/biz"
	"week04.homework.szs/internal/data"
	"week04.homework.szs/internal/service"
)

var ArticleRepoSet = wire.NewSet(data.NewMockArticleRepo, wire.Bind( new(biz.ArticleRepo), new(*data.MockArticleRepo)))

func InitializeArticle() *service.ArticleSvc {
	wire.Build(ArticleRepoSet, service.NewArticleSvc)
	return &service.ArticleSvc{}
}