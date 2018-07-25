package domainObjects

import "github.com/satori/go.uuid"

const EntityIdRegex = "[0-9a-fA-F-_]{1,128}"

type IEntity interface {
	GetId() uuid.UUID
}
