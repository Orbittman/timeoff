package models

import (
	"time"
)

type Food struct{
	ModelBase
	Name string
	Energy int
	Fat int
	Protein int
	Sweetness int
	Bitterness int
	Sourness int
	MinimumLevel int
	SeasonStart time.Time
	SeasonEnd time.Time
}

func (f *Food) DataKey() string {
	return f.Name
}