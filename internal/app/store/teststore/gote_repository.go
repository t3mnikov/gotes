package teststore

import (
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store"
)

// Gote repo for test memory store
type GoteRepository struct {
    store *Store
    gotes map[int]*model.Gote
}

// Create Gote
func (r GoteRepository) Create(g *model.Gote) error {
    err := g.Validate()
    if err != nil {
        return err
    }

    g.ID = len(r.gotes) + 1

    r.gotes[g.ID] = g

    return nil
}

// Find Gote by ID
func (r GoteRepository) FindByID(userID int, goteID int) (*model.Gote, error) {
    for _, g := range r.gotes {
        if g.ID == goteID && g.UserId == userID {
            return g, nil
        }

    }

    return nil, store.ErrRecordNotFound
}

// Find Gotes by User ID
func (r GoteRepository) FindByUserID(userID int) ([]*model.Gote, error) {
    result := []*model.Gote{}

    for _, g := range r.gotes {
        if g.UserId == userID {
            result = append(result, g)
        }
    }

    return result, nil
}

// Update Gote in memory store
func (r GoteRepository) Update(g *model.Gote) error {
    var ok = false

    for _, gote := range r.gotes {
        if gote.ID == g.ID && gote.UserId == g.UserId {
            ok = true
            break
        }
    }

    if !ok {
        return store.ErrRecordNotFound
    }

    r.gotes[g.ID] = g

    return nil
}

// Delete Gote from memory store
func (r GoteRepository) DeleteByID(userID int, goteID int) error {
    var ok = false

    for _, gote := range r.gotes {
        if gote.ID == goteID && gote.UserId == userID {
            ok = true
            break
        }
    }

    if !ok {
        return store.ErrRecordNotFound
    }

    delete(r.gotes, goteID)

    return nil
}
