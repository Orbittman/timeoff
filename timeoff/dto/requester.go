package dto

type Requester interface {
	ValidateHash() bool

	GetHash() string
}
