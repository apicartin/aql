package aql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	NumberFormatter  = " %s %s %d "
	StringFormatter  = " %s %s '%s' "
	DefaultFormatter = " %s %s %v "
)

//SQLParser -
type SQLParser struct {
}

//Parse  to sql based
func (sp SQLParser) Parse(f string, snakeCase bool) interface{} {
	initSQLMap()
	sql := ""
	conditions := []string{}
	r := d.Decode(f)

	for k, v := range r {
		if snakeCase {
			k = strcase.ToSnake(k)
		}
		if len(v) == 1 {
			cri1 := v[0]
			cond1 := sp.apply(k, cri1)
			conditions = append(conditions, cond1)
		} else if len(v) > 1 {

			// add or conditions
			orConds := []string{}
			for _, vc := range v {
				cond := sp.apply(k, vc)
				orConds = append(orConds, cond)
			}

			conditions = append(conditions, strings.Join(orConds, " or "))
		}

	}

	for i, s := range conditions {
		if i == len(conditions)-1 {
			sql = sql + fmt.Sprintf(" ( %s ) ", s)
		} else {
			sql = sql + fmt.Sprintf(" ( %s ) and ", s)
		}
	}
	return sql
}

//Sort -
func (sp SQLParser) Sort(f string) interface{} {
	sql := "order by "
	r := make(map[string]string)
	err := json.Unmarshal([]byte(f), &r)
	if err != nil {
		return ""
	}

	for k, v := range r {
		sql = sql + " " + k + " " + v
	}
	return sql
}

func (sp SQLParser) apply(fieldName string, m Criteria) string {

	formatSt := ""
	v := handleInt64ForSQL(m.Value)

	switch v.(type) {
	case float32:
		formatSt = NumberFormatter
		break
	case float64:
		formatSt = NumberFormatter
		break
	case string:
		formatSt = StringFormatter
		break
	case int64:
		formatSt = NumberFormatter
		break
	case int:
		formatSt = NumberFormatter
		break
	case int16:
		formatSt = NumberFormatter
		break
	case int32:
		formatSt = NumberFormatter
		break
	case bool:
		formatSt = NumberFormatter
		break
	default:
		formatSt = DefaultFormatter
	}

	res := fmt.Sprintf(formatSt, fieldName, m.Operator, v)

	return res
}

func handleInt64ForSQL(v interface{}) interface{} {

	switch v.(type) {
	case string:

		n, err := strconv.ParseInt(v.(string), 10, 64)
		if err == nil {

			if len(v.(string)) == 19 {
				return n
			}
		}
		break

	case []interface{}:
		val := ""
		for i, vv := range v.([]interface{}) {
			str1 := fmt.Sprintf("%v", handleInt64ForSQL(vv))
			if i == len(v.([]interface{}))-1 {
				val = val + str1
			} else {
				val = val + str1 + " , "
			}

		}
		v = val
	}
	return v
}
