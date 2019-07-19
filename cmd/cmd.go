package cmd

import (
	"errors"
	"log"

	"function/internal"
)

//-------------------------------------------
// Exported function
//-------------------------------------------

// MainLogic implements the main logic
func MainLogic(payload internal.Event) error {
	print(payload)

	if payload.Name == "" {
		return errors.New("empty name")
	}

	return nil
}

//-------------------------------------------
// Internal functions
//-------------------------------------------
func print(payload internal.Event) {
	log.Printf("%+v", payload)
}
