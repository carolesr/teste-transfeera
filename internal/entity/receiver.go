package entity

type Receiver struct {
	ID         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Pix        Pix     `json:"pix"`
	Bank       *string `json:"bank"`
	Agency     *string `json:"agency"`
	Account    *string `json:"account"`
	Status     Status  `json:"status"`
}
