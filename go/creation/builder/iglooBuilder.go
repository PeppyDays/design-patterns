package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType(windowType string) {
	b.windowType = windowType
}

func (b *IglooBuilder) setDoorType(doorType string) {
	b.doorType = doorType
}

func (b *IglooBuilder) setNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
