package Entity

import "github.com/google/uuid"

type USDBRL struct {
	ID          string
	Code        string
	Codein      string
	Name        string
	High        string
	Low         string
	VarBid      string
	PctChange   string
	Bid         string
	Ask         string
	Timestamp   string
	Create_date string
}

func NewUSDBRL(
	Code string,
	Codein string,
	Name string,
	High string,
	Low string,
	VarBid string,
	PctChange string,
	Bid string,
	Ask string,
	Timestamp string,
	Create_date string,
) *USDBRL {
	return &USDBRL{
		ID:          uuid.New().String(),
		Code:        Code,
		Codein:      Codein,
		Name:        Name,
		High:        High,
		Low:         Low,
		VarBid:      VarBid,
		PctChange:   PctChange,
		Bid:         Bid,
		Ask:         Ask,
		Timestamp:   Timestamp,
		Create_date: Create_date,
	}
}
