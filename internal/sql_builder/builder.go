package sql_builder

import (
	"fmt"

	"gorm.io/gorm"
)

type Operator string

const (
	Eq        Operator = "="
	Neq       Operator = "!="
	Gt        Operator = ">"
	Gte       Operator = ">="
	Lt        Operator = "<"
	Lte       Operator = "<="
	In        Operator = "in"
	Like      Operator = "like"
	NotLike   Operator = "not like"
	IsNull    Operator = "is null"
	IsNotNull Operator = "is not null"
)

type Filter struct {
	Key      string
	Operator Operator
	Value    string
	IsRaw    bool
}

func WithFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, filter := range filters {
		if filter.Operator == IsNull || filter.Operator == IsNotNull {
			db = db.Where(fmt.Sprintf("%s %s", filter.Key, filter.Operator))
		} else if filter.IsRaw {
			db = db.Where(fmt.Sprintf("%s %s %s", filter.Key, filter.Operator, filter.Value))
		} else {
			db = db.Where(fmt.Sprintf("%s %s ?", filter.Key, filter.Operator), filter.Value)
		}
	}

	return db
}
