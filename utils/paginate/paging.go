package paginate

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	FirstPage          = 1
	PaginationMinLimit = 10
	SortAscending      = "asc"
	SortDescending     = "desc"
	DefaultSortColumn  = "id"
)

type Datapaging struct {
	DateInTimestamp bool

	Limit int
	Page  int

	//OrderBy define the order property of the row, use it with format [field] [asc/desc].
	//Example []string{"distance desc"}
	OrderBy []string

	OrderByMulti []string

	//FilterColumn specify column as filter parameter
	FilterColumn string
	//FilterValue specify value of the column to be filtered
	FilterValue string

	DateLatest   *time.Time
	DateEarliest *time.Time

	DateBetweenPrefix string
}

func (pagination *Datapaging) BuildQueryGORM(db *gorm.DB) *gorm.DB {

	if pagination.WithLimit() {
		db = db.Limit(pagination.Limit)
	}

	if pagination.WithPageOffset() {
		db = db.Offset(pagination.GetOffset())
	}

	if pagination.WithOrderBy() {
		db = db.Order(fmt.Sprintf("%s %s", pagination.OrderBy[0], pagination.OrderBy[1]))
	}

	if pagination.WithOrderByMulti() {
		for _, orderData := range pagination.OrderByMulti {
			db = db.Order(orderData)
		}
	}

	if pagination.WithDateBetween() {
		createdAtCol := "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest.Unix(),
			pagination.DateLatest.AddDate(0, 0, 1).Unix())
	}

	if pagination.WithDateTimeBetween() {
		createdAtCol := "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest,
			pagination.DateLatest.AddDate(0, 0, 1))
	}

	return db
}

func (pagination *Datapaging) BuildQueryGORMWithParam(db *gorm.DB, namingTable string) *gorm.DB {

	if pagination.WithLimit() {
		db = db.Limit(pagination.Limit)
	}

	if pagination.WithPageOffset() {
		db = db.Offset(pagination.GetOffset())
	}

	if pagination.WithDateBetween() {
		createdAtCol := namingTable + "." + "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest.Unix(),
			pagination.DateLatest.AddDate(0, 0, 1).Unix())
	}

	if pagination.WithDateTimeBetween() {
		createdAtCol := namingTable + "." + "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest,
			pagination.DateLatest.AddDate(0, 0, 1))
	}

	return db
}

func (pagination *Datapaging) WithLimit() bool {
	if pagination.Limit != 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithPageOffset() bool {
	if pagination.Page != 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithOrderBy() bool {
	if len(pagination.OrderBy) > 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithOrderByMulti() bool {
	if len(pagination.OrderByMulti) > 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) GetOffset() int {
	return (pagination.Page - 1) * pagination.Limit
}

func (pagination *Datapaging) WithDateBetween() bool {
	if pagination.DateEarliest != nil && pagination.DateLatest != nil && !pagination.DateInTimestamp {
		return true
	}

	return false
}

func (pagination *Datapaging) WithDateTimeBetween() bool {
	if pagination.DateEarliest != nil && pagination.DateLatest != nil && pagination.DateInTimestamp {
		return true
	}

	return false
}
