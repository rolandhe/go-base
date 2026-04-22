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

type KvExtendRegisterMode int

const (
	KvExtendRegisterOverride KvExtendRegisterMode = 1
	KvExtendRegisterFirst    KvExtendRegisterMode = 2
	KvExtendRegisterLast     KvExtendRegisterMode = 3
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
	container        map[string]string
	createTime       int64
	qinfo            *QuickInfo
	kvFromHeaderFunc KvExtendFunc
	kvFromQueryFunc  KvExtendFunc
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
	n.kvFromHeaderFunc = bc.kvFromHeaderFunc
	n.kvFromQueryFunc = bc.kvFromQueryFunc
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

func (bc *BaseContext) GetHeaderValue(key string) any {
	if bc.kvFromHeaderFunc == nil {
		return nil
	}
	return bc.kvFromHeaderFunc(key)
}

func (bc *BaseContext) GetHeaderStringValue(key string) string {
	if bc.kvFromHeaderFunc == nil {
		return ""
	}
	v := bc.kvFromHeaderFunc(key)
	if v == nil {
		return ""
	}
	str, ok := v.(string)
	if !ok {
		return ""
	}
	return str
}

func (bc *BaseContext) GetQueryValue(key string) any {
	if bc.kvFromQueryFunc == nil {
		return nil
	}
	return bc.kvFromQueryFunc(key)
}

func (bc *BaseContext) GetQueryStringValue(key string) string {
	if bc.kvFromQueryFunc == nil {
		return ""
	}
	v := bc.kvFromQueryFunc(key)
	if v == nil {
		return ""
	}
	str, ok := v.(string)
	if !ok {
		return ""
	}
	return str
}

func (bc *BaseContext) RegisterKvFromHeaderFunc(kvHeaderFunc KvExtendFunc, mode KvExtendRegisterMode) {
	if bc.kvFromHeaderFunc == nil || mode == KvExtendRegisterOverride {
		bc.kvFromHeaderFunc = kvHeaderFunc
		return
	}
	if kvHeaderFunc == nil {
		return
	}
	old := bc.kvFromHeaderFunc
	if mode == KvExtendRegisterFirst {
		bc.kvFromHeaderFunc = func(key string) any {
			v := kvHeaderFunc(key)
			if v != nil {
				return v
			}
			return old(key)
		}
		return
	}
	bc.kvFromHeaderFunc = func(key string) any {
		v := old(key)
		if v != nil {
			return v
		}
		return kvHeaderFunc(key)
	}
}

func (bc *BaseContext) RegisterKvFromQueryFunc(kvQueryFunc KvExtendFunc, mode KvExtendRegisterMode) {
	if bc.kvFromQueryFunc == nil || mode == KvExtendRegisterOverride {
		bc.kvFromQueryFunc = kvQueryFunc
		return
	}
	if kvQueryFunc == nil {
		return
	}
	old := bc.kvFromQueryFunc
	if mode == KvExtendRegisterFirst {
		bc.kvFromQueryFunc = func(key string) any {
			v := kvQueryFunc(key)
			if v != nil {
				return v
			}
			return old(key)
		}
		return
	}
	bc.kvFromQueryFunc = func(key string) any {
		v := old(key)
		if v != nil {
			return v
		}
		return kvQueryFunc(key)
	}
}
