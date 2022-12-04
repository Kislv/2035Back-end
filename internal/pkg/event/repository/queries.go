package eventrepository

const (
	queryCheckEvent = `
	SELECT count(*)
	FROM events 
	WHERE title = $1 and longitude = $2 and latitude = $3;
	`
	queryCreateEvent = `
	INSERT INTO
    events (title, description, userId, longitude, latitude, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price)
	VALUES
    (
		$1,
        $2,
        $3, 
        $4,
		$5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13
    )
	RETURNING id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price;
	`
	
	queryGetEventListFirstPart = `
	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
	FROM events
	JOIN events_categories ON events.id = events_categories.eventId
	WHERE events_categories.category IN `
	queryGetEventListSecondPart = ` 
	ORDER BY events.id;
	`

	queryGetAllEventList = `
	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
	FROM events
	ORDER BY events.id;
	`

	queryGetCategoryList = `
	SELECT name, imagePath
	FROM categories
	ORDER BY categories.name;
	`
// 	queryCreateEventCategoryFirstPart = `
// 	INSERT INTO
//     events_categories (eventId, category)
// 	VALUES`
//     queryCreateEventCategorySecondPart = `
// 	(
// 		$`
// 	queryCreateEventCategoryThirdPart = `,
//         $`
// 	queryCreateEventCategoryForthPart = `
//     )`
// 	queryCreateEventCategoryFifthPart = `
// 	RETURNING eventId, category;`

	queryCreateEventCategory = `
	INSERT INTO
	events_categories (eventId, category)
	VALUES
	(
		$1,
		$2
	)
	RETURNING eventId, category;`

	queryGetCertainEvent = `
	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
	FROM events
	WHERE id = $1;
	`

	querySignUpUserForEvent = `
	INSERT INTO 
	users_events (eventId, userId)
	VALUES
	(
        $1,
        $2
	);
	`

	queryCancelSignUpUserForEvent = `
	DELETE FROM 
	users_events
	WHERE eventid = $1 and userid = $2;
	`

	queryGetUserAge = `
	SELECT age 
	FROM users
	WHERE id = $1;
	`

	queryGetEventAges = `
	SELECT minage, maxage
	FROM events
	WHERE id = $1;
	`
)
