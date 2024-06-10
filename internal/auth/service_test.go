package auth

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"assets/internal/mocks"
	"context"
	"errors"
	"testing"
	"time"
)

func TestService_Auth(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	repositoryMocked := mocks.NewRepositoryMocked()
	service := NewService(repositoryMocked, nil)

	var tests = []struct {
		*entity.AuthData
		testname        string
		wantError       error
		wantToken       string
		sessionDuration *time.Duration
	}{
		{&entity.AuthData{
			Login:    "valid",
			Password: "valid",
		}, "valid", nil, "12345", nil},
		{&entity.AuthData{
			Login:    "not valid",
			Password: "not valid",
		}, "not valid", serviceErrors.InvalidLoginPasswordError, "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			token, err := service.Auth(ctx, tt.AuthData, nil)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("got %s, want nil\n", err.Error())
			}
			if token != tt.wantToken {
				t.Errorf("got %s, want '%s'\n", token, tt.wantToken)
			}
		})
	}
}

func TestService_GetSession(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	repositoryMocked := mocks.NewRepositoryMocked()
	sessionDuration := time.Duration(1) * time.Hour
	service := NewService(repositoryMocked, &sessionDuration)

	var tests = []struct {
		token     string
		wantError error
	}{
		{"valid-token", nil},
		{"bad-token", serviceErrors.UnauthorizedError},
		{"expired", serviceErrors.UnauthorizedError},
		{"not-last-session-ok", nil},
		{"not-last-session-not-ok", serviceErrors.UnauthorizedError},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			_, err := service.GetActiveSessionByToken(ctx, tt.token)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("got %+v, want %+v\n", err, tt.wantError)
			}

		})
	}

}
