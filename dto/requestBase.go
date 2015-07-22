package dto

import (
	"crypto/md5"
	"encoding/base64"
)

type RequestBase struct {
	Checksum string
}

func (r RequestBase) Key() string{
	return "testing"
}

func (r RequestBase) ValidateChecksum() bool {
	checksum := r.GetHash()
	return r.Checksum == checksum
}

func (r RequestBase) GetHash() string {
	return CreateHash(r.Key())
}

func CreateHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
