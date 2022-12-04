package usrrepository

const (

	queryGetById = `
	SELECT id, username, email, imgsrc, phonenumber, age
	FROM users
	WHERE id = $1;
	`
	
	queryUpdAvatarByUsID = `
	UPDATE users 
	SET imgsrc = $2
	WHERE id = $1;
	`
	QueryGetCategory = `
	SELECT category
	FROM users_categories
	WHERE userId = $1;
	`
)
