package dto

type Requester interface {
	Key() string
	GetChecksum() string
}
