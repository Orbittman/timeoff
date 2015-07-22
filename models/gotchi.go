package models

import(
	"time"
)

type Gotchi struct{
	Name string
	Created time.Time
	Xp int
	Weight int
	Happiness int
	LastFed time.Time
	Health int
	MaxHealth int
}

func (g Gotchi) Init(){
	// Each time it is inaitialised we need to set it's current state
		
}

func (g *Gotchi) Feed(f Food) {
	
}