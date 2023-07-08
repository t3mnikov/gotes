package sqlstore_test

import (
    "github.com/stretchr/testify/assert"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store/sqlstore"
    "testing"
)

func TestUserRepository_Create(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    //TODO make test with already existing user

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)
}

func TestUserRepository_FindByID(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)

    u2, err := s.User().FindByID(u.ID)
    assert.True(t, u2.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)
    assert.True(t, u.ID > 0)
    assert.NoError(t, err)

    u2, err := s.User().FindByEmail(u.Email)
    assert.NoError(t, err)
    assert.NotNil(t, u2)
}
