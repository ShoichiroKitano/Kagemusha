package kagemusha

type Kagemusha struct {
	mocked Function
}

func Mock(mocked interface{}) Kagemusha {
	return Kagemusha{}
}

func (self Kagemusha) Return(returnValue interface{}) {
}

