package shared

import "encoding/base64"

func GetValueStr(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func GetPointerStr(s string) *string {
	return &s
}

func EncodeBase64(cursor []byte) string {
	return base64.StdEncoding.EncodeToString(cursor)
}

func DecodeBase64(cursor string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
