package user

const selectSQL = `
	SELECT
		users.id,
		users.first_name,
		users.last_name,
		users.company,
		users.email,
		users.password_digest,
		roles.name,
		users.last_seen,
		users.updated_at,
		users.created_at
`

const findAllSQL = selectSQL + `
	FROM users
	INNER JOIN roles
	ON users.role_id = roles.id
`

const findByIdSQL = selectSQL + `
	FROM users
	INNER JOIN roles
	ON users.role_id = roles.id
	WHERE users.id = $1
	LIMIT 1
`

const findByEmailSQL = selectSQL + `
	FROM users
	INNER JOIN roles
	ON users.role_id = roles.id
	WHERE users.email = $1
	LIMIT 1
`

const insertUserSQL = `
	INSERT INTO users(
		first_name,
		last_name,
		company,
		email,
		password_digest,
		role_id
	) VALUES (
		$1, $2, $3, $4, $5, (
			SELECT id
			FROM roles
			WHERE name = $6
		)
	) RETURNING id
`

const updateUserSQL = `
	UPDATE users
	SET
		first_name = $2,
		last_name = $3,
		company = $4,
		email = $5,
		password_digest = $6,
		role_id = (
			SELECT id
			FROM roles
			WHERE name = $7
		),
		updated_at = current_timestamp
	WHERE id = $1
`

const updateLastSeenSQL = `
	UPDATE users
	SET
		last_seen = current_timestamp
	WHERE id = $1
`

const deleteUserSQL = `
	DELETE FROM users
	WHERE id = $1
`
