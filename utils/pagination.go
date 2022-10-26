package utils

import (
	"encoding/json"
	"go-starter/dto"
	"go-starter/enums"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"
)

type PaginationOptions struct {
	Limit    int
	MaxLimit int
	Sort     dto.Sort
}

func Pagination(r *http.Request, options ...PaginationOptions) dto.Pagination {
	query := r.URL.Query()

	var (
		limit    int
		maxLimit int
		sort     dto.Sort
	)

	if len(options) > 0 {
		option := options[0]

		if option.Limit > 0 {
			limit = option.Limit
		}

		if option.MaxLimit > 0 {
			maxLimit = option.MaxLimit
		}

		if len(option.Sort.By) > 0 {
			sort.By = option.Sort.By
		}
		if option.Sort.Direction != 0 {
			sort.Direction = option.Sort.Direction
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

	if len(sort.By) == 0 {
		sort.By = "createdAt"
	}
	if !slices.Contains(
		[]enums.SortDirection{
			enums.Sort.Direction.Asc,
			enums.Sort.Direction.Desc,
		}, sort.Direction,
	) {
		sort.Direction = enums.Sort.Direction.Desc
	}

	return dto.Pagination{
		Limit:   limit,
		Offset:  limit * (page - 1),
		Keyword: keyword,
		Filter:  filter,
		Sort:    sort,
	}
}
