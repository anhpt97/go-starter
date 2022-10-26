package enums

type SortDirection int

type direction struct {
	Asc  SortDirection
	Desc SortDirection
}

var Sort = struct {
	Direction direction
}{
	Direction: direction{
		Asc:  1,
		Desc: -1,
	},
}
