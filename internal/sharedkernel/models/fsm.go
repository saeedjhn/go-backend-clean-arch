package models

import (
	"fmt"
)

type EventFSM string

type StateFSM string

type FSM struct {
	CurrentState StateFSM
	Transitions  map[StateFSM]map[EventFSM]StateFSM
}

func NewFSM(initialState StateFSM, transitions map[StateFSM]map[EventFSM]StateFSM) *FSM {
	return &FSM{
		CurrentState: initialState,
		Transitions:  transitions,
	}
}

func (fsm *FSM) ApplyEventFSM(event EventFSM) error {
	if next, ok := fsm.Transitions[fsm.CurrentState][event]; ok {
		fsm.CurrentState = next
		return nil
	}

	return fmt.Errorf("cannot apply event %q in state %q", event, fsm.CurrentState)
}

func (fsm *FSM) GetCurrentState() StateFSM {
	return fsm.CurrentState
}

// Example
// func OrderFSM() *fsm.FSM {
// 	return fsm.NewFSM(
// 		fsm.OrderCreated,
// 		map[fsm.OrderState]map[fsm.Event]fsm.OrderState{
// 			fsm.OrderCreated:         {fsm.PayOrder: fsm.PaymentPending},
// 			fsm.PaymentPending:       {fsm.StartManufacture: fsm.ManufacturingPending},
// 			fsm.ManufacturingPending: {fsm.StockOrder: fsm.StockPending},
// 			fsm.StockPending:         {fsm.DispatchOrder: fsm.DispatchPending},
// 			fsm.DispatchPending:      {fsm.DeliverOrder: fsm.DeliveryPending},
// 			fsm.DeliveryPending:      {fsm.ConfirmOrder: fsm.ConfirmationPending},
// 			fsm.ConfirmationPending:  {fsm.ConfirmOrder: fsm.OrderConfirmed},
// 		},
// 	)
// }

// type OrderService struct {
//	fsm fsm.StateMachine
// }
//
// func NewOrderService() *OrderService {
//	return &OrderService{
//		fsm: usecase.OrderFSM(),
//	}
// }
//
// func (s *OrderService) HandleEvent(event fsm.Event) {
//	err := s.fsm.ApplyEvent(event)
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Order state changed to:", s.fsm.GetCurrentState())
//	}
// }

// Way change state
// 1. API (REST/GraphQL/gRPC)
// /orders/{id}/status
// PATCH /orders/123/status
// Content-Type: application/json
// {
//    "event": "pay_order"
// }
// 2. (RabbitMQ, Kafka, NATS)
// 3. CRON Job
// 4. Webhook(other service)

// // State represents the current state of the order
// type State string
//
// const (
//	OrderCreated         State = "order_created"
//	PaymentPending       State = "payment_pending"
//	ManufacturingPending State = "manufacturing_pending"
//	StockMovePending     State = "stock_pending"
//	DispatchPending      State = "dispatch_pending"
//	DeliveryPending      State = "delivery_pending"
//	ConfirmationPending  State = "confirmation_pending"
//	OrderConfirmed       State = "order_confirmed"
//	RefundPending        State = "refund_pending"
//	OrderCancelled       State = "order_cancelled"
// )
//
// // Event represents an event that triggers a state transition
// type Event string
//
// const (
//	PaymentReceived   Event = "payment_received"
//	TableProduced     Event = "table_produced"
//	TableMovedToStock Event = "table_moved_to_stock"
//	TableDispatched   Event = "table_dispatched"
//	TableDelivered    Event = "table_delivered"
//	CustomerConfirmed Event = "customer_confirmed"
//	ManufacturingFail Event = "manufacturing_fail"
//	RefundCompleted   Event = "refund_completed"
// )
//
// // Transition represents a valid state transition
// type Transition struct {
//	From  State
//	Event Event
//	To    State
// }
//
// // ValidTransitions is a list of all valid transitions
// var ValidTransitions = []Transition{
//	{OrderCreated, PaymentReceived, PaymentPending},
//	{PaymentPending, TableProduced, ManufacturingPending},
//	{ManufacturingPending, TableMovedToStock, StockMovePending},
//	{StockMovePending, TableDispatched, DispatchPending},
//	{DispatchPending, TableDelivered, DeliveryPending},
//	{DeliveryPending, CustomerConfirmed, ConfirmationPending},
//	{ConfirmationPending, CustomerConfirmed, OrderConfirmed},
//	{ManufacturingPending, ManufacturingFail, RefundPending},
//	{RefundPending, RefundCompleted, OrderCancelled},
// }

// // TransitionState changes the state based on the current state and the event
// func TransitionState(currentState State, event Event) (State, error) {
//	for _, transition := range ValidTransitions {
//		if transition.From == currentState && transition.Event == event {
//			return transition.To, nil
//		}
//	}
//	return currentState, errors.New("invalid transition")
// }

// func main() {
//	// Initial state
//	currentState := OrderCreated
//
//	// Process events
//	events := []Event{PaymentReceived, TableProduced, ManufacturingFail, RefundCompleted}
//
//	for _, event := range events {
//		newState, err := TransitionState(currentState, event)
//		if err != nil {
//			fmt.Printf("Error: %v\n", err)
//			break
//		}
//		fmt.Printf("Transitioned from %s to %s on event %s\n", currentState, newState, event)
//		currentState = newState
//	}
// }
