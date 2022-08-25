package enums

type SortDirection string

type direction struct {
	Asc  SortDirection
	Desc SortDirection
}

var Sort = struct {
	Direction direction
}{
	Direction: direction{
		Asc:  "ASC",
		Desc: "DESC",
	},
}
