package cache

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/callbacks"

	"gorm.io/gorm"
)

const IdentifierPrefix = "gorm-caches::"

func (c *Caches) buildIdentifier(db *gorm.DB) string {
	return buildIdentifier(db, c.Conf.Pfx)
}
func buildIdentifier(db *gorm.DB, prefix ...string) string {
	// Build query identifier,
	//	for that reason we need to compile all arguments into a string
	//	and concat them with the SQL query itself
	callbacks.BuildQuerySQL(db)
	query := db.Statement.SQL.String()
	queryArgs := valueToString(db.Statement.Vars)
	pfx := IdentifierPrefix
	if len(prefix) > 0 && prefix[0] != "" {
		pfx = prefix[0]
	}
	identifier := fmt.Sprintf("%s%s-%s", pfx, query, queryArgs)
	return identifier
}

func valueToString(value any) string {
	valueOf := reflect.ValueOf(value)
	switch valueOf.Kind() {
	case reflect.Ptr:
		if valueOf.IsNil() {
			return "<nil>"
		}
		return valueToString(valueOf.Elem().Interface())
	case reflect.Map:
		var sb strings.Builder
		sb.WriteString("{")
		for i, key := range valueOf.MapKeys() {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s: %s", valueToString(key.Interface()), valueToString(valueOf.MapIndex(key).Interface())))
		}
		sb.WriteString("}")
		return sb.String()
	case reflect.Slice:
		valueSlice := make([]any, valueOf.Len())
		for i := range valueSlice {
			valueSlice[i] = valueToString(valueOf.Index(i).Interface())
		}
		return fmt.Sprintf("%v", valueSlice)
	default:
		return fmt.Sprintf("%v", value)
	}
}
