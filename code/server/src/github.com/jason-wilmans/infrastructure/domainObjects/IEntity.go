package domainObjects

import "github.com/satori/go.uuid"

type IEntity interface {
	GetId() uuid.UUID
}
