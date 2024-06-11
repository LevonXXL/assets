package mocks

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"net"
	"time"
)

type RepositoryMocked struct {
	pool *pgxpool.Pool
}

func (rm *RepositoryMocked) GetUserByLogin(ctx context.Context, login string, secret string) (*entity.User, error) {
	switch {
	case login == "valid":
		return &entity.User{Id: 1, Login: login, PasswordHash: "valid", CreatedAt: time.Now(), Secret: secret}, nil
	case login == "not valid":
		return &entity.User{Id: 2, Login: login, PasswordHash: "valid", CreatedAt: time.Now(), Secret: secret}, nil
	}
	return nil, nil
}
func (rm *RepositoryMocked) CreateSession(ctx context.Context, uid uint32, ip net.IP) (*entity.Session, error) {
	switch {
	case uid == 1:
		return &entity.Session{Id: "12345", UId: 1, CreatedAt: time.Now()}, nil
	}
	return nil, nil
}
func (rm *RepositoryMocked) GetSessionByToken(ctx context.Context, token string) (*entity.Session, error) {
	switch {
	case token == "valid-token":
		return &entity.Session{Id: "valid-token", UId: 1, CreatedAt: time.Now()}, nil
	case token == "bad-token":
		return nil, serviceErrors.UnauthorizedError
	case token == "expired":
		return &entity.Session{Id: "expired", UId: 1, CreatedAt: time.Now().Add(time.Duration(-5) * time.Hour)}, nil
	case token == "not-last-session-ok":
		return &entity.Session{Id: "not-last-session-ok", UId: 13, CreatedAt: time.Now()}, nil
	case token == "not-last-session-not-ok":
		return &entity.Session{Id: "not-last-session-not-ok", UId: 9999999, CreatedAt: time.Now()}, nil
	}

	return &entity.Session{}, nil
}

func (rm *RepositoryMocked) GetLastUserSession(ctx context.Context, uid uint32) (*entity.Session, error) {
	switch {
	case uid == 13:
		return &entity.Session{Id: "not-last-session-ok", UId: 13, CreatedAt: time.Now()}, nil
	case uid == 9999999:
		return &entity.Session{Id: "123456", UId: 13, CreatedAt: time.Now()}, nil
	}
	if uid == 13 {

	}
	return nil, serviceErrors.ErrNoRows
}

func (rm *RepositoryMocked) GetLastSessionByToken(ctx context.Context, token string) (*entity.Session, error) {
	switch {
	case token == "valid-token":
		return &entity.Session{Id: "valid-token", UId: 1, CreatedAt: time.Now()}, nil
	case token == "bad-token":
		return nil, serviceErrors.UnauthorizedError
	case token == "expired":
		return &entity.Session{Id: "expired", UId: 1, CreatedAt: time.Now().Add(time.Duration(-5) * time.Hour)}, nil
	case token == "not-last-session-ok":
		return &entity.Session{Id: "not-last-session-ok", UId: 13, CreatedAt: time.Now()}, nil
	case token == "not-last-session-not-ok":
		return &entity.Session{Id: "not-last-session-not-ok", UId: 9999999, CreatedAt: time.Now()}, nil
	}

	return &entity.Session{}, nil
}

func NewRepositoryMocked() *RepositoryMocked {
	return &RepositoryMocked{}
}
