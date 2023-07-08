package gotesserver

import (
    "github.com/labstack/echo/v4"
    "github.com/t3mnikov/gotes/internal/app/store"

    "database/sql"
    "github.com/sirupsen/logrus"
    "github.com/t3mnikov/gotes/internal/app/store/sqlstore"
)

// server struct
type server struct {
    logger *logrus.Logger
    store  store.Store
    echo   *echo.Echo
    config *Config
    // session, jwt ???
}

// Start server
func Start(config *Config) error {
    db, err := newDb(config.DatabaseURL)

    if err != nil {
        return err
    }

    defer db.Close()

    st := sqlstore.New(db)

    _, err = newServer(st, config)

    return err
}

// Make a new DB connection
func newDb(databaseURL string) (*sql.DB, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil

}

// Constructor for server
func newServer(store store.Store, config *Config) (*server, error) {
    e := echo.New()

    s := &server{
        store:  store,
        logger: logrus.New(),
        echo:   e,
        config: config,
    }

    s.configureHandlers()

    err := e.Start(config.BindAddr)
    if err != nil {
        return nil, err
    }

    return s, nil
}

// Constructor for test server
func newTestServer(store store.Store, config *Config) (*server, error) {
    e := echo.New()

    s := &server{
        store:  store,
        logger: logrus.New(),
        echo:   e,
        config: config,
    }

    s.configureHandlers()

    return s, nil
}
