package aql

import (
	"encoding/json"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//JSONDecoder -
type JSONDecoder struct {
	internalMap map[string]interface{}
}

//Decode -
func (jd JSONDecoder) Decode(jsonString string) map[string][]Criteria {

	resultMap := make(map[string][]Criteria)

	err := json.Unmarshal([]byte(jsonString), &resultMap)
	if err != nil {
		log.Errorln(err)
	}
	log.Infoln(resultMap)
	travereCriteria(&resultMap)
	return resultMap
}
func travereCriteria(m *map[string][]Criteria) {

	for _, v := range *m {
		for _, c := range v {
			traverseMap(&c)
		}
	}
}
func traverseMap(m *Criteria) {
	v := m.Value
	switch v.(type) {
	case float32:
		s := fmt.Sprintf("%.0f", v)
		a, _ := strconv.ParseInt(s, 10, 64)
		m.Value = a
		break
	case float64:
		s := fmt.Sprintf("%.0f", v)
		m.Value, _ = strconv.ParseInt(s, 10, 64)
		break
	case []interface{}:
		newAr := traverseArray(v.([]interface{}))
		m.Value = newAr
		break
	default:
		m.Value = v
	}

}

func traverseArray(a []interface{}) []interface{} {
	var result = new([]interface{})

	for _, v := range a {
		switch v.(type) {
		case float32:
			s := fmt.Sprintf("%.0f", v)
			r, _ := strconv.ParseInt(s, 10, 64)
			*result = append(*result, r)
			break
		case float64:
			s := fmt.Sprintf("%.0f", v)
			r, _ := strconv.ParseInt(s, 10, 64)
			*result = append(*result, r)
			break
		case []interface{}:
			traverseArray(a)
			break
		default:
			*result = append(*result, v)
		}
	}
	return *result
}
