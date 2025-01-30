CREATE TABLE admins
(
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    first_name  VARCHAR(191) NOT NULL,
    last_name   VARCHAR(191) NOT NULL,
    email       VARCHAR(191) NOT NULL UNIQUE,
    mobile      VARCHAR(191) NOT NULL,
    description TEXT,
    password    VARCHAR(191) NOT NULL,
    gender      ENUM('male', 'female') NOT NULL,
    status      ENUM('active', 'inactive') NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);