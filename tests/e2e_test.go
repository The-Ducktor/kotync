package tests

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *E2ETestSuite) SetupSuite() {
	suite.e = echo.New()
	// Initialize routes and middleware here
}

func (suite *E2ETestSuite) TestUserRegistration() {
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"email":"test@example.com","password":"password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), Register(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "token")
	}
}

func (suite *E2ETestSuite) TestUserLogin() {
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com","password":"password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), Login(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "token")
	}
}

func (suite *E2ETestSuite) TestGetUserInfo() {
	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer valid_token")
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), GetUserInfo(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "email")
	}
}

func (suite *E2ETestSuite) TestCreateManga() {
	req := httptest.NewRequest(http.MethodPost, "/manga", strings.NewReader(`{"title":"Test Manga","author":"Test Author"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), CreateManga(c)) {
		assert.Equal(suite.T(), http.StatusCreated, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "id")
	}
}

func (suite *E2ETestSuite) TestGetManga() {
	req := httptest.NewRequest(http.MethodGet, "/manga/1", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), GetManga(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "title")
	}
}

func (suite *E2ETestSuite) TestUpdateManga() {
	req := httptest.NewRequest(http.MethodPut, "/manga/1", strings.NewReader(`{"title":"Updated Manga","author":"Updated Author"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	if assert.NoError(suite.T(), UpdateManga(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Contains(suite.T(), rec.Body.String(), "title")
	}
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}
