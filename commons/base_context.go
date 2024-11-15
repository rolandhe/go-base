package commons

import (
	"maps"
)

const (
	TraceId    = "trace-id"
	Profile    = "profile"
	UID        = "Uid"
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
	QuickInfo() *QuickInfo
}

type QuickInfo struct {
	NotLogSqlConf bool
	Uid           int64
	UserType      string
	BUserRole     string
}

func NewBaseContext() BaseContext {
	return &baseContextImpl{
		container: map[string]string{},
		qinfo: &QuickInfo{
			NotLogSqlConf: false,
			Uid:           -1,
		},
	}
}

type baseContextImpl struct {
	container map[string]string
	qinfo     *QuickInfo
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

func (bc *baseContextImpl) QuickInfo() *QuickInfo {
	return bc.qinfo
}

func GetToken(bc BaseContext) string {
	return bc.Get(Token)
}

func GetShareToken(bc BaseContext) string {
	return bc.Get(ShareToken)
}
