package application

import "github.com/dayvd/matching-engine/internal/domain"

type CommandType uint8

const (
	CmdNewOrder CommandType = iota
	CmdCancel
	CmdAmend
)

type Command struct {
	Type CommandType

	Order domain.Order

	TargetID domain.OrderID

	NewPrice domain.Price
	NewQty   domain.Qty
}
