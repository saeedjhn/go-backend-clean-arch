### When to Use General or Specific Interfaces in Go

In Go projects, the choice between general interfaces (e.g., Sender, MessageSender) and specific interfaces (e.g.,
SMSSender, EmailSender) depends on the design goals and use case of the project.

#### 1. General Design (More Flexible and Reusable)

- Use Case:

  General interfaces are ideal when you need flexibility to handle multiple use cases (e.g., SMS, Email, Notification).


- Advantages:

    - Promotes reusability across different communication channels.
    - Keeps the codebase flexible for future expansions.
    - Reduces dependencies on specific implementations.
    ```
    type MessageSender interface {
      Send(destination, message string) error
    }
    ```

#### 2. Specific Design (Limited to One Use Case)

- Use Case:

  Specific interfaces are suitable when the behavior is limited to a single use case, like SMS or Email, and is unlikely
  to be reused for other purposes.


- Advantages:

    - Simpler to implement for a single service.
    - Explicitly defines the use case, making the code easier to understand in its specific context.
    ```
    type SMSSender interface {
      Send(destination, message string) error
    }
    ```

#### Answer to General or Specific?

- General interfaces are more commonly used, especially in libraries or packages designed for extensibility.
- Specific interfaces are better for tightly scoped functionality when you know the behavior will not extend beyond a
  single context.