package dto

type NewGotchiRequest struct{
	RequestBase
	Name string
}

func (r NewGotchiRequest) Key() string{
	return r.Name
}