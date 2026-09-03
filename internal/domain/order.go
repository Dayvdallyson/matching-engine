package domain

type Side uint8

const (
	Buy Side = iota
	Sell
)

func (s Side) String() string {
	switch s {
	case Buy:
		return "BUY"
	case Sell:
		return "SELL"
	default:
		return "UNKNOWN_SIDE"
	}
}

func (s Side) Opposite() Side {
	if s == Buy {
		return Sell
	}
	return Buy
}

type OrderType uint8

const (
	Limit OrderType = iota
	Market
)

func (t OrderType) String() string {
	switch t {
	case Limit:
		return "LIMIT"
	case Market:
		return "MARKET"
	default:
		return "UNKNOWN_TYPE"
	}
}

type TimeInForce uint8

const (
	GTC TimeInForce = iota
	IOC
	FOK
)

func (tif TimeInForce) String() string {
	switch tif {
	case GTC:
		return "GTC"
	case IOC:
		return "IOC"
	case FOK:
		return "FOK"
	default:
		return "UNKNOWN_TIF"
	}
}

type OrderStatus uint8

const (
	New OrderStatus = iota
	PartiallyFilled
	Filled
	Cancelled
	Rejected
)

func (st OrderStatus) String() string {
	switch st {
	case New:
		return "NEW"
	case PartiallyFilled:
		return "PARTIALLY_FILLED"
	case Filled:
		return "FILLED"
	case Cancelled:
		return "CANCELLED"
	case Rejected:
		return "REJECTED"
	default:
		return "UNKNOWN_STATUS"
	}
}

type Order struct {
	ID       OrderID
	Seq      SeqNum
	Price    Price
	Qty      Qty
	Remaining Qty

	Side   Side
	Type   OrderType
	TIF    TimeInForce
	Status OrderStatus
}

func (o *Order) IsFilled() bool {
	return o.Remaining == 0
}

func (o *Order) FilledQty() Qty {
	return o.Qty - o.Remaining
}
