CREATE TABLE role_wildcard_permissions
(
    role_id          BIGINT UNSIGNED,
    resource_pattern VARCHAR(191) NOT NULL,
    priority         INT       DEFAULT 0,
    permission_id    BIGINT UNSIGNED,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, resource_pattern),
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE ON UPDATE CASCADE
);