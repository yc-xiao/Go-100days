package m2

import "fmt"

var tempMap = make(map[string]string)

func init() {
	fmt.Println("mp3")
}

func GetMap(key string)  (string, bool){
	value, ok := tempMap[key]
	return value, ok
}

func AddMap(key, value string) {
	tempMap[key] = value
}
