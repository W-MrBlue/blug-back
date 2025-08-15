package article

import (
	"Blug/pkg/entities"
)

type ServiceInterface interface {
	AddArticle(title string, content string) (*entities.Article, error)
	GetArticle(id int) (*entities.Article, error)
	UpdateArticle(article *entities.Article) (*entities.Article, error)
	DeleteArticle(id int) error
	GetAllArticles() ([]*entities.Article, error)
	GetArticlesByPage(page, limit int) ([]*entities.Article, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) AddArticle(title, content string) (*entities.Article, error) {
	article := &entities.Article{
		Title:   title,
		Content: content,
	}
	if len(content) > 250 {
		article.Abstract = content[:250]
	} else {
		article.Abstract = content
	}
	if err := s.repository.addArticle(article); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Service) GetArticle(id int) (*entities.Article, error) {
	article := &entities.Article{
		Id: id,
	}
	article, err := s.repository.getArticle(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Service) GetAllClasses() ([]*entities.Class, error) {
	classes, err := s.repository.GetAllClasses()
	if err != nil {
		return nil, err
	}
	return classes, err
}

func (s *Service) GetAllArticles() ([]*entities.Article, error) {
	articles, err := s.repository.GetAllArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *Service) UpdateArticle(article *entities.Article) (*entities.Article, error) {
	updatedArticle, err := s.repository.UpdateArticle(article)
	if err != nil {
		return nil, err
	}
	return updatedArticle, nil
}

func (s *Service) DeleteArticle(id int) error {
	article := &entities.Article{
		Id: id,
	}
	if err := s.repository.DeleteArticle(article); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteArticleUndo(id int) error {
	article := &entities.Article{
		Id: id,
	}
	if err := s.repository.DeleteArticleUndo(article); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetArticlesByPage(page, limit int, showDeleted bool) ([]*entities.Article, error) {
	articles, err := s.repository.GetArticlesByPage(page, limit, showDeleted)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *Service) GetArticleCount() int64 {
	return s.repository.GetArticleCount()
}
