INSERT INTO `users` (`name`, `mobile`, `email`, `password`, `created_at`, `updated_at`)
VALUES ('Bob Smith', '09120000001', 'bob.smith@example.com',
        '$2a$10$nYTqqo0nkYbgTEB13Ue05.3FhGJ1fr9nVET1ZzJK5bwoG2gt.BoeO', NOW(), NOW()),
       ('Alice Johnson', '09120000002', 'alice.johnson@example.com',
        '$2a$10$hcAxK4H7mzmyTpU1FWpmlOmzKq688HqUJthTjK3DifT.ZZowG1eje', NOW(), NOW());
