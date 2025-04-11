package models

// Groups are used to combine multiple roles which are usually given to users together based on their
// business responsibilities. Groups are designed to simplify permissions management.
// In contrast to roles groups make business sense and often correlate with company’s organizational structure.

// The group owner (owner) is usually an identifier or name that identifies who manages or controls the group.
// Typically, this field is used to identify the person or service responsible for creating, editing,
// or deleting groups.

// Group represents a collection of roles for simplifying permission management.
// type Group struct {
// 	ID          types.ID    // Unique code for the group
// 	Name        string    // Name displayed to end users
// 	Description string    // Overview of the group
// 	Internal    bool      // Indicates if the group is predefined and unmodifiable
// 	Owner       types.ID    // Identifier of the group owner
// 	CreatedAt   time.Time // Timestamp for group creation
// 	UpdatedAt   time.Time // Timestamp for the last group update
// }

// CREATE TABLE admin_groups
// (
//    admin_id BIGINT UNSIGNED,
//    group_id BIGINT UNSIGNED,
//    PRIMARY KEY (admin_id, group_id),
//    FOREIGN KEY (admin_id) REFERENCES admins (id),
//    FOREIGN KEY (group_id) REFERENCES groups (id)
// );

// DROP TABLE IF EXISTS `admin_groups`;

// CREATE TABLE groups
// (
//    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
//    name        VARCHAR(191) NOT NULL,
//    description TEXT,
//    internal    BOOLEAN      NOT NULL DEFAULT FALSE,
//    owner       BIGINT UNSIGNED,
//    created_at  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
//    updated_at  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//    FOREIGN KEY (owner) REFERENCES admins (id)
// );

// DROP TABLE IF EXISTS `groups`;
