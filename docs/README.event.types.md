# Event Types in Domain-Driven Design (DDD)

## Overview

Event types define the structure and naming conventions for domain events in a system. These events capture significant
occurrences within a **bounded context** and help in decoupling different parts of the system.

### Recommended Naming Conventions

1. `{bounded_context}.{aggregate}.{event_name}.{subcategory}`
2. `{bounded_context}.{aggregate}.{event_name}`

Using a hierarchical structure allows flexibility while keeping events understandable and structured.

## Examples of Event Types

| Event Type                             | Description                                       |
|----------------------------------------|---------------------------------------------------|
| `users.account.created`                | A new user account was created                    |
| `users.account.updated.profile`        | A user updated their profile information          |
| `orders.payment.completed`             | A payment for an order was successfully completed |
| `orders.shipment.dispatched`           | An order was dispatched for delivery              |
| `auth.sessions.token.generated`        | A new authentication token was generated          |
| `auth.sessions.token.revoked`          | A user session token was revoked                  |
| `inventory.products.stock.depleted`    | A product's stock level reached zero              |
| `inventory.products.stock.replenished` | A product was restocked                           |
| `notifications.email.sent`             | An email notification was sent to a user          |
| `notifications.sms.failed`             | An SMS notification failed to send                |

## Summary

- **Event types help structure domain events** in a meaningful and scalable way.
- **Hierarchical namespacing** ensures clarity and future extensibility.
- **Consistent naming conventions** improve system observability and traceability.
