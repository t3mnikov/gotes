package store

import "github.com/t3mnikov/gotes/internal/app/model"

// UserRepository interface
type UserRepository interface {
    Create(user *model.User) error
    FindByID(int) (*model.User, error)
    FindByEmail(string) (*model.User, error)
    EmailExists(email string) (bool, error)
}

// GoteRepository interface
type GoteRepository interface {
    Create(gote *model.Gote) error
    FindByID(int, int) (*model.Gote, error)
    FindByUserID(int) ([]*model.Gote, error)
    Update(gote *model.Gote) error
    DeleteByID(userID int, goteID int) error
}
