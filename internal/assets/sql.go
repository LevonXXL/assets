package assets

const (
	getAssetSQL = `SELECT data, created_at
FROM assets WHERE name=$1 AND uid=$2`

	createAssetSQL = `INSERT INTO assets (name, uid, data) VALUES ($1, $2, $3) 
ON CONFLICT (name, uid) DO UPDATE SET (data, created_at) = (EXCLUDED.data, NOW())
RETURNING created_at`

	deleteAssetSQL = `DELETE FROM assets WHERE name=$1 AND uid=$2`

	getCountAssetSQL = `SELECT COUNT(*) AS cnt FROM assets`

	getListAssetsSQL = `SELECT name, uid, created_at FROM assets ORDER BY (name, uid) LIMIT $1 OFFSET $2`
)
