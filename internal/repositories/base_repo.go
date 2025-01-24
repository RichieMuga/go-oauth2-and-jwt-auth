// Package repositories contains the boilerplate for creating new repo to avoid alot of duplication
package repositories

import "gorm.io/gorm"

// BaseRepo is a config base for new repositories to be created
type BaseRepo struct {
	db *gorm.DB
}

// NewBaseRepo instanciates the NewBaseRepo function
func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

// DB contains the db to be used
func (b *BaseRepo) DB() *gorm.DB {
	return b.db
}
