package store

// Store interface of data models
type Store interface {
    User() UserRepository
    Gote() GoteRepository
}
