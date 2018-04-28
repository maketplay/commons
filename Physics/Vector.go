package Physics

import "math"

type Vector struct {
	x      float64
	y      float64
	length float64
}

func NewVector(from Point, to Point) *Vector {
	v := new(Vector)
	v.x = float64(to.PosX) - float64(from.PosX)
	v.y = float64(to.PosY) - float64(from.PosY)
	return v
}

func (v *Vector) Normalize() *Vector {
	length := v.Length()
	if length > 0 {
		v.Scale(1 / length)
	}
	return v
}

func (v *Vector) SetLength(length float64) *Vector {
	v.Scale(length / v.Length())
	return v
}

func (v *Vector) Invert() *Vector {
	v.x = -v.x
	v.y = -v.y
	return v
}

func (v *Vector) Scale(t float64) *Vector {
	v.x *= t
	v.y *= t
	return v
}

func (v *Vector) Sin() float64 {
	return v.y / v.Length()
}

func (v *Vector) Cos() float64 {
	return v.x / v.Length()
}

func (v *Vector) Angle() float64 {
	return math.Acos(v.Cos())
}

func (v *Vector) OppositeAngle() float64 {
	return math.Acos(v.Cos())
}

func (v *Vector) Length() float64 {
	return math.Hypot(float64(v.x), float64(v.y))
}

func (v *Vector) Add(vector *Vector) *Vector {
	v.x += vector.x
	v.y += vector.y
	return v
}

func (v *Vector) TargetFrom(point Point) Point {
	return Point{
		point.PosX + int(math.Round(v.x)),
		point.PosY + int(math.Round(v.y)),
	}
}