package entity

type Receiver struct {
	ID         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	PixKeyType string  `json:"pix_key_type"`
	PixKey     string  `json:"pix_key"`
	Bank       *string `json:"bank"`
	Agency     *string `json:"agency"`
	Account    *string `json:"account"`
	Status     *string `json:"status"`
}

type NewReceiver struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	PixKeyType string `json:"pix_key_type"`
	PixKey     string `json:"pix_key"`
}
