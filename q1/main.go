package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := ""
	sortedData, lenData := GetDataByPoint(data)
	fmt.Println("the data is: ", sortedData, "the lenght is: ", lenData)
}

func GetDataByPoint(data string) ([]string, int) {
	arrayData := strings.Split(data, ",")
	sortedData := sort.StringSlice(arrayData)
	sortedData.Sort()
	return sortedData, sortedData.Len()
}
