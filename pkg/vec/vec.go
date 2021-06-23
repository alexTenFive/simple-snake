package vec

import (
	"fmt"
	"math"
)

type (
	Vector struct {
		X, Y float64
	}
)

func (x *Vector) String() string {
	if x == nil {
		return ""
	}
	return fmt.Sprintf("(%.3f; %.3f)", x.X, x.Y)
}

func (x Vector) Length() float64 {
	return math.Hypot(x.X, x.Y)
}

func (x *Vector) LengthSqr() float64 {
	return x.X*x.X + x.Y*x.Y
}

func (x *Vector) Scale(f float64) *Vector {
	x.X *= f
	x.Y *= f
	return x
}

func (x *Vector) Clone() Vector {
	return Vector{x.X, x.Y}
}

func (x *Vector) CloneP() *Vector {
	return &Vector{x.X, x.Y}
}

func (x *Vector) Scaled(f float64) Vector {
	return Vector{x.X * f, x.Y * f}
}

func (x *Vector) ScaledP(f float64) *Vector {
	return &Vector{x.X * f, x.Y * f}
}

func (x *Vector) Invert() *Vector {
	x.X = -x.X
	x.Y = -x.Y
	return x
}

func (x *Vector) Inverted() Vector {
	return Vector{-x.X, -x.Y}
}

func (x *Vector) InvertedP() *Vector {
	return &Vector{-x.X, -x.Y}
}

func (x *Vector) Normalize() *Vector {
	sl := x.LengthSqr()
	if sl == 0 || sl == 1 {
		return x
	}
	return x.Scale(1 / math.Sqrt(sl))
}

func (x *Vector) Normalized() Vector {
	xx := x.Clone()
	xx.Normalize()
	return xx
}

func (x *Vector) NormalizedP() *Vector {
	xx := x.CloneP()
	xx.Normalize()
	return xx
}

// Add adds another vector to vec.
func (x *Vector) Add(v *Vector) *Vector {
	x.X += v.X
	x.Y += v.Y
	return x
}

// Sub subtracts another vector from vec.
func (x *Vector) Sub(v *Vector) *Vector {
	x.X -= v.X
	x.Y -= v.Y
	return x
}

// Mul multiplies the components of the vector with the respective components of v.
func (x *Vector) Mul(v *Vector) *Vector {
	x.X *= v.X
	x.Y *= v.Y
	return x
}

// Rotate rotates the vector counter-clockwise by angle.
func (x *Vector) Rotate(angle float64) *Vector {
	*x = x.Rotated(angle)
	return x
}

// Rotated returns a counter-clockwise rotated copy of the vector.
func (x *Vector) Rotated(angle float64) Vector {
	sinus := math.Sin(angle)
	cosinus := math.Cos(angle)
	return Vector{
		x.X*cosinus - x.Y*sinus,
		x.X*sinus + x.Y*cosinus,
	}
}

func Sub(a, b *Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

func SubP(a, b *Vector) *Vector {
	return &Vector{a.X - b.X, a.Y - b.Y}
}

func Dot(a, b *Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}

func Add(v1, v2 *Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func AddP(v1, v2 *Vector) *Vector {
	return &Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

// Angle returns the angle between two vectors.
func Angle(a, b *Vector) float64 {
	v := Dot(a, b) / (a.Length() * b.Length())
	// prevent NaN
	if v > 1. {
		v = v - 2
	} else if v < -1. {
		v = v + 2
	}
	return math.Acos(v)
}

func DegreesBetween(a, b *Vector) float64 {
	var dot = Dot(a, b)
	degree := math.Acos(dot) * (180.0 / math.Pi)
	return degree
}

func nearlyEqual(a, b float64) bool {
	epsilon := 0.000001

	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	if a == 0.0 || b == 0.0 || diff < math.SmallestNonzeroFloat64 {
		return diff < (epsilon * math.SmallestNonzeroFloat64)
	}

	absA := math.Abs(a)
	absB := math.Abs(b)

	return diff/math.Min(absA+absB, math.MaxFloat64) < epsilon
}

func Equal(a, b *Vector) bool {
	return nearlyEqual(a.X, b.X) && nearlyEqual(a.Y, b.Y)
}
