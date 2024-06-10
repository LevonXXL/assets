package auth

const (
	getUserByLoginSQL = `SELECT id, login, password_hash, created_at, encode(digest($1, 'md5'),'hex') AS secret 
FROM users 
WHERE login=$2`

	createSessionSQL = `INSERT INTO sessions (uid, ip) VALUES ($1, $2) RETURNING id, created_at`

	getSessionByTokenIdSQL = `SELECT id, uid, created_at, ip
FROM sessions 
WHERE id=$1`

	getLastUserSessionsSQL = `SELECT id, uid, created_at
FROM sessions 
WHERE uid=$1
ORDER BY created_at DESC 
LIMIT 1`
)
