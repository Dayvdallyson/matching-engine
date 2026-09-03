package application

import "github.com/dayvd/matching-engine/internal/domain"

type EventType uint8

const (
	EvAccepted EventType = iota
	EvTrade
	EvBookDelta
	EvCancelled
	EvAmended
	EvRejected
)

type RejectReason uint8

const (
	RejNone RejectReason = iota
	RejUnknownOrder
	RejFOKUnfillable
	RejPriceOutOfBand
	RejZeroQty
)

type Event struct {
	Type EventType
	Seq  domain.SeqNum

	OrderID domain.OrderID
	Status  domain.OrderStatus

	Trade domain.Trade

	Side     domain.Side
	Price    domain.Price
	LevelQty domain.Qty

	Reason RejectReason
}
