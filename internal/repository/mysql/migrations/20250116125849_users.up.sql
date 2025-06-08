-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text
CREATE TABLE `users`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name`       VARCHAR(191)                        NOT NULL,
    `mobile`     VARCHAR(191)                        NOT NULL UNIQUE,
    `email`      VARCHAR(191)                        NOT NULL UNIQUE,
    `password`   VARCHAR(191)                        NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (mobile, email)
);
