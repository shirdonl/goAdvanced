package pkg

import "math"

type Vector struct {
	X float32
	Y float32
}

// 坐标点相加
func (vector1 Vector) add(vector2 Vector) Vector {
	return Vector{vector1.X + vector2.X, vector1.Y + vector2.Y}
}

// 坐标点相减
func (vector1 Vector) sub(vector2 Vector) Vector {
	return Vector{vector1.X - vector2.X, vector1.Y - vector2.Y}
}

// 坐标点相乘
func (vector1 Vector) multi(Speed float32) Vector {
	return Vector{vector1.X * Speed, vector1.Y * Speed}
}

// 计算距离
func (vector1 Vector) distanceTo(vector2 Vector) float32 {
	dX := vector1.X - vector2.X
	dY := vector1.Y - vector2.Y
	distance := math.Sqrt(float64(dX*dX + dY*dY))
	return float32(distance)
}

// 矢量单位化
func (vector1 Vector) normalize() Vector {
	mag := vector1.X*vector1.X + vector1.Y*vector1.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vector{vector1.X * oneOverMag, vector1.Y * oneOverMag}
	} else {
		return Vector{0, 0}
	}
}
