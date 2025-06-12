package dto

type Colmeia struct {
	ID           *int    `json:"id"`
	ColmeiaID    *int    `json:"colmeia_id"`
	QRCode       *string `json:"qr_code"`
	Species      *string `json:"species"`
	StartingDate *string `json:"starting_date"`
	Status       *string `json:"status"`
}
