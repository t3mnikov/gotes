package gotesserver

import (
    "encoding/json"
    "github.com/golang-jwt/jwt/v5"
    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
    "github.com/t3mnikov/gotes/internal/app/model"
    "net/http"
    "strconv"
    "strings"
)

// server handlers
func (s *server) configureHandlers() {
    g := s.echo.Group("/gotes")
    g.Use(s.jwtConfigMiddleware())
    g.GET("", s.handleGetGotesList())
    g.GET("/:id", s.handleGetGotesOne())
    g.POST("", s.handleCreateGotes())
    g.DELETE("/:id", s.handleDeleteGotes())
    g.PUT("/:id", s.handleUpdateGotes())

    s.echo.POST("/tokens", s.handleTokens())
    rt := s.echo.Group("/tokens/refresh", s.jwtConfigMiddleware())
    rt.POST("", s.handleRefreshToken())

    usersRouter := s.echo.Group("/users")
    usersRouter.POST("", s.handleCreateUsers())
}

// JWT MMiddleware config
func (s *server) jwtConfigMiddleware() echo.MiddlewareFunc {
    config := echojwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims {
            return new(jwtCustomClaims)
        },
        SigningKey: []byte(s.config.TokenSecret),
    }

    return echojwt.WithConfig(config)
}

// Handler for creating users
func (s *server) handleCreateUsers() echo.HandlerFunc {
    type request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    return func(c echo.Context) error {
        req := &request{}
        err := json.NewDecoder(c.Request().Body).Decode(req)
        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        //TODO rewrite
        req.Email = strings.ToLower(req.Email)

        u := &model.User{
            Email:    req.Email,
            Password: req.Password,
        }

        ex, err := s.store.User().EmailExists(u.Email)
        if ex {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        err = s.store.User().Create(u)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        u.ClearAttributes()
        c.JSON(http.StatusCreated, u)

        return nil
    }
}

// Handler for getting gote
func (s *server) handleGetGotesOne() echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        gote, err := s.store.Gote().FindByID(claims.UserID, id)

        if err != nil {
            er := NewErrorResponse(http.StatusNotFound, err)
            c.JSON(er.Code, er)
            return err
        }

        return c.JSON(http.StatusOK, gote)
    }
}

// Handler for getting gotes list
func (s *server) handleGetGotesList() echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)

        gotes, err := s.store.Gote().FindByUserID(claims.UserID)
        if err != nil {
            er := NewErrorResponse(http.StatusNotFound, err)
            c.JSON(er.Code, er)
            return err
        }

        c.JSON(http.StatusOK, gotes)

        return err
    }
}

// Handler for creating gotes
func (s *server) handleCreateGotes() echo.HandlerFunc {
    type request struct {
        Name string `json:"name"`
        Text string `json:"text"`
    }

    return func(c echo.Context) error {
        req := &request{}
        err := json.NewDecoder(c.Request().Body).Decode(req)
        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)

        // find User
        u, err := s.store.User().FindByID(claims.UserID)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        g := &model.Gote{
            UserId: u.ID,
            Name:   req.Name,
            Text:   req.Text,
        }

        err = s.store.Gote().Create(g)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        return c.JSON(http.StatusCreated, g)
    }
}

// Handler for updating gotes
func (s *server) handleUpdateGotes() echo.HandlerFunc {
    type request struct {
        Name string `json:"name"`
        Text string `json:"text"`
    }

    return func(c echo.Context) error {
        req := &request{}
        err := json.NewDecoder(c.Request().Body).Decode(req)
        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)
        goteID, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        g := &model.Gote{
            ID:     goteID,
            UserId: claims.UserID,
            Name:   req.Name,
            Text:   req.Text,
        }

        err = s.store.Gote().Update(g)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        return c.JSON(http.StatusOK, g)
    }
}

// Handler for deleting gotes
func (s *server) handleDeleteGotes() echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            er := NewErrorResponse(http.StatusBadRequest, err)
            c.JSON(er.Code, er)
            return err
        }

        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)

        err = s.store.Gote().DeleteByID(claims.UserID, id)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        c.JSON(http.StatusAccepted, []int{})

        return nil
    }
}
