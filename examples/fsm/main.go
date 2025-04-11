package main

import (
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

const (
	_orderCreated   models.StateFSM = "order_created"
	_paymentDone    models.StateFSM = "payment_completed"
	_stockConfirmed models.StateFSM = "stock_confirmed"
	_shipped        models.StateFSM = "shipped"
	_delivered      models.StateFSM = "delivered"
)

const (
	_makePayment  models.EventFSM = "make_payment"
	_confirmStock models.EventFSM = "confirm_stock"
	_shipOrder    models.EventFSM = "ship_order"
	_deliverOrder models.EventFSM = "deliver_order"
)

func main() {
	transitions := map[models.StateFSM]map[models.EventFSM]models.StateFSM{
		_orderCreated:   {_makePayment: _paymentDone},
		_paymentDone:    {_confirmStock: _stockConfirmed},
		_stockConfirmed: {_shipOrder: _shipped},
		_shipped:        {_deliverOrder: _delivered},
	}

	fsm := models.NewFSM(_orderCreated, transitions)

	events := []models.EventFSM{_makePayment, _confirmStock, _shipOrder, _deliverOrder}

	for _, event := range events {
		if err := fsm.ApplyEventFSM(event); err != nil {
			log.Println("Error:", err)
			break
		}
	}

	log.Println("Final State:", fsm.GetCurrentState())
}
