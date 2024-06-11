package auth

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"context"
	"errors"
	"net"
	"time"
)

type ServiceAuth interface {
	Auth(context.Context, *entity.AuthData, net.IP) (string, error)
	GetActiveSessionByToken(context.Context, string) (*entity.Session, error)
}

type Service struct {
	repository      RepositoryAuth
	sessionDuration *time.Duration
}

func NewService(repository RepositoryAuth, sessionDuration *time.Duration) ServiceAuth {
	return &Service{repository: repository, sessionDuration: sessionDuration}
}

func (su *Service) Auth(ctx context.Context, authData *entity.AuthData, ip net.IP) (string, error) {

	user, err := su.repository.GetUserByLogin(ctx, authData.Login, authData.Password)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", serviceErrors.InvalidLoginPasswordError
	}

	if user.Secret != user.PasswordHash {
		return "", serviceErrors.InvalidLoginPasswordError
	}

	session, err := su.repository.CreateSession(ctx, user.Id, ip)
	if err != nil {
		return "", err
	}

	return session.Id, nil
}

func (su *Service) GetActiveSessionByToken(ctx context.Context, token string) (*entity.Session, error) {
	session, err := su.repository.GetSessionByToken(ctx, token)
	if err != nil {
		if errors.Is(err, serviceErrors.ErrNoRows) {
			return nil, serviceErrors.UnauthorizedError
		}
		return nil, err
	}

	//Если не установлена продолжительность сессии - сессии бессрочные,
	//иначе проверяем условие, что сессия еще не истекла
	if su.sessionDuration != nil &&
		session.CreatedAt.Add(*su.sessionDuration).Unix() < time.Now().Unix() {
		return nil, serviceErrors.UnauthorizedError
	}

	return session, nil
}
