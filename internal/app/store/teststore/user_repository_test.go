package teststore_test

import (
    "github.com/stretchr/testify/assert"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store/teststore"
    "testing"
)

func TestUserRepository_Create(t *testing.T) {
    s := teststore.New()
    u := model.TestUser(t)
    u2 := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    err = s.User().Create(u2)
    assert.Error(t, err)
    assert.Equal(t, u2.ID, 0)
}

func TestUserRepository_FindByID(t *testing.T) {
    s := teststore.New()

    u1 := model.TestUser(t)
    s.User().Create(u1)

    u2, err := s.User().FindByID(u1.ID)

    assert.NoError(t, err)
    assert.NotNil(t, u2)
}

func TestUserRepository_FindByID2(t *testing.T) {
    s := teststore.New()

    u, err := s.User().FindByID(100500)

    assert.Error(t, err)
    assert.Nil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
    s := teststore.New()

    email := "bla@example.com"
    u, err := s.User().FindByEmail(email)

    assert.Error(t, err)
    assert.Nil(t, u)

    u2 := model.TestUser(t)
    s.User().Create(u2)

    u3, err := s.User().FindByEmail(u2.Email)

    assert.NoError(t, err)
    assert.NotNil(t, u3)
}

func TestUserRepository_EmailExists(t *testing.T) {
    s := teststore.New()

    u := model.TestUser(t)
    err := s.User().Create(u)

    assert.NoError(t, err)

    _, err = s.User().EmailExists(u.Email)

    assert.Error(t, err)
}
