INSERT INTO
    users (username, password, email, phonenumber)
VALUES
    (
        'Ванька',
        '1234abcd', 
        'ivan@vk.ru',
        '79167777777'
    ),
    (
        'tmp1',
        '1234abcd', 
        'tmp1@vk.ru',
        '79167777777'

    ),
    (
        'tmp2',
        '1234abcd',
        'tmp2@vk.ru',
        '79167777777'
    ),
    (
        'tmp3',
        '1234abcd', 
        'tmp3@vk.ru',
        '79167777777'
    ),
    (
        'tmp4',
        '1234abcd',
        'tmp4@vk.ru',
        '79167777777'
    );

INSERT INTO
    events (title, rating, votesnum, userId, longitude, latitude, currentmembersquantity, creatingdate, startdate)
VALUES
    (
        'party1',
        '4.5', 
        10,
        1,
        1.0,
        1.0,
        5,
        '2022-04-10 15:47:24',
        '2022-04-10 15:47:24'
    ),
    (
        'party2',
        '4.9', 
        10,
        1,
        1.0,
        1.0,
        5,
        '2022-04-10 15:47:24',
        '2022-04-10 15:47:24'
    ),
    (
        'party3',
        '4.9', 
        10,
        1,
        1.0,
        1.0,
        5,
        '2022-04-10 15:47:24',
        '2022-04-10 15:47:24'
    );

    INSERT INTO
        events (poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price)
	VALUES
    (
		'/static/event/event.png',
        'userParty',
        '4.5', 
        10,
		'description',
        1,
        1.0,
        1.0,
        5,
        5,
        5,
        '2022-04-10 15:47:24',
        '2022-04-10 15:47:24',
        '2022-04-10 15:47:24',
		5,
		5,
		5
    );

INSERT INTO
    categories (name)
VALUES
    (
        'sport'
    ),
    (
        'house'
    ),
    (
        'business'
    ),
    (
        'learn'
    ),
    (
        'food'
    );

INSERT INTO
    events_categories (eventId, category)
VALUES
    (
        1,
        'sport'
    ),
    (
        2,
        'sport'
    ),
    (
        2,
        'business'
    );
