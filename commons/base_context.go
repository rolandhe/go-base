package commons

import (
	"maps"
	"time"
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
	Uid            int64
	UserName       string
	Mobile         string
	Email          string
	Roles          int64
	OpId           int64
	CompanyId      int64
	CompanyName    string
	TokenId        int64
	MemberExpireAt int64
}

func (info *UserInfo) AccountType() int {
	v := info.Uid & int64(MaskAccount)
	return int(v)
}

func (info *UserInfo) IsAttach() bool {
	if info.AccountType() == BAccount && info.OpId != info.Uid {
		return true
	}
	return false
}

func (info *UserInfo) IsRole(role int64) bool {
	return info.Roles&role == role
}

func (info *UserInfo) IsAdmin() bool {
	return info.IsRole(AdminRole)
}

func NewBaseContext() *BaseContext {
	return &BaseContext{
		container:  map[string]string{},
		createTime: time.Now().UnixMilli(),
		qinfo: &QuickInfo{
			NotLogSqlConf: false,
			UserInfo:      &UserInfo{},
		},
	}
}

type BaseContext struct {
	container  map[string]string
	createTime int64
	qinfo      *QuickInfo
}

func (bc *BaseContext) Put(key string, value string) {
	bc.container[key] = value
}

func (bc *BaseContext) Get(key string) string {
	return bc.container[key]
}

func (bc *BaseContext) GetCreateTime() int64 {
	return bc.createTime
}

func (bc *BaseContext) Clone() *BaseContext {
	n := &BaseContext{
		container:  map[string]string{},
		createTime: bc.createTime,
		qinfo: &QuickInfo{
			NotLogSqlConf: bc.qinfo.NotLogSqlConf,
			UserInfo:      &UserInfo{},
		},
	}
	*n.qinfo.UserInfo = *bc.qinfo.UserInfo
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
