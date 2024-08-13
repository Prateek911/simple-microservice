package model

type AccountInfo struct {
	Name      string
	CRN       int64
	Balance   int64
	Holdings  int64
	AccType   AccountType
	AccHealth AccountHealth
}

type AccountHealth string
type AccountType string

const (
	OK         AccountHealth = "ok"
	WARN       AccountHealth = "Warning"
	SUB1       AccountHealth = "Substandard 1"
	SUB2       AccountHealth = "Substandard 2"
	SUB3       AccountHealth = "Substandard 3"
	OVERDUE30  AccountHealth = "Overdue by 30 days"
	OVERDUE90  AccountHealth = "Overdue by 90 days"
	DELINQUENT AccountHealth = "Delinquent"
)

const (
	SAVING    AccountType = "Savings Account"
	CURRENT   AccountType = "Current Account"
	GL        AccountType = "GL Account"
	OVERDRAFT AccountType = "Overdraft Account"
)
