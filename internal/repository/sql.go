package repository

var (
	sqlSignUp            = `INSERT INTO users (name, username, password, user_role_id) VALUES(?, ?, ?, ?)`
	sqlGetUserRoleByCode = `SELECT id, name, code FROM users_role WHERE code = ?`
	sqlGetUserByUsername = `
		SELECT 
			users.id,
			users.name,
			username,
			password,
			ur.code AS code_role
		FROM users
		LEFT JOIN users_role ur ON ur.id = users.user_role_id
		WHERE username = ?`
)
