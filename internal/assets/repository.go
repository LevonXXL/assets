package assets

import (
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryAssets interface {
	CreateAsset(context.Context, *entity.Asset) (*entity.Asset, error)
	GetAsset(context.Context, *entity.Asset) (*entity.Asset, error)
	DeleteAsset(context.Context, uint32, string) (int64, error)
	GetCount(context.Context) (int, error)
	GetList(context.Context, int, int) ([]entity.Asset, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) RepositoryAssets {
	return &Repository{pool: pool}
}

func (ur *Repository) CreateAsset(ctx context.Context, asset *entity.Asset) (*entity.Asset, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	if err := conn.QueryRow(ctx, createAssetSQL, asset.Name, asset.UId, asset.Data).Scan(
		&asset.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = serviceErrors.ErrNoRows
		}
		return nil, err
	}

	return asset, nil
}

func (ur *Repository) GetAsset(ctx context.Context, asset *entity.Asset) (*entity.Asset, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	if err := conn.QueryRow(ctx, getAssetSQL, asset.Name, asset.UId).Scan(
		&asset.Data,
		&asset.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = serviceErrors.ErrNoRows
		}
		return nil, err
	}
	return asset, nil
}

func (ur *Repository) DeleteAsset(ctx context.Context, uid uint32, name string) (int64, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	ct, err := conn.Exec(ctx, deleteAssetSQL, name, uid)
	if err != nil {
		return 0, err
	}
	return ct.RowsAffected(), nil
}

func (ur *Repository) GetCount(ctx context.Context) (int, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var cnt int
	if err := conn.QueryRow(ctx, getCountAssetSQL).Scan(
		&cnt,
	); err != nil {
		return 0, err
	}
	return cnt, nil
}

func (ur *Repository) GetList(ctx context.Context, limit, offset int) ([]entity.Asset, error) {
	conn, err := ur.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var result []entity.Asset
	rows, err := conn.Query(ctx, getListAssetsSQL, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var asset entity.Asset
		if err = rows.Scan(&asset.Name, &asset.UId, &asset.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, asset)
	}

	return result, nil
}
