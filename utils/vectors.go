package utils

type Vector2 struct {
	X, Y int
}

func (v *Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector2) Min(v2 Vector2) Vector2 {
	return Vector2{
		X: Min(v.X, v2.X),
		Y: Min(v.Y, v2.Y),
	}
}

func (v *Vector2) Max(v2 Vector2) Vector2 {
	return Vector2{
		X: Max(v.X, v2.X),
		Y: Max(v.Y, v2.Y),
	}
}
