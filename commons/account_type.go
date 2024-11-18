package commons

const (
	AAccount = 0
	CAccount = 1
	BAccount = 2
)

func UniqAccountId(accountId int64, accountType int) int64 {
	return accountId<<6 | int64(accountType)
}
