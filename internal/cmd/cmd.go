package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/nszilard/log"

	"github.com/nszilard/aws-lambda-go/internal/models"
	"github.com/nszilard/aws-lambda-go/internal/pkg"
)

//-------------------------------------------
// Exported function
//-------------------------------------------

// Entrypoint implements the main logic
func Entrypoint(payload models.Event) (int, error) {
	pl, _ := json.MarshalIndent(payload, "", "  ")
	log.Debugf("Received event payload: %s", pl)

	if payload.Name == "" {
		return 0, fmt.Errorf("payload contains no name")
	}

	return pkg.Fibonacci(payload.Count), nil
}
