package data

import (
	"fmt"
	"strconv"
)

type Article struct {
	ID int64
	Title string
}

type MockArticleRepo struct {}

func NewMockArticleRepo() *MockArticleRepo {
	return &MockArticleRepo{}
}

func (m *MockArticleRepo) SaveArticle(article *Article) (id int64) {
	fmt.Printf("article title: %s, id: %d saved!\n", article.Title, article.ID)
	return article.ID
}

func (m *MockArticleRepo) GetArticleByID(id int64) (*Article, error) {
	return &Article{
		ID:    id,
		Title: "mock-" + strconv.Itoa(int(id)),
	}, nil
}

