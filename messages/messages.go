package messages

type ReceivePaymentRequest struct {
	Number      string  `json:"number"`
	Sender      string  `json:"sender"`
	Beneficiary string  `json:"beneficiary"`
	Amount      float32 `json:"amount"`
	Currency    string  `json:"currency"`
}

/*
type AcceptPaymentRequest struct {
	Number string `json:"number"`
}

type RejectPaymentRequest struct {
	Number string `json:"number"`
}
*/

type Payment struct {
	ID          int64   `json:"id"`
	Number      string  `json:"number"`
	Sender      string  `json:"sender"`
	Beneficiary string  `json:"beneficiary"`
	Amount      float32 `json:"amount"`
	Currency    string  `json:"currency"`
	CreatedAt   int64   `json:"created_at"`
	State       string  `json:"state"`
}
