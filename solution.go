package square

import (
	"math"
)

type IntCustomType int

const (
	SidesCircle IntCustomType = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum IntCustomType) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return (math.Pow(sideLen, 2) * math.Sqrt(sideLen)) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
