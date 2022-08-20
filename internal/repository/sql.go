package repository

var (
	// authentication
	sqlSignUp            = `INSERT INTO users (name, username, password, user_role_id) VALUES(?, ?, ?, ?)`
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
	sqlGetUserRoleByCode = `SELECT id, name, code FROM users_role WHERE code = ?`

	// tasks
	sqlCreateTask = `
		INSERT INTO tasks (title, description, user_id) VALUES(?, ?, ?)
	`
)
