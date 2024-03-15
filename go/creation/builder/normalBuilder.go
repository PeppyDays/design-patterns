package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType(windowType string) {
	b.windowType = windowType
}

func (b *NormalBuilder) setDoorType(doorType string) {
	b.doorType = doorType
}

func (b *NormalBuilder) setNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
