package entity

type PixKeyType string

const (
	CPF       PixKeyType = "CPF"
	CNPJ      PixKeyType = "CNPJ"
	Email     PixKeyType = "Email"
	Phone     PixKeyType = "Phone"
	RandomKey PixKeyType = "Random Key"
)

type Pix struct {
	KeyType PixKeyType `json:"key_type"`
	Key     string     `json:"key"`
}
