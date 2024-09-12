package main

import (
	"math"
)

type Vector [3]float64

func (v *Vector) GetX() float64 {
	return v[0]
}

func (v *Vector) GetY() float64 {
	return v[1]
}

func (v *Vector) GetZ() float64 {
	return v[2]
}

func (v *Vector) Negation() Vector {
	return [3]float64{-v[0], -v[1], -v[2]}
}

func (v *Vector) Add(u *Vector) Vector {
	return Vector{
		v[0] + u[0],
		v[1] + u[1],
		v[2] + u[2],
	}
}

func (v *Vector) Sub(u *Vector) Vector {
	uNeg := u.Negation()
	return v.Add(&uNeg)
}

func (v *Vector) MultiplyConst(m float64) Vector {
	return Vector{
		v[0] * m,
		v[1] * m,
		v[2] * m,
	}
}

func (v *Vector) Dot(u *Vector) Vector {
	return Vector{
		v[0] * u[0],
		v[1] * u[1],
		v[2] * u[2],
	}
}

func (v *Vector) Cross(u *Vector) Vector {
	return Vector{u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0]}
}

func (v *Vector) Divide(d float64) Vector {
	return v.MultiplyConst(1 / d)
}

func (v *Vector) LengthSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vector) UnitVector() Vector {
	return v.Divide(v.Length())
}
