-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text
CREATE TABLE `tasks` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    `title` VARCHAR(191) NOT NULL,
    `description` TEXT,
    `status` ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +migrate Down
DROP TABLE IF EXISTS `tasks`;
