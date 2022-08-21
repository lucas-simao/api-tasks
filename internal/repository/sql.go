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
		INSERT INTO tasks (title, description, created_by_user_id) VALUES(?, ?, ?)
	`
	sqlGetTasks = `
		SELECT
			t.id,
			t.title,
			t.description,
			cby.id AS created_by_id,
			cby.name AS created_by_name,
			t.created_at,
			COALESCE(dby.id, 0) AS deleted_by_id,
			COALESCE(dby.name, '') AS deleted_by_name,
			COALESCE(t.deleted_at, "") AS deleted_at,
			t.updated_at,
			COALESCE(t.finished_at, "") AS finished_at
		FROM tasks t
		LEFT JOIN users cby ON cby.id = t.created_by_user_id
		LEFT JOIN users dby ON dby.id = t.deleted_by_user_id
		WHERE true
	`
)
