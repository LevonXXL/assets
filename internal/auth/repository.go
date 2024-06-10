package auth

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net"
)

type RepositoryAuth interface {
	GetUserByLogin(context.Context, string, string) (*entity.User, error)
	CreateSession(context.Context, uint32, net.IP) (*entity.Session, error)
	GetSessionByToken(context.Context, string) (*entity.Session, error)
	GetLastUserSession(context.Context, uint32) (*entity.Session, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) RepositoryAuth {
	return &Repository{pool: pool}
}

func (ur *Repository) GetUserByLogin(ctx context.Context, login string, secret string) (*entity.User, error) {

	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	user := &entity.User{}

	if err := conn.QueryRow(ctx, getUserByLoginSQL, secret, login).Scan(
		&user.Id,
		&user.Login,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.Secret,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = serviceErrors.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

func (ur *Repository) CreateSession(ctx context.Context, uid uint32, ip net.IP) (*entity.Session, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	session := &entity.Session{
		UId: uid,
	}

	if err := conn.QueryRow(ctx, createSessionSQL, uid, ip).Scan(
		&session.Id,
		&session.CreatedAt,
	); err != nil {
		return nil, err
	}

	return session, nil
}

func (ur *Repository) GetSessionByToken(ctx context.Context, tokenId string) (*entity.Session, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	session := &entity.Session{}

	if err := conn.QueryRow(ctx, getSessionByTokenIdSQL, tokenId).Scan(
		&session.Id,
		&session.UId,
		&session.CreatedAt,
		&session.Ip,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = serviceErrors.ErrNoRows
		}
		return nil, err
	}

	return session, nil
}

func (ur *Repository) GetLastUserSession(ctx context.Context, uId uint32) (*entity.Session, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	session := &entity.Session{}

	if err := conn.QueryRow(ctx, getLastUserSessionsSQL, uId).Scan(
		&session.Id,
		&session.UId,
		&session.CreatedAt,
	); err != nil {
		return nil, err
	}

	return session, nil
}
