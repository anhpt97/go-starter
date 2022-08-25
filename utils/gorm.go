package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Model interface {
	TableName() string
}

// q := DB.Model()...Session(&gorm.Session{})
// q1 := q.Where("id = ?", id1)...
// q2 := q.Where("id = ?", id2)...

// q := DB.Model()...Where("id = ?", id)
// q1 := q.Session(&gorm.Session{})...
// q2 := q.Session(&gorm.Session{})...

// DB.Table("book"). // for tables without model structure
// 	Select("1 + 2 AS sum, \"abc\" AS title").
// 	Scan(&books)

// .Omit(clause.Associations) // skip all associations

func GetAllColumnNamesOfTableQuery(model Model) string {
	s := []string{}

	regexp := regexp.MustCompile(`column:\w+`)
	tableName := model.TableName()
	aliasTableName := cases.Title(language.Und).String(tableName)

	t := reflect.TypeOf(model)
	for i := 0; i < t.NumField(); i++ {
		columnName := strings.ReplaceAll(
			regexp.FindString(t.Field(i).Tag.Get("gorm")),
			"column:", "",
		)
		if len(columnName) == 0 {
			continue
		}
		s = append(s, fmt.Sprintf("%v AS %v",
			tableName+"."+columnName,
			aliasTableName+"__"+columnName,
		))
	}

	return strings.Join(s[:], ", ")
}
