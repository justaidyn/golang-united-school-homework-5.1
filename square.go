package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	a := Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
	return a
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return (s.a + s.a) * 2
}
