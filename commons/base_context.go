package commons

import "maps"

const (
	TraceId   = "trace-id"
	Profile   = "profile"
	UID       = "uid"
	Role      = "role"
	Possessor = "possessor"
)

type BaseContext interface {
	Put(key string, value string)
	Get(key string) string
	Clone() BaseContext
}

func NewBaseContext() BaseContext {
	return &baseContextImpl{
		container: map[string]string{},
	}
}

type baseContextImpl struct {
	container map[string]string
}

func (bc *baseContextImpl) Put(key string, value string) {
	bc.container[key] = value
}

func (bc *baseContextImpl) Get(key string) string {
	return bc.container[key]
}

func (bc *baseContextImpl) Clone() BaseContext {
	n := &baseContextImpl{
		container: map[string]string{},
	}
	maps.Copy(n.container, bc.container)
	return n
}
