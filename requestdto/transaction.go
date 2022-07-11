package requestdto

type TransactionSplitDetails struct {
	UserID     string  `json:"userID,omitempty"`
	PaidAmount float64 `json:"paidAmount,omitempty"`
	OwesAmount float64 `json:"owesAmount,omitempty"`
}

type TransactionRequest struct {
	ID           string                    `json:"id,omitempty"`
	Amount       float64                   `json:"amount,omitempty"`
	Comment      string                    `json:"comment,omitempty"`
	SplitDetails []TransactionSplitDetails `json:"splitDetails,omitempty"`
	GroupID      string                    `json:"groupID,omitempty"`
}
