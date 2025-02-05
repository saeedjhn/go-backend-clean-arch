# Outbox Pattern Configuration (`outbox.Config`)

The `outbox.Config` struct is used to configure the processing of events in the **Outbox Pattern**. Below is an
explanation of each field and how it impacts the behavior of the Outbox processing.

---

## Fields

### 1. **`IntervalInSeconds`**

- **Type**: `time.Duration`
- **Default**: `time.Second`
- **Description**:  
  Specifies how often the Outbox processing should run.  
  For example, `time.Second` means the `StartProcessing` function will be executed every **1 second**.  
  This defines how frequently the system checks for unpublished messages.

---

### 2. **`BatchSize`**

- **Type**: `int`
- **Default**: `2`
- **Description**:  
  Determines the number of messages fetched from the database in each processing cycle.  
  A value of `2` means that at most **two outbox events** will be processed per execution.  
  If more messages exist in the queue, they will be processed in the next execution.

---

### 3. **`RetryThreshold`**

- **Type**: `int`
- **Default**: `3`
- **Description**:  
  Defines the maximum retry attempts for failed message delivery.  
  A value of `3` means that if an event fails to publish **three times**, it will no longer be retried.  
  This prevents wasting resources on repeatedly failing messages.

---

## Summary

- **✅ Processing runs every 1 second.**
- **✅ Up to 2 messages are processed per execution.**
- **✅ Messages failing 3 times will no longer be retried.**

This configuration ensures that events are processed **efficiently** while avoiding unnecessary retries for problematic
messages. 🚀

---

## Example Configuration

Here’s an example of how to configure the `outbox.Config` struct:

```go
config := outbox.Config{
    IntervalInSeconds: 5 * time.Second, // Process every 5 seconds
    BatchSize:         10,              // Process up to 10 messages per cycle
    RetryThreshold:    5,               // Retry failed messages up to 5 times
}