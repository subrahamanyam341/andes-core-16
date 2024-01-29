package block

import "github.com/subrahamanyam341/andes-core-16/data"

// EmptyBlockCreator is able to create empty block instances
type EmptyBlockCreator interface {
	CreateNewHeader() data.HeaderHandler
	IsInterfaceNil() bool
}
