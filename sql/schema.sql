CREATE TABLE IF NOT EXISTS documents
(
    id           VARCHAR   NOT NULL,
    version      TIMESTAMP NOT NULL,
    content      TEXT      NOT NULL,
    language     VARCHAR   NOT NULL,
    PRIMARY KEY (id, version)
);