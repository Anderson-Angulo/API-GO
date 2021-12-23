import "go-api/model"

type Memory struct{
	CurrentID int
	Persons map[int]model.Person
}

func newMemory() Memory{
	persons:= make(map[int]model.Person)
	return Memory{
		CurrentID:0,
		Persons:persons
	}
}