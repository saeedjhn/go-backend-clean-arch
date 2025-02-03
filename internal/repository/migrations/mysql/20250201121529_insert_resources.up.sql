INSERT INTO resources (name, description, type, internal, created_at, updated_at)
VALUES ('/users', 'User management API', 'API', FALSE, NOW(), NOW()),
       ('/users/:id', 'Fetch user details', 'API', FALSE, NOW(), NOW()),
       ('/auth/login', 'User authentication endpoint', 'API', FALSE, NOW(), NOW()),
       ('/auth/logout', 'User logout endpoint', 'API', FALSE, NOW(), NOW()),
       ('/admin/dashboard', 'Admin dashboard API', 'API', TRUE, NOW(), NOW()),
       ('/admin/register', 'Admin dashboard API', 'API', TRUE, NOW(), NOW()),
       ('/admin/settings', 'Admin settings API', 'API', TRUE, NOW(), NOW());