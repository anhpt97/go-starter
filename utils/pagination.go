package utils

import (
	"encoding/json"
	"go-starter/dto"
	"go-starter/enums"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Sort struct {
	By        string
	Direction enums.SortDirection
}

type DefaultPagination struct {
	Limit    int
	MaxLimit int
	Sort     Sort
}

func Pagination(r *http.Request, paginations ...DefaultPagination) dto.Pagination {
	query := r.URL.Query()

	var (
		limit    int
		maxLimit int
		sort     Sort
	)

	if len(paginations) > 0 {
		defaultPagination := paginations[0]

		if defaultPagination.Limit > 0 {
			limit = defaultPagination.Limit
		}

		if defaultPagination.MaxLimit > 0 {
			maxLimit = defaultPagination.MaxLimit
		}

		if len(defaultPagination.Sort.By) > 0 {
			sort.By = defaultPagination.Sort.By
		}
		if len(defaultPagination.Sort.Direction) > 0 {
			sort.Direction = defaultPagination.Sort.Direction
		}
	} else {
		limit, _ = strconv.Atoi(query.Get(enums.Query.Limit))

		maxLimit = 100

		json.Unmarshal([]byte(query.Get(enums.Query.Sort)), &sort)
	}

	if limit < 1 || limit > maxLimit {
		limit = 10
	}

	page, _ := strconv.Atoi(query.Get(enums.Query.Page))
	if page < 1 {
		page = 1
	}

	keyword := query.Get(enums.Query.Keyword)

	filter := map[string]any{}
	json.Unmarshal([]byte(query.Get(enums.Query.Filter)), &filter)

	if len(strings.TrimSpace(sort.By)) == 0 {
		sort.By = "id"
	}
	if !slices.Contains(
		[]enums.SortDirection{
			enums.Sort.Direction.Asc,
			enums.Sort.Direction.Desc,
		}, enums.SortDirection(strings.ToUpper(strings.TrimSpace(string(sort.Direction)))),
	) {
		sort.Direction = enums.Sort.Direction.Desc
	}

	return dto.Pagination{
		Limit:   limit,
		Offset:  limit * (page - 1),
		Keyword: keyword,
		Filter:  filter,
		Order:   sort.By + " " + string(sort.Direction),
	}
}
