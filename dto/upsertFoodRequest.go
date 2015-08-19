package dto

type UpsertFoodRequest struct{
	RequestBase
	Name string
}

func (r UpsertFoodRequest) Key() string {
	return r.Name
}