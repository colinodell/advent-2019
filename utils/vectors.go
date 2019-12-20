package utils

type Vector2 struct {
	X, Y int
}

type Vector3 struct {
	X, Y, Z int
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

func (v *Vector2) Multiply(i int) Vector2 {
	return Vector2{
		X: v.X * i,
		Y: v.Y * i,
	}
}

func (v *Vector2) Copy() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}

func (v *Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v *Vector3) Compare(v2 Vector3) Vector3 {
	return Vector3{
		X: compare(v.X, v2.X),
		Y: compare(v.Y, v2.Y),
		Z: compare(v.Z, v2.Z),
	}
}

func compare(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}
