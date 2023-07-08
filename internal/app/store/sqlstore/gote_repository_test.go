package sqlstore_test

import (
    "github.com/stretchr/testify/assert"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store/sqlstore"
    "testing"
)

func TestGoteRepository_Create(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users", "gotes")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    g := model.Gote{
        UserId: u.ID,
        Name:   "test name",
        Text:   "such a new text of new gote",
    }

    err = s.Gote().Create(&g)

    assert.True(t, g.ID > 0)
    assert.True(t, g.CreatedAt > 0)
    assert.True(t, g.UpdatedAt > 0)
    assert.NoError(t, err)
    assert.NotNil(t, g)
}

func TestGoteRepository_Update(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users", "gotes")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    g := model.Gote{
        UserId: u.ID,
        Name:   "test name",
        Text:   "such a new text of new gote",
    }

    err = s.Gote().Create(&g)

    assert.True(t, g.ID > 0)
    assert.True(t, g.CreatedAt > 0)
    assert.True(t, g.UpdatedAt > 0)
    assert.NoError(t, err)
    assert.NotNil(t, g)

    g.Name = "updated name"
    g.Text = "updated text"

    err = s.Gote().Update(&g)
    assert.NoError(t, err)

}

func TestGoteRepository_DeleteByID(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users", "gotes")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    g := model.Gote{
        UserId: u.ID,
        Name:   "test name",
        Text:   "such a new text of new gote",
    }

    err = s.Gote().Create(&g)

    assert.True(t, g.ID > 0)
    assert.True(t, g.CreatedAt > 0)
    assert.True(t, g.UpdatedAt > 0)
    assert.NoError(t, err)
    assert.NotNil(t, g)

    err = s.Gote().DeleteByID(g.UserId, g.ID)
    assert.NoError(t, err)

    err = s.Gote().DeleteByID(g.UserId, g.ID)
    assert.Error(t, err)
}

func TestGoteRepository_FindByID(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users", "gotes")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    gote, err := s.Gote().FindByID(u.ID, 1)
    assert.Error(t, err)
    assert.Nil(t, gote)

    g := model.Gote{
        UserId: u.ID,
        Name:   "test name",
        Text:   "such a new text of new gote",
    }

    err = s.Gote().Create(&g)

    gote2, err := s.Gote().FindByID(u.ID, g.ID)
    assert.True(t, gote2.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, gote2)
}

func TestGoteRepository_FindByUserID(t *testing.T) {
    db, truncate := sqlstore.TestDB(t, DatabaseURL)
    defer truncate("users", "gotes")

    s := sqlstore.New(db)
    u := model.TestUser(t)

    err := s.User().Create(u)

    assert.True(t, u.ID > 0)
    assert.NoError(t, err)
    assert.NotNil(t, u)

    g := model.Gote{
        UserId: u.ID,
        Name:   "test name",
        Text:   "such a new text of new gote",
    }

    err = s.Gote().Create(&g)

    gotes, err := s.Gote().FindByUserID(u.ID)
    assert.True(t, len(gotes) > 0)

    for _, gote := range gotes {
        assert.True(t, checkIsGote(gote))
    }
}

func checkIsGote(i any) bool {
    switch i.(type) {
    case *model.Gote:
        return true
    }

    return false
}
