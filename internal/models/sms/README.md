# SMS Configuration Domain Design (DDD-Oriented)

This document describes the domain model design of an SMS configuration system following Domain-Driven Design (DDD)
principles.

---

## 🧠 Aggregate Root: `Config`

### Description:

The `Config` struct is the **aggregate root** of the SMS configuration domain. It encapsulates all required data to
define a full SMS provider configuration, including credentials, paths, type, and sender lines.

### Fields:

- `ID`: Unique identifier for this config.
- `Title`: A human-readable name for the configuration.
- `IsDefault`: Indicates whether this is the default config.
- `Priority`: Determines the order in which configs are selected.
- `Credentials`: Embedded struct for API credentials (see `Credential`).
- `Status`: Indicates whether the config is active, inactive, etc.
- `ProviderID`: Foreign key to the `Provider` aggregate.
- `SenderLineID`: Selected default sender line for this config.
- `Type`: Foreign key to a `Type` defining the use-case (e.g., OTP, marketing).
- `CreatedAt`, `UpdatedAt`: Timestamps for lifecycle tracking.

> **Note:** This struct is the primary entry point for interacting with the SMS sending configuration system.

---

## 🧩 Value Object: `Credential`

### Description:

Represents credentials needed for communicating with an external SMS provider.

### Fields:

- `APIKey`, `SecretKey`, `Username`, `Password`: Credentials required for authentication.

> Since these are tightly coupled to a provider's requirements and don't carry identity, they are implemented as a *
*Value Object**.

---

## 📦 Entity Or (Aggregate Root): `Provider`

### Description:

Defines an external SMS provider. It is a separate aggregate but referenced by `Config`.

### Fields:

- `ID`: Unique identifier for the provider.
- `Name`: Friendly name (e.g., Twilio, Kavenegar).
- `Slug`: URL-safe identifier.
- `Description`: Optional details.
- `Website`: Provider's official URL.
- `APIPaths`: Value object holding API endpoint paths.
- `Status`: Indicates provider availability.
- `SenderLines`: A list of sender line IDs owned by this provider.
- `CreatedAt`, `UpdatedAt`: Timestamps.

> This entity may live in its own bounded context if providers are managed independently.

---

## 🧩 Value Object: `APIPath`

### Description:

Holds all relevant API endpoints a provider supports.

### Fields:

- `BaseURL`: Root URL of the provider’s API.
- Optional endpoints like `SendAPIPath`, `ReportAPIPath`, etc.

> Immutable within the lifecycle of the `Provider`. Treated as a **Value Object**.

---

## 📦 Entity: `SenderLine`

### Description:

Represents a single SMS sender number provided by an SMS vendor.

### Fields:

- `ID`: Unique identifier.
- `Number`: Phone number used for sending SMS.
- `Capacity`: Optional quota information.
- `IsActive`: Indicates whether it's usable.
- `Description`: Optional.
- `ProviderID`: Links to the owning `Provider`.

> Can belong to multiple configs or contexts, depending on design. If SenderLines are shared, they may be managed in a
> separate module.

---

## 📦 Entity Or (Aggregate Root): `Type`

### Description:

Describes the intent or use-case of a config (e.g., OTP, marketing, etc).

### Fields:

- `ID`: Unique identifier.
- `Name`: Short name (`otp`, `marketing`, `report`, etc.).
- `Description`: Clarifies usage.

> Useful for enforcing business rules, e.g., "OTP must have response validation."

---

## 🔄 Relationships

- `Config` → `Provider`: Many-to-One (via `ProviderID`)
- `Config` → `SenderLine`: One-to-One/Optional (via `SenderLineID`)
- `Config` → `Type`: One-to-One (via `Type`)
- `Provider` → `SenderLines`: One-to-Many

---

## ✅ DDD Principles Applied

| Principle              | Implementation Example                           |
|------------------------|--------------------------------------------------|
| Aggregate Roots        | `Config`, `Provider`, `Type`                     |
| Value Objects          | `Credential`, `APIPath`                          |
| Entities with Identity | `Config`, `Provider`, `SenderLine`, `Type`       |
| Ubiquitous Language    | `SenderLine`, `Config`, `Credential`, `Provider` |
| Separation of Concerns | Each struct focuses on a single responsibility   |

---

## 📌 Future Improvements

- Introduce domain services for logic like validating sender line uniqueness or default status rules.
- Enforce invariants like "Only one default config per type".
- Use domain events (e.g., `ConfigCreated`, `ProviderStatusChanged`).
