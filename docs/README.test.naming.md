## Naming tests in Golang

### Naming conventions

Let’s take a look at some of the most established naming strategies for tests.

1. #### Roy Osherove’s naming strategy

   ##### [UnitOfWork_StateUnderTest_ExpectedBehavior]

   A unit of work is a use case in the system that startes with a public method and ends up with one of three types of
   results: a return value/exception, a state change to the system which changes its behavior, or a call to a third
   party (when we use mocks). so a unit of work can be a small as a method, or as large as a class, or even multiple
   classes. as long is it all runs in memory, and is fully under our control.

   ##### Example
   ```
   Test_LoginRequest_EmptyMobileField_ReturnsMobileFieldError
   Test_ProcessPayment_InvalidCardNumber_ReturnsInvalidCardError
   Test_CreateUser_EmptyBody_BadRequestErrorThrown
   Test_AddUser_EmptyEmail_ReturnsEmailFieldError
   Test_DeleteUser_NonExistentUser_ReturnsUserNotFoundError
   Test_LoadProfile_InvalidUserId_ReturnsUserNotFoundError
   Test_VerifyEmail_InvalidToken_ReturnsInvalidTokenError
   Test_SendEmail_InvalidAddress_ReturnsInvalidEmailError
   Test_SearchItems_EmptyQuery_ReturnsNoResults
   Test_UpdateUser_InvalidEmail_ReturnsEmailFormatError
   Test_AuthenticateUser_InvalidCredentials_ReturnsAuthenticationFailedError
   ```

--- 

### References

[Naming tests in Golang](https://medium.com/getground/naming-tests-in-golang-c58c188bb9a1)

[osherove](https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html)