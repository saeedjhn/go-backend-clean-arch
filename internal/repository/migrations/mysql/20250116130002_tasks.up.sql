-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text
CREATE TABLE `tasks`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id`     BIGINT UNSIGNED,
    `title`       VARCHAR(191) NOT NULL,
    `description` TEXT,
    `status`      ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);
