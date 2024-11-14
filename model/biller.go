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
	Name            string `json:"name"`
}

// Struktur untuk response list biller dari API eksternal
// type ListBillerResponse struct {
// 	Data []Biller `json:"data"`
// }
