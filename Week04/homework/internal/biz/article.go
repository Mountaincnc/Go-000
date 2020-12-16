package biz

import (
	"week04.homework.szs/internal/data"
)

type ArticleRepo interface {
	GetArticleByID(id int64) (*data.Article, error)
	SaveArticle(art *data.Article) int64
}
