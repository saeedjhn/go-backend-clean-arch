package main

import (
	"log"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
)

const (
	_orderCreated   entity.StateFSM = "order_created"
	_paymentDone    entity.StateFSM = "payment_completed"
	_stockConfirmed entity.StateFSM = "stock_confirmed"
	_shipped        entity.StateFSM = "shipped"
	_delivered      entity.StateFSM = "delivered"
)

const (
	_makePayment  entity.EventFSM = "make_payment"
	_confirmStock entity.EventFSM = "confirm_stock"
	_shipOrder    entity.EventFSM = "ship_order"
	_deliverOrder entity.EventFSM = "deliver_order"
)

func main() {
	transitions := map[entity.StateFSM]map[entity.EventFSM]entity.StateFSM{
		_orderCreated:   {_makePayment: _paymentDone},
		_paymentDone:    {_confirmStock: _stockConfirmed},
		_stockConfirmed: {_shipOrder: _shipped},
		_shipped:        {_deliverOrder: _delivered},
	}

	fsm := entity.NewFSM(_orderCreated, transitions)

	events := []entity.EventFSM{_makePayment, _confirmStock, _shipOrder, _deliverOrder}

	for _, event := range events {
		if err := fsm.ApplyEventFSM(event); err != nil {
			log.Println("Error:", err)
			break
		}
	}

	log.Println("Final State:", fsm.GetCurrentState())
}
