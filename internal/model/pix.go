package model

type Pix struct {
	KeyType string `bson:"key_type"`
	Key     string `bson:"key"`
}
