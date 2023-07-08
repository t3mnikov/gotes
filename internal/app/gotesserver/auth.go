package gotesserver

import (
    "encoding/json"
    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "github.com/t3mnikov/gotes/internal/app/model"
    "net/http"
    "strings"
    "time"
)

// User tokens
// AuthToken to make restricted routes
// RefreshToken to release a new auth token one
type userTokens struct {
    AuthToken    string `json:"auth_token"`
    RefreshToken string `json:"refresh_token"`
}

// Custom claims for storing adv params
type jwtCustomClaims struct {
    UserID int    `json:"user_id"`
    Name   string `json:"name"`
    Admin  bool   `json:"admin"`
    jwt.RegisteredClaims
}

// Generate user tokens
func (s *server) makeUserTokens(u *model.User) (*userTokens, error) {
    token := s.makeUserAuthToken(u)
    refresh := s.makeUserRefreshToken(u)

    t, err := token.SignedString([]byte(s.config.TokenSecret))
    if err != nil {
        return nil, err
    }

    r, err := refresh.SignedString([]byte(s.config.TokenSecret))
    if err != nil {
        return nil, err
    }

    return &userTokens{t, r}, nil
}

// Generate auth token
func (s *server) makeUserAuthToken(u *model.User) *jwt.Token {
    claims := &jwtCustomClaims{
        u.ID,
        u.Email,
        false,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.config.TokenExpiredHours))),
        },
    }

    // create the token with claims
    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// Generate refresh token
func (s *server) makeUserRefreshToken(u *model.User) *jwt.Token {
    claims := &jwtCustomClaims{
        u.ID,
        u.Email,
        false,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(7*24))),
        },
    }

    // create the token with claims
    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// Similar to singing up
// User is given a tokens
func (s *server) handleTokens() echo.HandlerFunc {
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

        req.Email = strings.ToLower(req.Email)

        u, err := s.store.User().FindByEmail(req.Email)
        if err != nil || !u.ComparePassword(req.Password) {
            er := NewErrorResponse(http.StatusUnauthorized, errWrongEmailOrPassword)
            c.JSON(er.Code, er)
            return err
        }

        tokens, err := s.makeUserTokens(u)

        if err != nil {
            er := NewErrorResponse(http.StatusInternalServerError, errFailedIssueToken)
            c.JSON(er.Code, er)
            return err
        }

        return c.JSON(http.StatusOK, tokens)
    }
}

// Handler for given new pair token if actual is expired
func (s *server) handleRefreshToken() echo.HandlerFunc {

    return func(c echo.Context) error {
        if c.Get("user") == nil {
            return nil
        }

        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*jwtCustomClaims)

        u, err := s.store.User().FindByID(claims.UserID)
        if err != nil {
            er := NewErrorResponse(http.StatusNotAcceptable, err)
            c.JSON(er.Code, er)
            return err
        }

        tokens, err := s.makeUserTokens(u)

        if err != nil {
            er := NewErrorResponse(http.StatusInternalServerError, errFailedIssueToken)
            c.JSON(er.Code, er)
            return err
        }

        return c.JSON(http.StatusOK, tokens)
    }
}
