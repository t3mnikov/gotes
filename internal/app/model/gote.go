package model

import (
    validation "github.com/go-ozzo/ozzo-validation"
)

// Gote data model
type Gote struct {
    ID        int    `json:"id"`
    UserId    int    `json:"user_id"`
    Name      string `json:"name"`
    Text      string `json:"text"`
    CreatedAt int    `json:"created_at"`
    UpdatedAt int    `json:"updated_at"`
}

// Gotes validation
func (g *Gote) Validate() error {
    return validation.ValidateStruct(
        g,
        validation.Field(&g.UserId, validation.Required),
        validation.Field(&g.Name, validation.Required),
        validation.Field(&g.Text, validation.Required),
    )
}
