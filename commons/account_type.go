package commons

const (
	AAccount    = 62
	BAccount    = 31
	CAccount    = 15
	MaskAccount = 0x3F
)

const (
	CommonCompany   = 61
	PracticeCompany = 17
	PlatformCompany = 3
)

func GetPlatformCompanyId() int64 {
	return uniqAccountId(1, PlatformCompany)
}

func GetPracticeCompanyId() int64 {
	return uniqAccountId(2, PracticeCompany)
}

func uniqAccountId(accountId int64, accountType int) int64 {
	return accountId<<6 | int64(accountType)
}
