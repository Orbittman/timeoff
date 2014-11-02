package dto

type Requester interface {
	ValidateHash() bool
}
