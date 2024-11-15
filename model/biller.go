package model

// Struktur untuk akun biller
type BillerAccount struct {
	BillerAccountID string `json:"biller_account_id"`
	Name            string `json:"name"`
	BillAmount      int64  `json:"bill_amount"`
	Paid            bool   `json:"paid"`
}

// Struktur untuk Biller yang memiliki beberapa akun
type Biller struct {
	BillerID string                   `json:"biller_id"`
	Name     string                   `json:"name"`
	Accounts map[string]BillerAccount `json:"accounts"` // Map untuk menyimpan akun berdasarkan ID
}

type PayBillerRequest struct {
	BillerID        string `json:"biller_id"`
	BillerAccountID string `json:"biller_account_id"`
	Amount          int64  `json:"amount"`
}

type BillerAccountData struct {
	BillerID        string `json:"biller_id"`
	BillerName      string `json:"biller_name"`
	BillerAccountID string `json:"biller_account_id"`
	Name            string `json:"name"`
	BillAmount      int    `json:"bill_amount"`
	Paid            bool   `json:"paid"`
}

type BillerAccountResponse struct {
	Message string            `json:"message"`
	Data    BillerAccountData `json:"data"`
}
