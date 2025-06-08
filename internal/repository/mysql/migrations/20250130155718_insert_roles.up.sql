INSERT INTO roles (name, description, internal, created_at, updated_at)
VALUES
-- General Roles
('Super Admin', 'Highest level of access in the system. Can manage users, roles, and permissions.', TRUE, NOW(), NOW()),
('Admin', 'Has the ability to manage users and system settings. May have permission to create or delete critical data.',
 TRUE, NOW(), NOW()),
('Moderator', 'Responsible for monitoring content and approving or removing it.', TRUE, NOW(), NOW()),
('User', 'Limited access to basic system features.', FALSE, NOW(), NOW()),
('Guest', 'A user who has not yet logged into the system.', FALSE, NOW(), NOW()),

-- Organizational Roles
('CEO', 'Access to all reports and managerial data. Can make high-level company decisions.', TRUE, NOW(), NOW()),
('Manager', 'Access to data related to their team or department.', TRUE, NOW(), NOW()),
('Employee', 'Access to tasks, projects, and job-related information.', FALSE, NOW(), NOW()),
('HR', 'Can view and manage employee information.', TRUE, NOW(), NOW()),

-- IT & Software Development Roles
('Developer', 'Access to project code and Git repositories. Can make changes to code.', TRUE, NOW(), NOW()),
('DevOps Engineer', 'Manages servers, CI/CD pipelines, and system operations.', TRUE, NOW(), NOW()),
('QA Engineer', 'Has access to the system for quality testing and bug reporting.', TRUE, NOW(), NOW()),
('Support Engineer', 'Access to user information and debugging tools.', TRUE, NOW(), NOW()),

-- E-commerce Roles
('Vendor', 'Manages their own products and orders.', FALSE, NOW(), NOW()),
('Customer', 'Purchases products and views order history.', FALSE, NOW(), NOW()),
('Delivery Personnel', 'Views orders for delivery purposes.', FALSE, NOW(), NOW());