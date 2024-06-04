-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text
create TABLE `users` (
                       `id` INT AUTO_INCREMENT,
                       `name` VARCHAR(191),
                       `email` VARCHAR(191),
                       `password` VARCHAR(191) NOT NULL,
                       `mobile` VARCHAR(191) NOT NULL,
                       `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (id),
                        UNIQUE(mobile, email)
);

-- +migrate Down
drop table IF EXISTS users;