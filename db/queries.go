package db

const createGVTable = `CREATE TABLE GEO_VELOCITY_EVENTS (
    ID         INTEGER       PRIMARY KEY AUTOINCREMENT
                             NOT NULL,
    UUID       BLOB (16)     UNIQUE
                             NOT NULL,
    LOGIN_TIME BIGINT (10)   NOT NULL,
    USERNAME   VARCHAR (255) NOT NULL,
    CREATED    DATETIME      DEFAULT (CURRENT_TIMESTAMP),
    IP_ADDRESS TEXT          NOT NULL,
    LAT        DECIMAL       NOT NULL,
    LONG       DECIMAL       NOT NULL,
    RADIUS     INTEGER (3)   NOT NULL
);
`
