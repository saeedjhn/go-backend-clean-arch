-- +migrate Up
INSERT INTO `users` (`name`, `mobile`, `email`, `password`, `created_at`, `updated_at`)
VALUES ('Bob Smith', '09120000001', 'bob.smith@example.com', 'password123', NOW(), NOW()),
       ('Alice Johnson', '09120000002', 'alice.johnson@example.com', 'password123', NOW(), NOW()),
       ('Charlie Brown', '09120000003', 'charlie.brown@example.com', 'password123', NOW(), NOW()),
       ('Diana Prince', '09120000004', 'diana.prince@example.com', 'password123', NOW(), NOW()),
       ('Eve Adams', '09120000005', 'eve.adams@example.com', 'password123', NOW(), NOW()),
       ('Frank Miller', '09120000006', 'frank.miller@example.com', 'password123', NOW(), NOW()),
       ('Grace Hopper', '09120000007', 'grace.hopper@example.com', 'password123', NOW(), NOW()),
       ('Henry Ford', '09120000008', 'henry.ford@example.com', 'password123', NOW(), NOW()),
       ('Ivy Lee', '09120000009', 'ivy.lee@example.com', 'password123', NOW(), NOW()),
       ('Jack Daniels', '09120000010', 'jack.daniels@example.com', 'password123', NOW(), NOW());

-- +migrate Down
DROP TABLE IF EXISTS `users`;