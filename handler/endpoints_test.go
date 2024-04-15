package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/SawitProRecruitment/UserService/util"
)

func TestLogin(t *testing.T) {
	e := echo.New()
	mock := gomock.NewController(t)
	mockInterface := repository.NewMockRepositoryInterface(mock)

	defer mock.Finish()

	s := &Server{
		Repository: mockInterface,
		JWTKey:     []byte("123"),
	}

	hashPassword, _ := util.GenerateHash("testPassword")

	tests := []struct {
		name           string
		request        *http.Request
		header         func(request *http.Request)
		mock           func()
		expectedResp   string
		expectedStatus int
	}{
		{
			name:    "positive case",
			request: httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone_number": "testPhoneNumber", "password": "testPassword"}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserDataByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserData{
					Password: hashPassword,
				}, repository.UserId{
					UserID: 1,
				}, nil)
			},
			expectedResp:   "{\"user_id\":1}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:    "negative case fail get user data",
			request: httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone_number": "testPhoneNumber", "password": "testPassword"}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserDataByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserData{}, repository.UserId{}, errors.New("empty row"))
			},
			expectedResp:   "{\"message\":\"User not found\"}\n",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "negative case wrong password",
			request: httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone_number": "testPhoneNumber", "password": ""}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserDataByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserData{
					Password: "",
				}, repository.UserId{
					UserID: 1,
				}, nil)
			},
			expectedResp:   "{\"message\":\"Wrong password\"}\n",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			test.header(test.request)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(test.request, rec)
			err := s.Login(ctx)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, rec.Code)
			assert.Equal(t, test.expectedResp, rec.Body.String())
		})
	}
}

func TestRegister(t *testing.T) {
	e := echo.New()
	mock := gomock.NewController(t)
	mockInterface := repository.NewMockRepositoryInterface(mock)

	defer mock.Finish()

	s := &Server{
		Repository: mockInterface,
		JWTKey:     []byte("123"),
	}

	tests := []struct {
		name           string
		request        *http.Request
		header         func(request *http.Request)
		mock           func()
		expectedResp   string
		expectedStatus int
	}{
		{
			name:    "positive case",
			request: httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"phone_number": "+62000000001", "password": "Te$tp4ssword", "full_name": "testName"}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{}, nil)
				mockInterface.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(repository.UserId{
					UserID: 1,
				}, nil)
			},
			expectedResp:   "{\"user_id\":1}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:    "negative case user exist",
			request: httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"phone_number": "+62000000001", "password": "Te$tp4ssword", "full_name": "testName"}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{
					UserID: 1,
				}, nil)
			},
			expectedResp:   "{\"message\":\"Phone number already exist\"}\n",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "negative case invalid request",
			request: httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"phone_number": "000000001", "password": "testpassword", "full_name": "t"}`)),
			header: func(request *http.Request) {
				request.Header.Set("Content-Type", "application/json")
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{}, nil)
			},
			expectedResp:   "{\"message\":\"{\\\"phone_number\\\":\\\"Phone number must be at least 10 characters\\\",\\\"full_name\\\":\\\"Full name must be at least 3 characters\\\",\\\"password\\\":\\\"Password must contain at least 1 special (non alpha-numeric) characters\\\"}\"}\n",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			test.header(test.request)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(test.request, rec)
			err := s.Registration(ctx)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, rec.Code)
			assert.Equal(t, test.expectedResp, rec.Body.String())
		})
	}
}

func TestGetUser(t *testing.T) {
	e := echo.New()
	mock := gomock.NewController(t)
	mockInterface := repository.NewMockRepositoryInterface(mock)

	defer mock.Finish()

	s := &Server{
		Repository: mockInterface,
		JWTKey:     []byte("123"),
	}

	tests := []struct {
		name           string
		request        *http.Request
		header         func(request *http.Request)
		mock           func()
		expectedResp   string
		expectedStatus int
	}{
		{
			name:    "positive case",
			request: httptest.NewRequest(http.MethodGet, "/get_profile", nil),
			header: func(request *http.Request) {
				expirationTime := time.Now().Add(15 * time.Minute)
				token, _ := util.CreateToken(1, expirationTime, s.JWTKey)
				authorizationCookie := new(http.Cookie)
				authorizationCookie.Name = "authorization"
				authorizationCookie.Value = token
				authorizationCookie.Expires = expirationTime
				request.AddCookie(authorizationCookie)
			},
			mock: func() {
				mockInterface.EXPECT().GetUserDataById(gomock.Any(), gomock.Any()).Return(repository.UserData{
					PhoneNumber: "123",
					FullName:    "test",
				}, nil)
			},
			expectedResp:   "{\"full_name\":\"test\",\"phone_number\":\"123\"}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "negative case not login",
			request:        httptest.NewRequest(http.MethodGet, "/get_profile", nil),
			header:         func(request *http.Request) {},
			mock:           func() {},
			expectedResp:   "{\"message\":\"Please login\"}\n",
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			test.header(test.request)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(test.request, rec)
			err := s.GetProfile(ctx)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, rec.Code)
			assert.Equal(t, test.expectedResp, rec.Body.String())
		})
	}
}

func TestUpdateUser(t *testing.T) {
	e := echo.New()
	mock := gomock.NewController(t)
	mockInterface := repository.NewMockRepositoryInterface(mock)

	defer mock.Finish()

	s := &Server{
		Repository: mockInterface,
		JWTKey:     []byte("123"),
	}

	phoneNumber := "123"
	fullName := "test"

	tests := []struct {
		name           string
		request        *http.Request
		params         generated.UpdateProfileParams
		header         func(request *http.Request)
		mock           func()
		expectedResp   string
		expectedStatus int
	}{
		{
			name:    "positive case update full name and phone number",
			request: httptest.NewRequest(http.MethodPut, "/update_profile", nil),
			params: generated.UpdateProfileParams{
				PhoneNumber: &phoneNumber,
				FullName:    &fullName,
			},
			header: func(request *http.Request) {
				expirationTime := time.Now().Add(15 * time.Minute)
				token, _ := util.CreateToken(1, expirationTime, s.JWTKey)
				authorizationCookie := new(http.Cookie)
				authorizationCookie.Name = "authorization"
				authorizationCookie.Value = token
				authorizationCookie.Expires = expirationTime
				request.AddCookie(authorizationCookie)
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{}, nil)
				mockInterface.EXPECT().UpdateUserData(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResp:   "{\"user_id\":1}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:    "positive case update phone number",
			request: httptest.NewRequest(http.MethodPut, "/update_profile", nil),
			params: generated.UpdateProfileParams{
				PhoneNumber: &phoneNumber,
			},
			header: func(request *http.Request) {
				expirationTime := time.Now().Add(15 * time.Minute)
				token, _ := util.CreateToken(1, expirationTime, s.JWTKey)
				authorizationCookie := new(http.Cookie)
				authorizationCookie.Name = "authorization"
				authorizationCookie.Value = token
				authorizationCookie.Expires = expirationTime
				request.AddCookie(authorizationCookie)
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{}, nil)
				mockInterface.EXPECT().UpdatePhoneNumber(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResp:   "{\"user_id\":1}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:    "positive case update full name",
			request: httptest.NewRequest(http.MethodPut, "/update_profile", nil),
			params: generated.UpdateProfileParams{
				FullName: &fullName,
			},
			header: func(request *http.Request) {
				expirationTime := time.Now().Add(15 * time.Minute)
				token, _ := util.CreateToken(1, expirationTime, s.JWTKey)
				authorizationCookie := new(http.Cookie)
				authorizationCookie.Name = "authorization"
				authorizationCookie.Value = token
				authorizationCookie.Expires = expirationTime
				request.AddCookie(authorizationCookie)
			},
			mock: func() {
				mockInterface.EXPECT().UpdateFullName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResp:   "{\"user_id\":1}\n",
			expectedStatus: http.StatusOK,
		},
		{
			name:    "negative case not login",
			request: httptest.NewRequest(http.MethodPut, "/update_profile", nil),
			params: generated.UpdateProfileParams{
				PhoneNumber: &phoneNumber,
				FullName:    &fullName,
			},
			header:         func(request *http.Request) {},
			mock:           func() {},
			expectedResp:   "{\"message\":\"Please login\"}\n",
			expectedStatus: http.StatusForbidden,
		},
		{
			name:    "negative case update phone number exist",
			request: httptest.NewRequest(http.MethodPut, "/update_profile", nil),
			params: generated.UpdateProfileParams{
				PhoneNumber: &phoneNumber,
				FullName:    &fullName,
			},
			header: func(request *http.Request) {
				expirationTime := time.Now().Add(15 * time.Minute)
				token, _ := util.CreateToken(1, expirationTime, s.JWTKey)
				authorizationCookie := new(http.Cookie)
				authorizationCookie.Name = "authorization"
				authorizationCookie.Value = token
				authorizationCookie.Expires = expirationTime
				request.AddCookie(authorizationCookie)
			},
			mock: func() {
				mockInterface.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(repository.UserId{
					UserID: 1,
				}, nil)
			},
			expectedResp:   "{\"message\":\"Phone number already exist\"}\n",
			expectedStatus: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			test.header(test.request)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(test.request, rec)
			err := s.UpdateProfile(ctx, test.params)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, rec.Code)
			assert.Equal(t, test.expectedResp, rec.Body.String())
		})
	}
}
