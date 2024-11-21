package commons

import (
	"maps"
)

const (
	TraceId    = "trace-id"
	Profile    = "profile"
	Platform   = "platform"
	Token      = "token"
	PrivateUid = "private-uid"
)

const (
	ShareToken = "stoken"
	ShareScene = "sscene"
)

type QuickInfo struct {
	NotLogSqlConf bool

	*UserInfo
}

type UserInfo struct {
	Uid       int64
	BUserRole string
	OpId      int64
}

func (info *UserInfo) AccountType() int {
	v := info.Uid & int64(MaskAccount)
	return int(v)
}

func NewBaseContext() *BaseContext {
	return &BaseContext{
		container: map[string]string{},
		qinfo: &QuickInfo{
			NotLogSqlConf: false,
			UserInfo:      &UserInfo{},
		},
	}
}

type BaseContext struct {
	container map[string]string
	qinfo     *QuickInfo
}

func (bc *BaseContext) Put(key string, value string) {
	bc.container[key] = value
}

func (bc *BaseContext) Get(key string) string {
	return bc.container[key]
}

func (bc *BaseContext) Clone() *BaseContext {
	n := &BaseContext{
		container: map[string]string{},
	}
	maps.Copy(n.container, bc.container)
	return n
}

func (bc *BaseContext) QuickInfo() *QuickInfo {
	return bc.qinfo
}

func GetToken(bc *BaseContext) string {
	return bc.Get(Token)
}

func GetPlatform(bc *BaseContext) string {
	return bc.Get(Platform)
}

func GetProfile(bc *BaseContext) string {
	return bc.Get(Profile)
}
