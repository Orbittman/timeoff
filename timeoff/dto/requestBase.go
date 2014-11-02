package dto

import (
	"crypto/md5"
	"encoding/base64"
)

type RequestBase struct {
	Key string

	Hash string
}

func (r RequestBase) ValidateHash() bool {
	hash := r.GetHash()
	return r.Hash == hash
}

func (r RequestBase) GetHash() string {
	return CreateHash(r.Key)
}

func CreateHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
