package gotesserver

import (
    "bytes"
    "encoding/json"
    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store/teststore"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGotesServer_Auth(t *testing.T) {
    store := teststore.New()
    s, err := newTestServer(store, NewTestConfig())
    assert.NoError(t, err)

    u := model.TestUser(t)
    err = store.User().Create(u)
    assert.NoError(t, err)

    testCases := []struct {
        name         string
        tokenFunc    func() string
        expectedCode int
    }{
        {
            name: "auth",
            tokenFunc: func() string {
                tokens, err := s.makeUserTokens(u)
                assert.NoError(t, err)
                return tokens.AuthToken
            },
            expectedCode: http.StatusOK,
        },
        {
            name: "not auth",
            tokenFunc: func() string {
                return "kadabra"
            },
            expectedCode: http.StatusUnauthorized,
        },
    }

    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            rec := httptest.NewRecorder()

            req, err := http.NewRequest(http.MethodGet, "/gotes", nil)
            assert.NoError(t, err)
            req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
            req.Header.Add("Authorization", "Bearer "+testCase.tokenFunc())

            s.echo.ServeHTTP(rec, req)

            assert.Equal(t, testCase.expectedCode, rec.Code)
        })
    }
}

func TestGotesServer_HandleCreateUser2(t *testing.T) {
    store := teststore.New()
    s, err := newTestServer(store, NewTestConfig())
    assert.NoError(t, err)

    testCases := []struct {
        name         string
        payload      any
        expectedCode int
    }{
        {
            name: "ok",
            payload: map[string]string{
                "email":    "johnvan@caneghem.com",
                "password": "mk89nj54",
            },
            expectedCode: http.StatusCreated,
        },
        {
            name:         "invalid payload",
            payload:      "kadabra",
            expectedCode: http.StatusBadRequest,
        },
        {
            name: "create user",
            payload: map[string]string{
                "email":    "user1@mail.com",
                "password": "123321",
            },
            expectedCode: http.StatusCreated,
        },
        {
            name: "cannot create because email exists",
            payload: map[string]string{
                "email":    "user1@mail.com",
                "password": "123321",
            },
            expectedCode: http.StatusNotAcceptable,
        },
        {
            name: "wrong email",
            payload: map[string]string{
                "email":    "user1mail.com",
                "password": "123321",
            },
            expectedCode: http.StatusNotAcceptable,
        },
        {
            name: "wrong password",
            payload: map[string]string{
                "email":    "user1@mail.com",
                "password": "-",
            },
            expectedCode: http.StatusNotAcceptable,
        },
    }

    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            rec := httptest.NewRecorder()

            b := &bytes.Buffer{}
            err := json.NewEncoder(b).Encode(testCase.payload)
            assert.NoError(t, err)

            req, err := http.NewRequest(http.MethodPost, "/users", b)
            assert.NoError(t, err)
            req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

            s.echo.ServeHTTP(rec, req)

            assert.Equal(t, testCase.expectedCode, rec.Code)
        })
    }
}
