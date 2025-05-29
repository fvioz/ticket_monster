package persistence

type Zone struct {
	ID       int64
	Capacity int64
	Price    float64
	Name     string
	Numbered bool
}
