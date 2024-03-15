package builder

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType("a")
	d.builder.setWindowType("b")
	d.builder.setNumFloor(10)
	return d.builder.getHouse()
}
