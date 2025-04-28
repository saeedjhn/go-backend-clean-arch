package kavenegar

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/kavenegar/kavenegar-go"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type Config struct {
	Provider string `mapstructure:"Provider"`
	Sender   string `mapstructure:"sender"`
	APIKey   string `mapstructure:"api_key"`
}

type Client struct {
	config   Config
	Provider *kavenegar.Kavenegar
}

func New(config Config) *Client {
	return &Client{config: config, Provider: kavenegar.New(config.APIKey)}
}

func (c *Client) SendSingle(receptor string, message string) (int, error) {
	// var params *kavenegar.MessageSendParam
	res, err := c.Provider.Message.Send(c.config.Sender, []string{receptor}, message, nil)
	if err != nil {
		return 0, c.errorHandling(err)
	}

	if len(res) == 0 {
		return 0, ErrEmptyRecipient
	}

	return res[0].MessageID, nil
}

func (c *Client) SendSingleAt(receptor string, message string, duration time.Duration) (int, error) {
	params := &kavenegar.MessageSendParam{
		Date: time.Now().Add(duration),
	}
	res, err := c.Provider.Message.Send(c.config.Sender, []string{receptor}, message, params)
	if err != nil {
		return 0, c.errorHandling(err)
	}

	if len(res) == 0 {
		return 0, ErrEmptyRecipient
	}

	return res[0].MessageID, nil
}

func (c *Client) SendBulk(receptors []string, message string) ([]int, error) {
	// var params *kavenegar.MessageSendParam
	res, err := c.Provider.Message.Send(c.config.Sender, receptors, message, nil)
	if err != nil {
		return nil, c.errorHandling(err)
	}

	if len(res) == 0 {
		return nil, ErrEmptyRecipient
	}

	recIDs := c.extractMessageIDs(res)

	return recIDs, nil
}

func (c *Client) GetStatus(recIDs []int) ([]contract.Status, error) {
	status, err := c.Provider.Message.Status(c.intsToStrings(recIDs))
	if err != nil {
		return nil, c.errorHandling(err)
	}

	es := c.extractStatus(status)

	return es, nil
}

func (c *Client) extractMessageIDs(res []kavenegar.Message) []int {
	var recIDs []int
	for _, re := range res {
		recIDs = append(recIDs, re.MessageID)
	}

	return recIDs
}

func (c *Client) extractStatus(status []kavenegar.MessageStatus) []contract.Status {
	var s []contract.Status
	for _, messageStatus := range status {
		s = append(s, contract.Status{
			RecID:      messageStatus.MessageId,
			StatusCode: messageStatus.Status,
			StatusText: messageStatus.StatusText,
		})
	}

	return s
}

func (c *Client) intsToStrings(nums []int) []string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	return strs
}

func (c *Client) errorHandling(err error) error {
	var apiErr *kavenegar.APIError
	var httpErr *kavenegar.HTTPError

	switch {
	case errors.As(err, &apiErr):
		return fmt.Errorf("kavenegar API error: %w", apiErr)
	case errors.As(err, &httpErr):
		return fmt.Errorf("kavenegar HTTP error: %w", httpErr)
	default:
		return fmt.Errorf("unexpected error from kavenegar: %w", err)
	}
}
