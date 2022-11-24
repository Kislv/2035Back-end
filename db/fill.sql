INSERT INTO
    users (username, password, email, phonenumber, age)
VALUES
    (
        'Ванька',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'ivan@vk.ru',
        '79167777777',
        '18'
    ),
    (
        'tmp1',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp1@vk.ru',
        '79167777777',
        '59'

    ),
    (
        'tmp2',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp2@vk.ru',
        '79167777777',
        '16'
    ),
    (
        'tmp3',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp3@vk.ru',
        '79167777777',
        '17'
    ),
    (
        'tmp4',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp4@vk.ru',
        '79167777777',
        '20'
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
    categories (name, imagePath)
VALUES
    (
        'sport',
        '/static/event/sport.png'
    ),
    (
        'house',
        '/static/event/sport.png'
    ),
    (
        'business',
        '/static/event/sport.png'
    ),
    (
        'learn',
        '/static/event/sport.png'
    ),
    (
        'food',
        '/static/event/sport.png'
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

INSERT INTO
    users_categories (userId, category)
	VALUES
	(
        3,
        'food'
	),
	(
        3,
        'business'
	);
