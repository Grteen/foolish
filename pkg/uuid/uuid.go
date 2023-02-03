package uuid

import (
	uuid "github.com/satori/go.uuid"
)

func GetUUid() string {
	id := uuid.NewV4()
	return id.String()
}
