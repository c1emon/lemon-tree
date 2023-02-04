package httpx

import (
	"net/http"
	"regexp"
	"strconv"
)

type PaginationQuery struct {
	PaginationParams
}

func PaginationFromQuery(req *http.Request) *PaginationQuery {
	paginationQuery := &PaginationQuery{
		PaginationParams{Size: 10, Total: 100, Page: 0, Sorts: make([]SortItem, 0)},
	}

	queries := req.URL.Query()
	if s := queries.Get("size"); s != "" {
		if size, err := strconv.Atoi(s); err == nil {
			paginationQuery.Size = size
		}
	}
	if p := queries.Get("page"); p != "" {
		if page, err := strconv.Atoi(p); err == nil {
			paginationQuery.Page = page
		}
	}
	if t := queries.Get("total"); t != "" {
		if total, err := strconv.Atoi(t); err == nil {
			paginationQuery.Total = total
		}
	}
	if s := queries.Get("sort"); s != "" {
		reg := regexp.MustCompile("(?i)(desc|asc)\\(([a-zA-Z]+[0-9a-zA-Z_]*)\\)")
		for _, k := range reg.FindAllStringSubmatch(s, -1) {
			paginationQuery.Sorts = append(paginationQuery.Sorts, SortItem{Field: k[2], Order: ParserOrder(k[1])})
		}
	}

	return paginationQuery
}
