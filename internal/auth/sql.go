package auth

const (
	getUserByLoginSQL = `SELECT id, login, password_hash, created_at, encode(digest($1, 'md5'),'hex') AS secret 
FROM users 
WHERE login=$2`

	createSessionSQL = `INSERT INTO sessions (uid, ip) VALUES ($1, $2) RETURNING id, created_at`

	getSessionByTokenSQL = `SELECT id, uid, created_at, ip
FROM sessions 
WHERE id=$1`

	getLastSessionByTokenSQL = `SELECT id, uid, created_at, ip
FROM active_sessions 
WHERE id=$1`
)
