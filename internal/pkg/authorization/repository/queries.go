package autrepository

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc, password, phonenumber
	FROM users
	WHERE email = $1;
	`

	queryGetById = `
	SELECT id, username, email, imgsrc, phonenumber
	FROM users
	WHERE id = $1;
	`

	queryAddUser = `
	INSERT INTO
		users (username, email, password, phonenumber, age)
	VALUES
		($1, $2, $3, $4, $5)
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
