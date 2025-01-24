// base_repo.go contains the boilerplate for creating new repo to avoid alot of duplication
package repositories

import "gorm.io/gorm"

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

func (b *BaseRepo) DB() *gorm.DB {
	return b.db
}
