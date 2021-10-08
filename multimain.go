package multimain

import "github.com/robdimsdale/uuid"

func MustUUID() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uuid
}
