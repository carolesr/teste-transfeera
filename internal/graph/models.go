package graph

type NewReceiver struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	PixKeyType string `json:"pixKeyType"`
	PixKey     string `json:"pixKey"`
}

type Pix struct {
	KeyType string `json:"keyType"`
	Key     string `json:"key"`
}

type Receiver struct {
	ID         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Pix        *Pix    `json:"pix"`
	Bank       *string `json:"bank"`
	Agency     *string `json:"agency"`
	Account    *string `json:"account"`
	Status     *string `json:"status"`
}
