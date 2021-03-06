package aql

var sqlOpMap = make(map[string]string)
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
	mongoOpMap["like"] = "$regex"
}
func initSQLMap() {
	sqlOpMap["="] = " %s = '%v' "
	sqlOpMap["!="] = " %s != '%v' "
	sqlOpMap[">"] = " %s > '%v' "
	sqlOpMap["<"] = " %s < '%v' "
	sqlOpMap["<>"] = " %s <> '%v' "
	sqlOpMap["<="] = " %s <= '%v' "
	sqlOpMap[">="] = " %s >= '%v' "
}
