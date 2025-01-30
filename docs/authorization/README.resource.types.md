# Resource Types and Examples

## **1️⃣ "module" (Module)**

📌 **Description:** A specific system module that requires access control.

🔹 **Examples:**

| ID | Name             | Type   | Description                     |
|----|------------------|--------|---------------------------------|
| 1  | User Management  | module | Manages users and roles         |
| 2  | Order Processing | module | Handles orders and transactions |
| 3  | Reporting        | module | Generates reports and analytics |

---

## **2️⃣ "file" (File)**

📌 **Description:** A protected file that requires restricted access.

🔹 **Examples:**

| ID | Name         | Type | Description               |
|----|--------------|------|---------------------------|
| 4  | invoices.pdf | file | Financial invoices        |
| 5  | config.yaml  | file | Server configuration file |
| 6  | logo.png     | file | Company logo              |

---

## **3️⃣ "API" (API Endpoint)**

📌 **Description:** An API endpoint that needs role-based access control.

🔹 **Examples:**

| ID | Name             | Type | Description                       |
|----|------------------|------|-----------------------------------|
| 7  | GET /users       | API  | Fetches the list of users         |
| 8  | POST /orders     | API  | Creates a new order               |
| 9  | DELETE /products | API  | Deletes a product from the system |

---

## **Other Possible Types**

If needed, we can extend the `Type` field with additional categories such as:

- **"database"** → Represents access to database tables (e.g., `users`, `orders`).
- **"service"** → Represents a microservice in the system (e.g., `payment-service`).
- **"queue"** → Represents a message queue that needs permission control (e.g., `order-events-queue`).

