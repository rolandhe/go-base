package commons

import (
	"maps"
	"strconv"
)

const (
	TraceId    = "trace-id"
	Profile    = "profile"
	UID        = "uid"
	Role       = "role"
	Possessor  = "possessor"
	Platform   = "platform"
	Lang       = "lang"
	Token      = "token"
	ShareToken = "share-token"
	RemoteIp   = "remote-ip"
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

func GetUid(bc BaseContext) int64 {
	sUid := bc.Get(UID)
	if sUid == "" {
		return 0
	}
	uid, _ := strconv.ParseInt(sUid, 10, 64)
	return uid
}

func GetToken(bc BaseContext) string {
	return bc.Get(Token)
}

func GetShareToken(bc BaseContext) string {
	return bc.Get(ShareToken)
}
