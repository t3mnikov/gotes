package sqlstore

import (
    "database/sql"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store"
    "time"
)

// Gote repo for DB store
type GoteRepository struct {
    store *Store
}

// Create Gote in DB
func (r *GoteRepository) Create(g *model.Gote) error {
    err := g.Validate()
    if err != nil {
        return err
    }

    ts := int(time.Now().Unix())
    g.CreatedAt = ts
    g.UpdatedAt = ts

    err = r.store.db.QueryRow(
        "insert into gotes (user_id, name, text, created_at, updated_at) values ($1,$2,$3,$4,$5) returning id",
        g.UserId, g.Name, g.Text, g.CreatedAt, g.UpdatedAt,
    ).Scan(&g.ID)

    return err
}

// Find Gote in DB
func (r *GoteRepository) FindByID(userID int, goteID int) (*model.Gote, error) {
    g := &model.Gote{}

    err := r.store.db.QueryRow(
        "select * from gotes where user_id = $1 and id = $2", userID, goteID,
    ).Scan(&g.ID, &g.UserId, &g.Name, &g.Text, &g.CreatedAt, &g.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, store.ErrRecordNotFound
        }
    }

    return g, nil
}

// Find Gotes in DB by user
func (r *GoteRepository) FindByUserID(userID int) ([]*model.Gote, error) {
    gotes := []*model.Gote{}

    rows, err := r.store.db.Query(
        "select * from gotes where user_id = $1", userID,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, store.ErrRecordNotFound
        }
    }

    defer rows.Close()

    for rows.Next() {
        g := &model.Gote{}
        err := rows.Scan(&g.ID, &g.UserId, &g.Name, &g.Text, &g.CreatedAt, &g.UpdatedAt)
        if err != nil {
            return nil, err
        }
        gotes = append(gotes, g)
    }

    return gotes, nil
}

// Update Gote in DB
func (r *GoteRepository) Update(g *model.Gote) error {
    err := g.Validate()
    if err != nil {
        return err
    }

    g.UpdatedAt = int(time.Now().Unix())

    _, err = r.store.db.Exec(
        "update gotes set name = $1, text = $2, updated_at = $3 where id = $4 and user_id = $5",
        g.Name, g.Text, g.UpdatedAt, g.ID, g.UserId,
    )

    if err != nil {
        return err
    }

    return nil
}

// Delete Gote from DB
func (r *GoteRepository) DeleteByID(userID int, goteID int) error {
    g, err := r.FindByID(userID, goteID)
    if err != nil {
        return err
    }

    _, err = r.store.db.Exec(
        "delete from gotes where id = $1",
        g.ID,
    )

    return err
}
