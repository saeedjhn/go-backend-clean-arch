Feature: User Registration

  As a new user
  I want to register with my name, mobile, email and password
  So that I can access the system

  Background:
    Given the user service is running

#  Scenario: Registration fails when mobile is empty
#    When I register with name "Ali", mobile "", email "ali@test.com", and password "StrongPassword123"
#    Then the registration should fail with error "mobile is required"
#
#  Scenario: Registration fails when email is invalid
#    When I register with name "Ali", mobile "09121111111", email "ali@invalid", and password "StrongPassword123"
#    Then the registration should fail with error "email format is invalid"
#
#  Scenario: Registration fails when password is too short
#    When I register with name "Ali", mobile "09121111111", email "ali@test.com", and password "123"
#    Then the registration should fail with error "password is too short"

#  Scenario: Registration fails due to insecure password
#    Given I register with name "Mona", mobile "09124444444", email "mona@test.com", and password "Password123"
#    When the password "Password123" insecure
#    Then the registration should fail with error "password is insecure"

  Scenario: Registration fails due to duplicate mobile number
    Given I register with name "Sara", mobile "09121234567", email "sara@test.com", and password "Password123"
    When the mobile "09121234567" is already registered
    Then the registration should fail with error "mobile number is not unique"

  Scenario: Registration fails due to duplicate email address
    Given I register with name "Amir", mobile "09127777777", email "existing@test.com", and password "Password123"
    When the email "existing@test.com" is already registered
    Then the registration should fail with error "email address is not unique"

  Scenario: Registration fails due to repository error when checking mobile uniqueness
    Given I register with name "Reza", mobile "09121111111", email "reza@test.com", and password "Password123"
    When the repository returns an error when checking mobile "09121111111"
    Then the registration should fail with an internal error

  Scenario: Registration fails due to repository error when checking email uniqueness
    Given I register with name "Reza", mobile "09121111111", email "reza@test.com", and password "Password123"
    When the repository returns an error when checking email "reza@test.com"
    Then the registration should fail with an internal error

  Scenario: Registration fails due to password hashing error
    Given I register with name "Reza", mobile "09121111111", email "reza@test.com", and password "some-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-passwordsome-bad-password"
    When the password hashing fails
    Then the registration should fail with an internal error

  Scenario: Registration fails due to repository error when creating user
    Given I register with name "Omid", mobile "09129999999", email "omid@test.com", and password "Password123"
    When creating user in repository fails
    Then the registration should fail with an internal error

  Scenario: Registration fails due to outbox service error
    Given I register with name "Nima", mobile "09128888888", email "nima@test.com", and password "Password123"
    When the outbox service fails when saving events
    Then the registration should fail with an internal error

  Scenario: Successful registration
    When I register with name "Ali", mobile "09121234567", email "ali@test.com", and password "StrongPassword123"
    Then the registration should be successful
    And I should receive my user information in the response

