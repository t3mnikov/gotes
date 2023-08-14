package teststore_test

import (
    "github.com/stretchr/testify/assert"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store/teststore"
    "testing"
)

func TestGoteRepository_Create(t *testing.T) {
    s := teststore.New()

    u := model.TestUser(t)

    err := s.User().Create(u)
    assert.NoError(t, err)

    g := model.TestGote(t)
    g.UserId = u.ID

    err = s.Gote().Create(g)

    assert.NoError(t, err)
    assert.True(t, g.ID > 0)
}

func TestGoteRepository_FindByID(t *testing.T) {
    s := teststore.New()

    g, err := s.Gote().FindByID(1, 1)
    assert.Error(t, err)

    u := model.TestUser(t)

    err = s.User().Create(u)
    assert.NoError(t, err)

    g = model.TestGote(t)
    g.UserId = u.ID

    err = s.Gote().Create(g)

    assert.NoError(t, err)
    assert.True(t, g.ID > 0)

    g1, err := s.Gote().FindByID(u.ID, g.ID)
    assert.NoError(t, err)
    assert.NotNil(t, g1)
}

func TestGoteRepository_FindByUserID(t *testing.T) {
    s := teststore.New()

    u := model.TestUser(t)

    err := s.User().Create(u)
    assert.NoError(t, err)

    g := model.TestGote(t)
    g.UserId = u.ID

    err = s.Gote().Create(g)

    assert.NoError(t, err)
    assert.True(t, g.ID > 0)

    // --
    g2 := model.TestGote(t)
    g2.UserId = u.ID

    err = s.Gote().Create(g2)
    assert.NoError(t, err)

    gl, err := s.Gote().FindByUserID(u.ID)
    assert.NoError(t, err)
    assert.True(t, len(gl) == 2)
}

func TestGoteRepository_Update(t *testing.T) {
    s := teststore.New()

    u := model.TestUser(t)

    err := s.User().Create(u)
    assert.NoError(t, err)

    g := model.TestGote(t)
    g.UserId = u.ID

    err = s.Gote().Create(g)

    assert.NoError(t, err)
    assert.True(t, g.ID > 0)

    name := "any name"
    text := "any text"

    g.Name = name
    g.Text = text
    err = s.Gote().Update(g)

    assert.NoError(t, err)
    assert.Equal(t, g.Name, name)
    assert.Equal(t, g.Text, text)
}

func TestGoteRepository_DeleteByID(t *testing.T) {
    s := teststore.New()

    u := model.TestUser(t)

    err := s.User().Create(u)
    assert.NoError(t, err)

    g := model.TestGote(t)
    g.UserId = u.ID

    err = s.Gote().Create(g)

    assert.NoError(t, err)
    assert.True(t, g.ID > 0)

    err = s.Gote().DeleteByID(u.ID, g.ID)
    assert.NoError(t, err)

    err = s.Gote().DeleteByID(u.ID, g.ID)
    assert.Error(t, err)
}
