package sqlstore

import (
    "database/sql"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store"
    "time"
)

// User DB repository
type UserRepository struct {
    store *Store
}

// Check User already exists
func (r *UserRepository) EmailExists(email string) (bool, error) {
    c := 0
    err := r.store.db.QueryRow(
        `select count(id) as c from users where email ilike $1 limit 1`, email,
    ).Scan(&c)

    if err != nil {
        return true, err
    }

    if c > 0 {
        return true, store.ErrEmailAlreadyExists
    }

    return false, nil
}

// Create a User
func (r *UserRepository) Create(u *model.User) error {
    err := u.Validate()
    if err != nil {
        return err
    }

    err = u.BeforeCreate()
    if err != nil {
        return err
    }

    ts := time.Now().Unix()
    u.CreatedAt = int(ts)
    u.UpdatedAt = int(ts)

    err = r.store.db.QueryRow(
        "insert into users (email, password_hash, created_at, updated_at) values ($1, $2, $3, $4) returning id",
        u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt,
    ).Scan(&u.ID)

    return err
}

// FindByID Find User by ID
func (r *UserRepository) FindByID(ID int) (*model.User, error) {
    u := &model.User{}

    err := r.store.db.QueryRow(
        "select * from users where id = $1",
        ID,
    ).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, store.ErrRecordNotFound
        }
    }

    return u, nil
}

// FindByEmail Find User by Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
    u := &model.User{}

    err := r.store.db.QueryRow(
        "select * from users where email = $1",
        email,
    ).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, store.ErrRecordNotFound
        }
    }

    return u, nil
}
