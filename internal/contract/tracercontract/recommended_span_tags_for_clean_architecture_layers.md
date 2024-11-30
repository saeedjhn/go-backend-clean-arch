### Recommended Span Tags for Clean Architecture Layers

Using appropriate tags in spans helps provide detailed insights into the system's performance, errors, and specific
operation details. Below are the suggested tags for each layer in a Clean Architecture setup, with additional
considerations for better observability.

1. #### Delivery Layer (HTTP / gRPC / Messaging Layer)
   This layer manages requests and responses. Tags here should focus on input/output details.
    - Recommended Tags:
        - `http.method`: HTTP method (e.g., GET, POST, PUT, DELETE).
        - `http.url`: Request URL.
        - `http.status_code`: Response status code (e.g., 200, 404, 500).
        - `http.route`: Route matched for the request (e.g., /users/:id).
        - `http.client_ip`: Client's IP address.
        - `http.user_agent`: User agent string from the client.
        - `rpc.method (for gRPC)`: The invoked method name (e.g., UserService.GetUser).
        - `message.queue.name` (for messaging): Queue name for the message.
        - `delivery.protocol`: Protocol used (e.g., HTTP/1.1, HTTP/2, gRPC).
        - `delivery.duration_ms`: Time taken for request processing in this layer.


2. #### Usecase Layer
   This layer handles the core business logic. Tags here highlight the operation being performed and its dependencies.
    - Recommended Tags:
        - `usecase.name`: Name of the business operation (e.g., GetUser, CreateOrder).
        - `usecase.execution_time: Execution time for the use case (if measured manually).
        - `dependency.name`: Name of any external dependencies used (e.g., PostgreSQL, Redis).
        - `usecase.parameters`: Input parameters (ensure no sensitive data is logged).
        - `usecase.result`: Result of the operation (e.g., success, failure).
        - `usecase.failure_reason`: Reason for failure if applicable.


3. #### Repository Layer
   This layer interacts with the database or external services. Tags should provide details about queries or external
   service calls.
    - Recommended Tags:
        - `db.system`: Database type (e.g., postgresql, mysql, redis).
        - `db.statement`: Executed query (redact sensitive details).
        - `db.operation`: Type of database operation (e.g., SELECT, INSERT, UPDATE, DELETE).
        - `db.table`: Database table involved.
        - `db.rows_affected`: Number of rows affected (if applicable).
        - `db.duration_ms`: Query execution time (if measured manually).
        - `external.service.name`: Name of the external service (if applicable).
        - `external.service.endpoint`: Endpoint of the external service.
        - `external.service.status_code`: Response code from the service.


4. #### Middleware Layer
   Middleware applies common functionalities like authentication and logging. Tags here track shared logic applied to
   requests.
    - Recommended Tags:
        - `auth.user_id`: ID of the authenticated user.
        - `auth.roles`: User roles (e.g., admin, user).
        - `request.id`: Unique request identifier (if available).
        - `middleware.name`: Middleware name (e.g., TracingMiddleware, AuthMiddleware).
        - `auth.result`: Authentication result (success, failure).
        - `middleware.execution_time`: Time taken by middleware.


5. #### External API Layer
   For calls to external APIs or services, tags capture API-specific details.
    - Recommended Tags:
        - `external.api.name: Name of the external API (e.g., Stripe, SendGrid).
        - `external.api.endpoint: Endpoint called (e.g., /v1/payments).
        - `external.api.method: HTTP method used (e.g., POST, GET).
        - `external.api.response_time_ms: API response time.
        - `external.api.status_code: Status code of the response.