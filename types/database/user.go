package types

import (
	"fmt"

	"github.com/google/uuid"
	viewmodel "github.com/lunaorg/luna-main-api/types/viewmodel"
)

type registerUserDB struct {
	ID string `dynamodbav:"user_id"`
	PK string `dynamodbav:"pk"`
	SK string `dynamodbav:"sk"`

	Name string `dynamodbav:"name,omitempty"`
}

func ParseRegisterUserViewmodel(item viewmodel.RegisterUser) (*registerUserDB, error) {
	id := uuid.New()

	buildPk := func(login string, id string) string {
		return fmt.Sprintf("%s#%s#%s", "USER", login, id)
	}

	return &registerUserDB{
		ID:   id.String(),
		PK:   buildPk(item.Login, id.String()),
		SK:   "INFO",
		Name: item.Name,
	}, nil
}
