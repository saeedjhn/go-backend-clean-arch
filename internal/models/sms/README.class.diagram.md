classDiagram
    class Config {
        +ID: types.ID
        +Title: string
        +IsDefault: bool
        +Priority: int
        +Template: Template
        +Status: Status
        +ProviderID: types.ID
        +SenderLineID: types.ID
        +Type: types.ID
        +CreatedAt: time.Time
        +UpdatedAt: time.Time
    }

    class Template {
        +Content: map[Lang]string
        -lang: Lang
        +GetLocalizedContent(Lang) string
    }

    class Provider {
        +ID: types.ID
        +Name: ProviderName
        +Slug: string
        +Description: string
        +Website: string
        +APIPaths: APIPath
        +Credentials: Credential
        +Status: Status
        +SenderLines: types.ID[]
        +CreatedAt: time.Time
        +UpdatedAt: time.Time
    }

    class Credential {
        +APIKey: string
        +SecretKey: string
        +Username: string
        +Password: string
    }

    class APIPath {
        +BaseURL: string
        +Path: string
    }

    class SenderLine {
        +ID: types.ID
        +Number: string
        +Capacity: int
        +IsActive: bool
        +Description: string
        +ProviderID: types.ID
        +CreatedAt: time.Time
        +UpdatedAt: time.Time
    }

    class Type {
        +ID: types.ID
        +Name: TypeName
        +Description: string
        +CreatedAt: time.Time
        +UpdatedAt: time.Time
    }

    %% Enum Definitions
    class Lang {
        <<enumeration>>
        FA
        EN
        AR
        Default
    }

    class Status {
        <<enumeration>>
        Active
        Inactive
    }

    class ProviderName {
        <<enumeration>>
        MelliPayamak
        Kavenegar
        SMSIR
        FarazSMS
        Payamito
        Farapayamak
        Unknown
    }

    class TypeName {
        <<enumeration>>
        OTP
        Report
    }

    %% Relationships
    Config *-- Template
    Config --> Provider
    Config --> SenderLine
    Config --> Type
    Provider *-- Credential
    Provider *-- APIPath
    Provider "1" --> "many" SenderLine
    Template ..> Lang
    Type --> TypeName
