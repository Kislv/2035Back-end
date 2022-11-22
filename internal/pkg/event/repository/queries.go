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

	queryGetEventList = `
	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
	FROM events
	JOIN events_categories ON events.id = events_categories.eventId
	WHERE events_categories.category = $1
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
)
