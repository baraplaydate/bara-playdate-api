package paginate

import (
	"bara-playdate-api/constant"
	"fmt"
	"strconv"

	"github.com/spf13/cast"
)

func PreparePagination(params map[string]string, allowedSortColumns []string) (paging Datapaging) {

	search := params["search"]
	sortBy := params["sort_by"]
	sortDirection := params["sort_direction"]
	page := cast.ToInt(params["page"])

	limit := PaginationMinLimit

	if limitStr, ok := params["limit"]; ok && limitStr != "" {
		l, err := strconv.Atoi(limitStr)

		if err == nil && l > 0 {
			limit = l
		} else {
			limit = PaginationMinLimit
		}
	}

	if page < FirstPage {
		page = FirstPage
	}

	paging = Datapaging{
		Limit:       limit,
		Page:        page,
		OrderBy:     []string{prepareSortBy(sortBy, allowedSortColumns), prepareSortDirection(sortDirection)},
		FilterValue: search,
	}

	return paging
}

func prepareSortBy(param string, allowedSortColumns []string) string {

	sortByFound := false
	for _, allowedSortColumn := range allowedSortColumns {
		if allowedSortColumn == param {
			sortByFound = true
			break
		}
	}

	if !sortByFound {
		return DefaultSortColumn
	}

	return param
}

func prepareSortDirection(param string) string {

	allowedSortDirections := []string{SortAscending, SortDescending}

	sortDirectionFound := false
	for _, allowedSortDirection := range allowedSortDirections {
		if allowedSortDirection == param {
			sortDirectionFound = true
			break
		}
	}

	if !sortDirectionFound {
		return SortAscending
	}

	return param
}

func PrepareStatusValues(statusList []string) ([]string, error) {
	statusMap := map[string]string{
		"ACTIVE":    constant.ACTIVE,
		"NONACTIVE": constant.NONACTIVE,
	}

	var statusValues []string
	for _, status := range statusList {
		if val, ok := statusMap[status]; ok {
			statusValues = append(statusValues, val)
		} else {
			return nil, fmt.Errorf("status %s is not valid", status)
		}
	}
	return statusValues, nil
}
