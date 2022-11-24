DROP TABLE IF EXISTS users              CASCADE;
DROP TABLE IF EXISTS events             CASCADE;
DROP TABLE IF EXISTS categories         CASCADE;
DROP TABLE IF EXISTS events_categories  CASCADE;
DROP TABLE IF EXISTS users_categories  CASCADE;

CREATE TABLE users (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    username                            VARCHAR(30) NOT NULL,
    password                            VARCHAR(200) NOT NULL,
    email                               VARCHAR(50) NOT NULL,
    phonenumber                         VARCHAR(20) NOT NULL,
    age                                 BIGINT NOT NULL,
-- // TODO write valid path
    imgsrc                              VARCHAR(50) DEFAULT '/static/avatars/profile.svg'
);

CREATE TABLE events (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
-- // TODO write valid path
    poster                              VARCHAR(50) DEFAULT '/static/event/event.png',
    title                               VARCHAR(200) NOT NULL,
    rating                              DOUBLE PRECISION DEFAULT 0,
    votesnum                            BIGINT DEFAULT 0,
    description                         VARCHAR(3000),
    userId                              BIGINT REFERENCES users (id) ON DELETE CASCADE,
    longitude                           DOUBLE PRECISION NOT NULL,
    latitude                            DOUBLE PRECISION NOT NULL,
    currentmembersquantity              BIGINT DEFAULT 1,
    maxmembersquantity                  BIGINT DEFAULT 0,
    minmembersquantity                  BIGINT DEFAULT 0,
    creatingdate                        TIMESTAMP NOT NULL,
    startdate                           TIMESTAMP NOT NULL,
    enddate                             TIMESTAMP,
    minage                              SMALLINT DEFAULT 0,
    maxage                              SMALLINT DEFAULT 0,
    price                               BIGINT DEFAULT 0
);

CREATE TABLE categories (
    name                                VARCHAR(100) NOT NULL PRIMARY KEY,
    imagePath                           VARCHAR(100) DEFAULT '/static/event/sport.svg'
);

CREATE TABLE events_categories (
    eventId                            BIGINT REFERENCES events (id),
    category                            VARCHAR(100) REFERENCES categories (name),
    CONSTRAINT events_categories_id     PRIMARY KEY (eventId, category)
);

CREATE TABLE users_categories (
    userId                              BIGINT REFERENCES users (id),
    category                            VARCHAR(100) REFERENCES categories (name),
    CONSTRAINT users_categories_id     PRIMARY KEY (userId, category)
);
