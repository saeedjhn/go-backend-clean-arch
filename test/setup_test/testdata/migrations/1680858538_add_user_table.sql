-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text
CREATE TABLE `users` (
                       `id` INT AUTO_INCREMENT PRIMARY KEY,
                       `name` VARCHAR(191),
                       `mobile` VARCHAR(191) NOT NULL,
                       `email` VARCHAR(191),
                       `password` VARCHAR(191) NOT NULL,
                       `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       UNIQUE(mobile, email)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;