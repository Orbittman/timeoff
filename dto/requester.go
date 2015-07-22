package dto

type Requester interface {
	ValidateChecksum() bool
	GetHash() string
	Key() string
}
