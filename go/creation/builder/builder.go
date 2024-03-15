package builder

type Builder interface {
	setWindowType(windowType string)
	setDoorType(doorType string)
	setNumFloor(numFloor int)
	getHouse() House
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
