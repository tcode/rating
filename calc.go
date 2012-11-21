package chessboard

import (
  "mystring"
	"strings"
	"fmt"
)


type Board string

type Piece struct {
	Color string
	Type  string
	Row   int
	Col   int
}



func Parse(str string) string { // but what should this return ?
	ree := ""
	for x :=0; x < len(str); x++ {
		value := mystring.GetElem(str, x)
		// Here I would just want to write some if statements
		// for where my code should pass the value off to get parsed.
		if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" ||
			value == "5" || value == "6" || value == "7" || value == "8" {
			p := ParseNumber(value)
			np := mystring.Join(p)
			ree += np
		} else {
			ree += value
			ree += " "
		}
	}
	return ree
}

func Display(str string) {
	lines = string.Split(str, "/")
	for x:=0;x>8;x++ {
		
		for y:=0;y>len(lines[x]);y++ {
			pstr := ParseNumber(mystring.GetElem(lines[x], 0)) 
			fmt.Printf("%v", pstr)
		}
	}
}

func ParseFEN(s string) {
	
}

// this function only takes a single element list
func ParseNumber(s string) []string {
	a:=[]string{}
	if s == "1" {
		a = append(a, ".")
	} else if s == "2" {
		for x:=0;x<2;x++ {
			a = append(a, ".")
		}
	} else if s == "3" {
		for x:=0;x<3;x++ {
			a = append(a, ".")
		}
	} else if s == "4" {
		for x:=0;x<4;x++ {
			a = append(a, ".")
		}
	} else if s == "5" {
		for x:=0;x<5;x++ {
			a = append(a, ".")
		}
	} else if s == "6" {
		for x:=0;x<6;x++ {
			a = append(a, ".")
		}
	} else if s == "7" {
		for x:=0;x<7;x++ {
			a = append(a, ".")
		}
	} else if s == "8" {
		for x:=0;x<8;x++ {
			a = append(a, ".")
		}
	}
	return a
}

func ParseUnit(s string) Piece {
	p:=Piece{}
	lines := strings.Split(s, "/")
	for y:=0;y<len(lines[0]);y++ {
		a := mystring.GetElem(lines[0], y)
		if a == "p" {
			p.Type = "Pawn"
			p.Color = "Black"
			p.Col = y
		} else if a == "P" {
			p.Type = "Pawn"
			p.Color = "White"
			p.Col = y
		} else if a == "n" {
			p.Type = "Knight"
			p.Color = "Black"
			p.Col = y
		} else if a == "N" {
			p.Type = "Knight"
			p.Color = "White"
			p.Col = y
		} else if a == "b" {
			p.Type = "Bishop"
			p.Color = "Black"
			p.Col = y
		} else if a == "B" {
			p.Type = "Bishop"
			p.Color = "White"
			p.Col = y
 		} else if a == "r" {
			p.Type = "Rook"
			p.Color = "Black"
			p.Col = y
		} else if a == "R" {
			p.Type = "Rook"
			p.Color = "White" 
			p.Col = y
		} else if a == "q" {
			p.Type = "Queen"
			p.Color = "Black" 
			p.Col = y
		} else if a == "Q" {
			p.Type = "Queen"
			p.Color = "White" 
			p.Col = y
		} else if a == "k" {
			p.Type = "King"
			p.Color = "Black" 
			p.Col = y
		} else if a == "K" {
			p.Type = "King"
			p.Color = "White" 
			p.Col = y
		} else {
			
		}
	}
	return p
}




// How do I make something which keeps a decent representation of the board.?







