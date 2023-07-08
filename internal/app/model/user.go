package model

import (
    validation "github.com/go-ozzo/ozzo-validation"
    "github.com/go-ozzo/ozzo-validation/is"
    "golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
    ID           int    `json:"id"`
    Email        string `json:"email"`
    Password     string `json:"password,omitempty"`
    PasswordHash string `json:"-"`
    CreatedAt    int    `json:"created_at"`
    UpdatedAt    int    `json:"updated_at"`
}

// For before creating User
func (u *User) BeforeCreate() error {
    if len(u.Password) > 0 {
        enc, err := encryptString(u.Password)

        if err != nil {
            return err
        }

        u.PasswordHash = enc
    }

    return nil
}

// Clear attributes
func (u *User) ClearAttributes() {
    u.Password = ""
}

// Compare password with user hased one
func (u *User) ComparePassword(password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}

// User validation
func (u *User) Validate() error {
    return validation.ValidateStruct(
        u,
        validation.Field(&u.Email, validation.Required, is.Email),
        validation.Field(&u.Password, validation.By(requiredIf(u.PasswordHash == "")), validation.Length(6, 50)),
    )
}

// Encrypting pwd
func encryptString(s string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    if err != nil {
        return "", err
    }

    return string(b), nil
}
