package rating

import (
	"math"
)

type Player struct {
	P float64 // short for points
	R float64 // short for rounds
	Ra float64 // Short for rating
	E float64 // Short for expected
}

/*
 The above code is part of a implementation of some bonus score checking thing.
 Not implemented yet.
*/

// I could make a struct where I store the k factor, so that you just set it
// then you only need to change it if you want to calculate for another 
// k-factor. 


/*
 I would like to have a variable containing the Player.
 But then I might want more than one, which seems to be a bit hacky.
 But I would perhaps also like to have a commmand center to issue the commaands to the other functions.
*/


func InitPlayer(P, R, Ra float64 ) Player {
	N := Player{P, R, Ra, 0}
	return N
} // Returns a Player "object" with 4 fields for float64, everything in this package is float64


func Expected(rating, opp float64) float64 {
	temp := math.Pow(10, (opp-rating)/400)
	expected := 1/(1+temp)
	return expected
} // This is perhaps the single most important function in this file.
// Since all other functions are depending on this, or at least they should be.


func Loss(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp-1.0)*k
	return loss
}

func Win(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp)*k
	return loss

}

func Draw(rating, opp float64) float64 {
	k := float64(45)
	exp := Expected(rating, opp)
	loss := (1.0-exp-0.5)*k
	return loss
}

func PWin(a Player, opp float64) Player {
	rating := a.Ra
	res := Win(rating, opp)
	a.Ra = res
	a.R += 1
	a.P +=1
	a.E += Expected(rating, opp)
	return a
}

func PDraw(a Player, opp float64) Player {
	rating := a.Ra
	res := Draw(rating, opp)
	a.Ra = res
	a.R += 1
	a.P += 0.5
	a.E += Expected(rating, opp)
	return a
}

func PLoss(a Player, opp float64) Player {
	rating := a.Ra
	res := Loss(rating, opp)
	a.Ra = res
	a.R += 1
	a.P += 0
	a.E += Expected(rating, opp)
	return a
}


// New functions for use with a player, and not as a single calculation, the above
// Are not designed to be used for a things other than quick single calculations on one game rating.
