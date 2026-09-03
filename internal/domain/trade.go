package domain

type Trade struct {
	Seq        SeqNum
	Price      Price
	Qty        Qty
	Aggressor  OrderID
	Resting    OrderID
	AggrSide   Side
}
