package dto

import (
	"crypto/md5"
	"encoding/base64"
)

type RequestBase struct {
	Checksum string
}

func (r RequestBase) GetChecksum() string {
	return r.Checksum
}

func ValidateChecksum(r Requester) bool {
	checksum := GetHash(r)
	return r.GetChecksum() == checksum
}

func GetHash(r Requester) string {
	return CreateHash(r.Key())
}

func CreateHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
