package bcrypt_test

import (
	"fmt"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
)

func Test_Generate_ValidCost_ReturnsHashedString(t *testing.T) {
	t.Parallel()

	str := "password123"
	cost := bcrypt.DefaultCost

	hashedStr, err := bcrypt.Generate(str, cost)

	if err != nil {
		t.Errorf("bcrypt.Generate failed with error: %v", err)
	}
	if hashedStr == "" {
		t.Error("Expected a non-empty hashed string, got empty string")
	}
}

func Test_Generate_CostGreaterThanMaxCost_ReturnsError(t *testing.T) {
	t.Parallel()

	str := "password123"
	cost := bcrypt.MaxCost + 1

	_, err := bcrypt.Generate(str, cost)

	if err == nil {
		t.Error("Expected error for cost greater than MaxCost, got nil")
	}
}

func TestGenerate_CostLessThanMinCost_UsesDefaultCost(t *testing.T) {
	t.Parallel()

	str := "password123"
	cost := bcrypt.MinCost - 1

	hashedStr, err := bcrypt.Generate(str, cost)

	if err != nil {
		t.Errorf("bcrypt.Generate failed with error: %v", err)
	}
	if hashedStr == "" {
		t.Error("Expected a non-empty hashed string, got empty string")
	}
}

func Test_CompareHashAndSTR_ValidString_ReturnsNoError(t *testing.T) {
	t.Parallel()

	str := "password123"
	cost := bcrypt.DefaultCost
	hashedStr, _ := bcrypt.Generate(str, cost)

	err := bcrypt.CompareHashAndSTR(hashedStr, str)

	if err != nil {
		t.Errorf("bcrypt.CompareHashAndSTR failed with error: %v", err)
	}
}

func Test_CompareHashAndSTR_InvalidString_ReturnsError(t *testing.T) {
	t.Parallel()

	str := "password123"
	invalidStr := "wrongpassword"
	cost := bcrypt.DefaultCost
	hashedStr, _ := bcrypt.Generate(str, cost)

	err := bcrypt.CompareHashAndSTR(hashedStr, invalidStr)

	if err == nil {
		t.Error("Expected error for invalid string, got nil")
	}
}

func TestGenerateAndCompareHashAndSTR_VariousCosts_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	str := "password123"
	// costs := []bcrypt.Cost{bcrypt.MinCost, bcrypt.DefaultCost, bcrypt.MaxCost}
	costs := []bcrypt.Cost{bcrypt.MinCost, bcrypt.DefaultCost}

	for _, cost := range costs {
		t.Run(fmt.Sprintf("Cost_%d", cost), func(t *testing.T) {
			hashedStr, err := bcrypt.Generate(str, cost)
			if err != nil {
				t.Fatalf("bcrypt.Generate failed with error: %v", err)
			}

			err = bcrypt.CompareHashAndSTR(hashedStr, str)
			if err != nil {
				t.Errorf("bcrypt.CompareHashAndSTR failed with error: %v", err)
			}
		})
	}
}

func Test_CompareHashAndSTR_InvalidHash_ReturnsError(t *testing.T) {
	t.Parallel()

	invalidHash := "invalidhash"
	str := "password123"

	err := bcrypt.CompareHashAndSTR(invalidHash, str)

	if err == nil {
		t.Error("Expected error for invalid hash, got nil")
	}
}

func Test_Generate_EmptyString_ReturnsHashedString(t *testing.T) {
	t.Parallel()

	str := ""
	cost := bcrypt.DefaultCost

	hashedStr, err := bcrypt.Generate(str, cost)

	if err != nil {
		t.Errorf("bcrypt.Generate failed with error: %v", err)
	}
	if hashedStr == "" {
		t.Error("Expected a non-empty hashed string, got empty string")
	}
}

func Test_CompareHashAndSTR_EmptyString_ReturnsNoError(t *testing.T) {
	t.Parallel()

	str := ""
	cost := bcrypt.DefaultCost
	hashedStr, _ := bcrypt.Generate(str, cost)

	err := bcrypt.CompareHashAndSTR(hashedStr, str)

	if err != nil {
		t.Errorf("bcrypt.CompareHashAndSTR failed with error: %v", err)
	}
}

func TestGenerateAndCompareHashAndSTR_MinCost_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	str := "password123"
	cost := bcrypt.MinCost

	hashedStr, err := bcrypt.Generate(str, cost)
	if err != nil {
		t.Fatalf("bcrypt.Generate failed with error: %v", err)
	}

	err = bcrypt.CompareHashAndSTR(hashedStr, str)
	if err != nil {
		t.Errorf("bcrypt.CompareHashAndSTR failed with error: %v", err)
	}
}
