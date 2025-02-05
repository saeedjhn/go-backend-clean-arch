# Outbox Pattern in Event-Driven Architecture

## Overview

The **Outbox Pattern** ensures reliable event publishing in event-driven architectures by storing events in a database table before publishing them to a messaging system. This approach guarantees consistency between the database and the messaging system, reducing the risk of message loss.

---

## How It Works

### 1. Storing Events in the Database

When a business operation occurs (e.g., creating an order), an event is stored in the `outbox_events` table instead of being directly published to the messaging system. This ensures that the event is persisted in the database as part of the same transaction as the business operation.

### 2. Executing a Background Processor

A separate **scheduler** periodically scans the `outbox_events` table for unpublished events. This scheduler is responsible for fetching and publishing events to the messaging system.

### 3. Publishing Events to a Messaging System

Unpublished events are fetched from the `outbox_events` table and sent to a messaging system such as **Kafka** or **RabbitMQ**. This step ensures that events are eventually delivered to consumers.

### 4. Updating Event Status

If an event is successfully published, the `is_published` field in the `outbox_events` table is updated to `TRUE`. This prevents duplicate processing of the same event.

---

## Benefits of the Outbox Pattern

- **Guaranteed Message Delivery**: Ensures that messages are eventually published, even if the service crashes before publishing.
- **Transactional Messaging**: Maintains consistency between the database and the messaging system by using database transactions.
- **Improved Reliability**: Prevents lost messages in microservices architectures, ensuring fault tolerance.

---

## Implementation Details

This package provides an implementation of the Outbox Pattern in Go, featuring:

### 1. **Scheduler**
- Uses `gocron` to periodically fetch and publish events from the `outbox_events` table.
- Ensures that events are processed in a timely manner.

### 2. **Repository**
- Interfaces with **PostgreSQL** to retrieve and update events in the `outbox_events` table.
- Provides methods for inserting new events and marking events as published.

### 3. **Publisher**
- Sends events to an external messaging system (e.g., Kafka, RabbitMQ).
- Handles retries and error handling to ensure reliable message delivery.

---

## Example Workflow

1. **Business Operation**:
    - A user places an order, and an `OrderCreated` event is generated.
    - The event is stored in the `outbox_events` table as part of the same database transaction.

2. **Background Processor**:
    - The scheduler periodically scans the `outbox_events` table for unpublished events.
    - It fetches the `OrderCreated` event and sends it to the messaging system.

3. **Event Publishing**:
    - The event is successfully published to Kafka.
    - The `is_published` field in the `outbox_events` table is updated to `TRUE`.

---