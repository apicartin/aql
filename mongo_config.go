package aql

var mongoOpMap = make(map[string]string)

func initMap() {
	mongoOpMap["="] = "$eq"
	mongoOpMap["!="] = "$ne"
	mongoOpMap["<>"] = "$ne"
	mongoOpMap[">"] = "$gt"
	mongoOpMap[">="] = "$gte"
	mongoOpMap["<"] = "$lt"
	mongoOpMap["<="] = "$lte"
	mongoOpMap["in"] = "$in"
	mongoOpMap["nin"] = "$nin"
}
