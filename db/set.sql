DROP TABLE IF EXISTS users              CASCADE;
DROP TABLE IF EXISTS events             CASCADE;
DROP TABLE IF EXISTS categories         CASCADE;
DROP TABLE IF EXISTS events_categories  CASCADE;

CREATE TABLE users (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    username                            VARCHAR(30) NOT NULL,
    password                            VARCHAR(20) NOT NULL,
    email                               VARCHAR(50) NOT NULL,
    phonenumber                         VARCHAR(20) NOT NULL,
    imgsrc                              VARCHAR(50) DEFAULT '/static/avatars/profile.svg'
);

CREATE TABLE events (
    id                                  BIGINT NOT NULL PRIMARY KEY,
    poster                              VARCHAR(50) NOT NULL,
    title                               VARCHAR(200) NOT NULL,
    rating                              DOUBLE PRECISION NOT NULL,
    votesnum                            BIGINT NOT NULL,
    description                         VARCHAR(3000) NOT NULL,
    organization                        VARCHAR(200) NOT NULL,
    longitude                           DOUBLE PRECISION NOT NULL,
    latitude                            DOUBLE PRECISION NOT NULL,
    currentmembersquantity              INTEGER NOT NULL,
    maxmembersquantity                  INTEGER,
    minmembersquantity                  INTEGER,
    creatingdate                        TIMESTAMP NOT NULL,
    startdate                           TIMESTAMP NOT NULL,
    enddate                             TIMESTAMP,
    minage                              SMALLINT,
    maxage                              SMALLINT,
    price                               INTEGER
);

CREATE TABLE categories (
    name                                VARCHAR(100) NOT NULL PRIMARY KEY
);

CREATE TABLE events_categories (
    event_id                            BIGINT REFERENCES events (id),
    category                            VARCHAR(100) REFERENCES categories (name),
    CONSTRAINT events_categories_id     PRIMARY KEY (event_id, category)
);
