package persistence

type BasePlan struct {
	ID       int64
	SellMode string
	Title    string
	Plans    []Plan
}
