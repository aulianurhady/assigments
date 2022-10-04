package transports

import (
	"time"
)

type Request struct {
	CustomerName  string    `json:"costumerName"`
	OrderedAt     time.Time `json:"orderedAt"`
	CustomerItems []Item    `json:"items"`
}

type Item struct {
	LineItemID  int    `json:"lineItemID,omitempty"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
