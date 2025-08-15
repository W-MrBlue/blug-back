package article

import (
	"Blug/pkg/entities"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (r *Repository) addArticle(article *entities.Article) error {
	if err := r.Db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) getArticle(article *entities.Article) (*entities.Article, error) {
	if err := r.Db.First(article, article.Id).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *Repository) GetAllArticles() ([]*entities.Article, error) {
	var articles []*entities.Article
	if err := r.Db.Where("status != ?", "deleted").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) GetAllClasses() ([]*entities.Class, error) {
	var classes []*entities.Class
	if err := r.Db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *Repository) GetArticleCount() int64 {
	var cnt int64
	r.Db.Model(&entities.Article{}).Count(&cnt)
	return cnt
}

func (r *Repository) UpdateArticle(article *entities.Article) (*entities.Article, error) {
	// First get the existing article
	existingArticle := &entities.Article{}
	if err := r.Db.First(existingArticle, article.Id).Error; err != nil {
		return nil, err
	}

	// Update the fields
	if err := r.Db.Model(existingArticle).Updates(map[string]interface{}{
		"Title":     article.Title,
		"Content":   article.Content,
		"UpdatedAt": time.Now()}).Error; err != nil {
		return nil, err
	}

	// Return the updated article
	return existingArticle, nil
}

func (r *Repository) DeleteArticle(article *entities.Article) error {
	if err := r.Db.Model(article).Updates(map[string]interface{}{
		"Status":    "deleted",
		"UpdatedAt": time.Now()}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteArticleUndo(article *entities.Article) error {
	if err := r.Db.Model(article).Updates(map[string]interface{}{
		"Status":    "draft",
		"UpdatedAt": time.Now()}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetArticlesByPage(page, limit int, showDeleted bool) ([]*entities.Article, error) {
	var articles []*entities.Article
	if showDeleted {
		if err := r.Db.Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.Db.Where("status != ?", "deleted").Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
			return nil, err
		}
	}
	return articles, nil
}
