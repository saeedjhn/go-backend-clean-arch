package kavenegar

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Name() contract.Name {
	// TODO implement me
	panic("implement me")
}

func (c *Client) WithEndpoint(_ string) contract.SMSSender {
	// TODO implement me
	panic("implement me")
}

func (c *Client) WithCredentials(_, _ string) contract.SMSSender {
	// TODO implement me, signature: username, password
	panic("implement me")
}

func (c *Client) WithAPIKey(_ string) contract.SMSSender {
	// TODO implement me
	panic("implement me")
}

func (c *Client) CallSendSingleAPI(_, _, _ string, _ bool) error {
	// TODO implement me, signature: to, from, text, isFlash
	panic("implement me")
}

func (c *Client) CallSendBulkAPI(_ []string, _, _ string, _ bool) error {
	// TODO implement me, signature: to, from, text, isFlash
	panic("implement me")
}

func (c *Client) CallSendByBaseNumber(_, _, _ string) error {
	// TODO implement me, signature: text, to, bodyID
	panic("implement me")
}

// func (c *Client) FetchDeliveryStatus(_ string) (string, error) {
// 	// TODO implement me, signature: recId
// 	panic("implement me")
// }

// func (c *Client) FetchBatchDeliveryStatuses(recIds []string) ([]string, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) FetchMessages(location, from string, index, count int) ([]contract.Message, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) FetchMessagesByDate(
// location, from string, index, count int, dateFrom, dateTo time.Time)
// ([]contract.Message, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) FetchUserMessagesByDate(
// location, from string, index, count int, dateFrom, dateTo time.Time
// ) ([]contract.Message, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) GetCredit() (float64, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) GetBasePrice() (float64, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) GetSmsPrice(from, text string, irancellCount, mtnCount int) (float64, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) GetUserNumbers() ([]string, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) GetInboxCount(isRead bool) (int, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) CallSendAdvanced(
// to []string, from, text string, isFlash bool, udh string, recIds, statuses []string
// ) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) FetchMessageReception(msgId string, fromRows int) (contract.MessageDetail, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (c *Client) RemoveMessages(location string, msgIds []string) error {
// 	// TODO implement me
// 	panic("implement me")
// }
