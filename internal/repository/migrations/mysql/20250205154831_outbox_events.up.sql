CREATE TABLE outbox_events
(
    id              BIGINT PRIMARY KEY AUTO_INCREMENT,
    type            VARCHAR(191) NOT NULL,
    payload         BLOB         NOT NULL,
    is_published    BOOLEAN  DEFAULT FALSE,
    retried_count   INT UNSIGNED DEFAULT 0,
    last_retried_at DATETIME DEFAULT NULL,
    published_at    DATETIME DEFAULT NULL
);
