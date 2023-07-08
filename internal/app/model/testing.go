package model

import "testing"

// Testing User
func TestUser(t *testing.T) *User {
    return &User{
        Email:    "user@user.org",
        Password: "123123",
    }
}
