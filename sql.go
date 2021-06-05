package aql

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/iancoleman/strcase"
)

//SqlParser -
type SQLParser struct {
}

//Parse  to sql based
func (mp SQLParser) Parse(f string, snakeCase bool) interface{} {
	initSqlMap()
	sql := ""
	conditions := []string{}
	r := d.Decode(f)

	for k, v := range r {
		if snakeCase {
			k = strcase.ToSnake(k)
		}

		cri1 := v[0]
		cond1 := fmt.Sprintf(sqlOpMap[cri1.Operator], k, handleInt64ForSql(cri1.Value))
		if len(v) == 2 {
			// and operator
			cri2 := v[1]

			cond2 := fmt.Sprintf(sqlOpMap[cri2.Operator], k, handleInt64ForSql(cri2.Value))
			conditions = append(conditions, cond1+" and "+cond2)
		} else if len(v) == 1 {

			conditions = append(conditions, cond1)
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

func (mp SQLParser) Sort(f string) interface{} {
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

//handleInt64ForSql -
func handleInt64ForSql(v interface{}) interface{} {

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
			str1 := fmt.Sprintf("%v", handleInt64ForSql(vv))
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
