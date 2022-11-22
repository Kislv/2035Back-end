package autrepository

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc, password, phonenumber
	FROM users
	WHERE email = $1;
	`

	queryGetById = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE id = $1;
	`

	queryAddUser = `
	INSERT INTO
		users (username, email, password, phonenumber)
	VALUES
		($1, $2, $3, $4)
	RETURNING id;
	`
	
	queryCreateUserCategory = `
	INSERT INTO
	users_categories (userId, category)
	VALUES
	(
		$1,
		$2
	)
	RETURNING userId, category;`

)
