package assets

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"assets/pkg/pagination"
	"context"
	"errors"
)

type ServiceAssets interface {
	CreateAsset(context.Context, uint32, string, []byte) (*entity.Asset, error)
	GetAsset(context.Context, uint32, string) ([]byte, error)
	DeleteAsset(context.Context, uint32, string) (int64, error)
	GetList(context.Context, *pagination.Pages) (*pagination.Pages, error)
}

type Service struct {
	repository RepositoryAssets
}

func NewService(assetRepository RepositoryAssets) ServiceAssets {
	return &Service{repository: assetRepository}
}

func (su *Service) CreateAsset(ctx context.Context, uid uint32, name string, data []byte) (*entity.Asset, error) {
	asset := &entity.Asset{
		UId:  uid,
		Name: name,
		Data: data,
	}
	asset, err := su.repository.CreateAsset(
		ctx,
		asset,
	)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (su *Service) GetAsset(ctx context.Context, uid uint32, name string) ([]byte, error) {
	asset := &entity.Asset{
		UId:  uid,
		Name: name,
	}
	asset, err := su.repository.GetAsset(
		ctx,
		asset,
	)
	if err != nil {
		if errors.Is(err, serviceErrors.ErrNoRows) {
			err = serviceErrors.NotFoundError
		}
		return nil, err
	}

	return asset.Data, nil
}

func (su *Service) DeleteAsset(ctx context.Context, uid uint32, name string) (int64, error) {
	return su.repository.DeleteAsset(ctx, uid, name)
}

func (su *Service) GetList(ctx context.Context, pages *pagination.Pages) (*pagination.Pages, error) {
	//Можно добавить через COUNT(*) OVER() AS total
	count, err := su.repository.GetCount(ctx)
	if err != nil {
		return pages, err
	}
	pages = pagination.New(pages.Page, pages.PerPage, count)

	pages.Items, err = su.repository.GetList(ctx, pages.Limit(), pages.Offset())
	if err != nil && !errors.Is(err, serviceErrors.ErrNoRows) {
		return nil, err
	}

	return pages, nil
}
