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

	queryBindBasicPlaylists = `
	INSERT INTO
    	users_playlists (user_id, playlist_id)
	VALUES
    	($1, $2),
    	($1, $3);
	`
)
