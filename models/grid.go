package models

type GridDataRow struct {
	Id       string
	Name     string
	Status   string
	Position string
	Email    string
	Badges   []string
	Img      string
}

type GridData []GridDataRow
