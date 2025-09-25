package storage

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, err
	}
	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Add(c *Contact) error {
	return s.db.Create(c).Error
}

func (s *GORMStore) GetAll() ([]*Contact, error) {
	var out []*Contact
	if err := s.db.Order("id asc").Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (s *GORMStore) GetByID(id int) (*Contact, error) {
	var c Contact
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact introuvable")
		}
		return nil, err
	}
	return &c, nil
}

func (s *GORMStore) Update(id int, name, email string) error {
	c, err := s.GetByID(id)
	if err != nil {
		return err
	}
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	return s.db.Save(c).Error
}

func (s *GORMStore) Delete(id int) error {
	res := s.db.Delete(&Contact{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("contact introuvable")
	}
	return nil
}
