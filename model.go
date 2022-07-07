package main

type Request struct {
	ID            int         `json:"ID"`
	Amount        float32     `json:"Amount"`
	Currency      string      `json:"Currency"`
	CustomerEmail string      `json:"CustomerEmail"`
	SplitInfo     []SplitInfo `json:"SplitInfo"`
}

type SplitInfo struct {
	SplitType     string  `json:"SplitType"`
	SplitValue    float32 `json:"SplitValue"`
	SplitEntityID string  `json:"SplitEntityID"`
}

type Response struct {
	ID             int              `json:"ID"`
	Balance        float32          `json:"Balance"`
	SplitBreakDown []SplitBreakDown `json:"SplitBreakDown"`
}

type SplitBreakDown struct {
	SplitEntityID string  `json:"SplitEntityID"`
	Amount        float32 `json:"Amount"`
}

type BalanceTracker struct {
	Balance          float32
	RatioOpenBalance struct {
		Balance float32
		Set     bool
		Total   float32
	}
}
