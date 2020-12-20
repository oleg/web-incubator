CREATE USER gwp WITH PASSWORD 'gwp';
CREATE DATABASE gwp OWNER gwp;

\connect gwp;

CREATE TABLE posts
(
    id      serial primary key,
    content text,
    author  varchar(255)
);

ALTER TABLE posts OWNER TO gwp;