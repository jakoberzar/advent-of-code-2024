package main

type Direction interface {
	Eligible(x, y int) bool
	Next(x, y int) (int, int)
}

type Up struct {
	MaxY int
	MaxX int
	Len  int
}

func (o Up) Eligible(x, y int) bool {
	return y >= o.Len-1 && y < o.MaxY
}

func (o Up) Next(x, y int) (int, int) {
	return x, y - 1
}

type Left struct {
	MaxY int
	MaxX int
	Len  int
}

func (o Left) Eligible(x, y int) bool {
	return x >= o.Len-1 && x < o.MaxX
}

func (o Left) Next(x, y int) (int, int) {
	return x - 1, y
}

type Down struct {
	MaxY int
	MaxX int
	Len  int
}

func (o Down) Eligible(x, y int) bool {
	return y <= o.MaxY-o.Len && y >= 0
}

func (o Down) Next(x, y int) (int, int) {
	return x, y + 1
}

type Right struct {
	MaxY int
	MaxX int
	Len  int
}

func (o Right) Eligible(x, y int) bool {
	return x <= o.MaxX-o.Len && x >= 0
}

func (o Right) Next(x, y int) (int, int) {
	return x + 1, y
}

type LeftUp struct {
	MaxY int
	MaxX int
	Len  int
}

func (o LeftUp) Eligible(x, y int) bool {
	return y >= o.Len-1 && y < o.MaxY && x >= o.Len-1 && x < o.MaxX
}

func (o LeftUp) Next(x, y int) (int, int) {
	return x - 1, y - 1
}

type RightUp struct {
	MaxY int
	MaxX int
	Len  int
}

func (o RightUp) Eligible(x, y int) bool {
	return x <= o.MaxX-o.Len && x >= 0 && y >= o.Len-1 && y < o.MaxY
}

func (o RightUp) Next(x, y int) (int, int) {
	return x + 1, y - 1
}

type LeftDown struct {
	MaxY int
	MaxX int
	Len  int
}

func (o LeftDown) Eligible(x, y int) bool {
	return y <= o.MaxY-o.Len && y >= 0 && x >= o.Len-1 && x < o.MaxX
}

func (o LeftDown) Next(x, y int) (int, int) {
	return x - 1, y + 1
}

type RightDown struct {
	MaxY int
	MaxX int
	Len  int
}

func (o RightDown) Eligible(x, y int) bool {
	return x <= o.MaxX-o.Len && x >= 0 && y <= o.MaxY-o.Len && y >= 0
}

func (o RightDown) Next(x, y int) (int, int) {
	return x + 1, y + 1
}
