package models

type Zone struct {
	ID       int64   `xml:"id"`
	Capacity int64   `xml:"capacity"`
	Price    float64 `xml:"price"`
	Name     string  `xml:"name"`
	Numbered bool    `xml:"numbered"`
	Plan     *Plan   `pg:"rel:has-one"`
}
