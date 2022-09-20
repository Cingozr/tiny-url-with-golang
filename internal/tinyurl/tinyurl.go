package tinyurl

import (
	"github.com/jinzhu/gorm"
)

type TinyUrl struct {
	gorm.Model
	OriginalUrl string
	TinyUrl     string
}

type TinyUrlService struct {
	DB *gorm.DB
}

type ITinyUrlService interface {
	GetAll() ([]TinyUrl, error)
	Create(tinyUrl TinyUrl) (TinyUrl, error)
	GetByUrlId(ID uint) (TinyUrl, error)
	DeleteByUrlId(ID uint) error
}

func NewTinyUrlService(db *gorm.DB) *TinyUrlService {
	return &TinyUrlService{
		DB: db,
	}
}

func (s *TinyUrlService) GetAll() ([]TinyUrl, error) {
	var tinyUrls []TinyUrl
	if result := s.DB.Find(&tinyUrls); result.Error != nil {
		return tinyUrls, result.Error
	}
	return tinyUrls, nil
}

func (s *TinyUrlService) Create(tinyUrl TinyUrl) (TinyUrl, error) {
	if result := s.DB.Save(&tinyUrl); result.Error != nil {
		return TinyUrl{}, result.Error
	}
	return tinyUrl, nil
}

func (s *TinyUrlService) GetByUrlId(ID uint) (TinyUrl, error) {
	var tinyUrl TinyUrl
	if result := s.DB.First(&tinyUrl, ID); result.Error != nil {
		return TinyUrl{}, result.Error
	}

	return tinyUrl, nil
}

func (s *TinyUrlService) DeleteByUrlId(ID uint) error {
	if result := s.DB.Delete(&TinyUrl{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}
