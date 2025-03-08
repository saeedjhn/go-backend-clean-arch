package main

import (
	"log"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
)

func main() {
	// Example 1: Create a set of roles
	adminRole := entity.Role{
		ID:          1,
		Name:        "Admin",
		Description: "Has full access to all resources",
		Internal:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	managerRole := entity.Role{
		ID:          2,
		Name:        "Manager",
		Description: "Has access to manage users and content",
		Internal:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Example 2: Create resources
	userManagementAPI := entity.Resource{
		ID:          1,
		Name:        "GET /users",
		Description: "Fetches the list of users",
		Type:        "API",
		Internal:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	orderProcessingAPI := entity.Resource{
		ID:          2,
		Name:        "POST /orders",
		Description: "Creates a new order",
		Type:        "API",
		Internal:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Example 3: Define permissions for roles on resources
	adminPermissions := entity.Permission{
		Allow: entity.RWXD{R: true, W: true, X: true, D: true}, // Full permissions
		Deny:  entity.RWXD{},
	}

	managerPermissions := entity.Permission{
		Allow: entity.RWXD{R: true, W: true, X: false, D: false}, // Read and Write permissions only
		Deny:  entity.RWXD{X: true, D: true},                     // Deny Execute and Delete actions
	}

	// Example 4: Assign permissions to roles
	adminRolePermission := entity.RoleResourcePermission{
		RoleID:      adminRole.ID,
		ResourceID:  userManagementAPI.ID,
		Permissions: adminPermissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	managerRolePermission := entity.RoleResourcePermission{
		RoleID:      managerRole.ID,
		ResourceID:  orderProcessingAPI.ID,
		Permissions: managerPermissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Example 5: Create an admin user and assign roles
	adminUser := entity.Admin{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Mobile:      "123-456-7890",
		Description: "Super Admin with full access",
		Password:    "securepassword",
		Roles:       []entity.Role{adminRole, managerRole},
		Gender:      entity.MaleGender,
		Status:      entity.AdminActiveStatus,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Example Output
	log.Printf("Admin User: %+v\n", adminUser)
	log.Printf("Role Permission: %+v\n", adminRolePermission)
	log.Printf("Manager Permission: %+v\n", managerRolePermission)
}
