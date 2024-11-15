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
	QuickInfo() *QuickInfo
}

type QuickInfo struct {
	NotLogSqlConf bool
	uid           int64
}

func NewBaseContext() BaseContext {
	return &baseContextImpl{
		container: map[string]string{},
		qinfo: &QuickInfo{
			NotLogSqlConf: false,
			uid:           -1,
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

func GetUid(bc BaseContext) int64 {
	if bc.QuickInfo().uid != -1 {
		return bc.QuickInfo().uid
	}
	sUid := bc.Get(UID)
	if sUid == "" {
		bc.QuickInfo().uid = 0
		return 0
	}
	bc.QuickInfo().uid, _ = strconv.ParseInt(sUid, 10, 64)
	return bc.QuickInfo().uid
}

func GetToken(bc BaseContext) string {
	return bc.Get(Token)
}

func GetShareToken(bc BaseContext) string {
	return bc.Get(ShareToken)
}
