package usecases

import (
	"fmt"

	db "github.com/jpmoraess/appointment-api/db/sqlc"
)

type Register struct {
	store db.Store
}

func NewRegister(store db.Store) *Register {
	return &Register{store: store}
}

func (r *Register) Execute() error {
	fmt.Println("registrando...")
	return nil
}
