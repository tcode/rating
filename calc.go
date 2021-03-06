package rating

/*
 Copyright Tobias Lindgaard
 LICENSE GPLv3 or later
 for further details read the LICENSE file.
*/

import (
	"math"
)

type Player struct {
	P float64 // short for points
	R float64 // short for rounds
	Ra float64 // Short for rating
	E float64 // Short for expected
	K float64 // Slot for the K factor.
}
// I could make a struct where I store the k factor, so that you just set it
// then you only need to change it if you want to calculate for another 
// k-factor. 


/*
 I would like to have a variable containing the Player.
 But then I might want more than one, which seems to be a bit hacky.
 But I would perhaps also like to have a commmand center to issue the
 commands to the other functions.
*/

func InitPlayer(Ra, K float64 ) Player {
	N := Player{0, 0, Ra, 0, K}
	return N
} // Returns a Player "object" with 4 fields for float64, everything in this package is float64


func Expected(rating, opp float64) float64 {
	temp := math.Pow(10, (opp-rating)/400)
	expected := 1/(1+temp)
	return expected
} // This is perhaps the single most important function in this file.
// Since all other functions are depending on this, or at least they should be.


func win(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp)*k
	return loss
}

func draw(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp-0.5)*k
	return loss
}

func loss(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp-1.0)*k
	return loss
}

func (a *Player) PWin(opp float64) *Player {
	res := win(a.Ra, opp)
	a.Ra += res
	a.R += 1
	a.P +=1
	a.E += Expected(a.Ra, opp)
	return a
}

func (a *Player) PDraw(opp float64) *Player {
	res := draw(a.Ra, opp)
	a.Ra += res
	a.R += 1
	a.P += 0.5
	a.E += Expected(a.Ra, opp)
	return a
}

func(a *Player) PLoss(opp float64) *Player {
	res := loss(a.Ra, opp)
	a.Ra += res
	a.R += 1
	a.P += 0
	a.E += Expected(a.Ra, opp)
	return a
}


func Check_Bonus(a Player) float64 {
	bon := float64(0)
	if a.R >= 5 {
		bon = 1.5
	} else if a.R > 7 {
		bon = 2
	} else if a.R > 9 {
		bon = 2.5
	} else if a.R > 11 {
		bon = 3
	} else if a.R > 13 {
		bon = 3.5
	}
	extra := a.P - bon
	val := float64(0)
	if extra > 0 {
		val = extra
	}
	return val
}
