# SMS Service - Clean Architecture Structure (Go)

This document outlines a complete Clean Architecture structure for implementing an SMS service (e.g. MelliPayamak) in
Go. The architecture includes:

- Delivery Layer (HTTP Handlers)
- Usecase Layer (Business Logic)
- Domain Layer (Interfaces)
- Repository Layer (External Systems Integration)

---

## 📁 Project Structure

```
internal/
├── domain/
│   └── sms/
│       └── interface.go        // SMSService interface
│
├── usecase/
│   └── sms/
│       └── usecase.go          // Usecase implementations
│
├── repository/
│   └── sms/
│       └── mellipayamak.go     // SMS repository implementation
│
└── delivery/
    └── http/
        └── sms_handler.go      // HTTP handlers
```

---

## 📜 Domain Interface (`domain/sms/interface.go`)

```go
type SMSService interface {
SendSMS(to, from, text string, isFlash bool) error
SendSimpleSMS(to []string, from, text string, isFlash bool) error
SendWithSharedServiceLine(text, to, bodyId string) error
GetDeliveryStatus(recId string) (string, error)
GetBatchDeliveryStatuses(recIds []string) ([]string, error)
ListMessages(location, from string, index, count int) ([]Message, error)
ListMessagesByDate(location, from string, index, count int, dateFrom, dateTo time.Time) ([]Message, error)
ListUserMessagesByDate(location, from string, index, count int, dateFrom, dateTo time.Time) ([]Message, error)
GetCreditBalance() (float64, error)
GetBaseTariff() (float64, error)
CalculateSMSPrice(from, text string, irancellCount, mtnCount int) (float64, error)
ListSenderNumbers() ([]string, error)
CountInboxMessages(isRead bool) (int, error)
SendAdvancedSMS(to []string, from, text string, isFlash bool, udh string, recIds, statuses []string) error
GetMessageDetails(msgId string, fromRows int) (MessageDetail, error)
DeleteInboxMessages(location string, msgIds []string) error
}
```

---

## 💼 Usecase Methods (`usecase/sms/usecase.go`)

```go
type smsUsecase struct {
repo sms.Repository
}

func (u *smsUsecase) SendSMS(...) error                       { ... }
func (u *smsUsecase) SendSimpleSMS(...) error                 { ... }
func (u *smsUsecase) SendWithSharedServiceLine(...) error    { ... }
func (u *smsUsecase) GetDeliveryStatus(...) (string, error)  { ... }
func (u *smsUsecase) GetBatchDeliveryStatuses(...) ([]string, error) { ... }
func (u *smsUsecase) ListMessages(...) ([]Message, error)    { ... }
func (u *smsUsecase) ListMessagesByDate(...) ([]Message, error) { ... }
func (u *smsUsecase) ListUserMessagesByDate(...) ([]Message, error) { ... }
func (u *smsUsecase) GetCreditBalance(...) (float64, error)  { ... }
func (u *smsUsecase) GetBaseTariff(...) (float64, error)     { ... }
func (u *smsUsecase) CalculateSMSPrice(...) (float64, error) { ... }
func (u *smsUsecase) ListSenderNumbers(...) ([]string, error) { ... }
func (u *smsUsecase) CountInboxMessages(...) (int, error)    { ... }
func (u *smsUsecase) SendAdvancedSMS(...) error              { ... }
func (u *smsUsecase) GetMessageDetails(...) (MessageDetail, error) { ... }
func (u *smsUsecase) DeleteInboxMessages(...) error          { ... }
```

---

## 🧱 Repository Interface (`repository/sms/interface.go`)

```go
type Repository interface {
CallSendAPI(to, from, text string, isFlash bool) error
CallSendSimpleSMS(to []string, from, text string, isFlash bool) error
CallSendByBaseNumber(text, to, bodyId string) error
FetchDeliveryStatus(recId string) (string, error)
FetchBatchDeliveryStatuses(recIds []string) ([]string, error)
FetchMessages(location, from string, index, count int) ([]Message, error)
FetchMessagesByDate(location, from string, index, count int, dateFrom, dateTo time.Time) ([]Message, error)
FetchUserMessagesByDate(location, from string, index, count int, dateFrom, dateTo time.Time) ([]Message, error)
GetCredit() (float64, error)
GetBasePrice() (float64, error)
GetSmsPrice(from, text string, irancellCount, mtnCount int) (float64, error)
GetUserNumbers() ([]string, error)
GetInboxCount(isRead bool) (int, error)
CallSendAdvanced(to []string, from, text string, isFlash bool, udh string, recIds, statuses []string) error
FetchMessageReception(msgId string, fromRows int) (MessageDetail, error)
RemoveMessages(location string, msgIds []string) error
}
```

---

## 🌐 RESTful API Endpoints

| HTTP Method | Endpoint                  | Description                    |
|-------------|---------------------------|--------------------------------|
| `POST`      | `/sms/send`               | Send basic SMS                 |
| `POST`      | `/sms/send/simple`        | Send multiple SMS              |
| `POST`      | `/sms/send/shared-line`   | Send using shared service line |
| `POST`      | `/sms/status`             | Get delivery status            |
| `POST`      | `/sms/status/batch`       | Get batch delivery statuses    |
| `GET`       | `/sms/messages`           | List messages                  |
| `GET`       | `/sms/messages/date`      | List messages by date          |
| `GET`       | `/sms/messages/user-date` | List user messages by date     |
| `GET`       | `/sms/credit`             | Get account balance            |
| `GET`       | `/sms/price/base`         | Get base price                 |
| `POST`      | `/sms/price/calculate`    | Calculate SMS price            |
| `GET`       | `/sms/numbers`            | Get sender numbers             |
| `GET`       | `/sms/inbox/count`        | Get inbox message count        |
| `POST`      | `/sms/send/advanced`      | Send advanced SMS              |
| `GET`       | `/sms/message/details`    | Get message details            |
| `DELETE`    | `/sms/inbox`              | Delete inbox messages          |