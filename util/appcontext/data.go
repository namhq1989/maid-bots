package appcontext

import "github.com/google/uuid"

func generateID() string {
	id, _ := uuid.NewV7()
	return id.String()
}
