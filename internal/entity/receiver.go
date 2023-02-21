package entity

type Receiver struct {
	ID         string
	Identifier string
	Name       string
	Email      string
	Pix        Pix
	Bank       *string
	Agency     *string
	Account    *string
	Status     Status
}
