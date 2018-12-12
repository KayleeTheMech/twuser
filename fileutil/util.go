package fileutil

import "strconv"

func StringToIntArray(stringArray []string) []int64 {
	var intArray []int64
	for _, str := range stringArray {
		intValue, err := strconv.ParseInt(str, 10, 64)
		checkError(err)
		intArray = append(intArray, intValue)
	}
	return intArray
}

func ReadTwitterIds(fileName string) []int64 {
	lines, err := ReadLinesFromFile(fileName)
	checkError(err)
	return StringToIntArray(lines)
}

func ReadTwitterIdsFromFilesInPath(path string) (ids []int64) {
	files, err := ListFilesInDir(path)
	checkError(err)

	for _, fileName := range files {
		newIds := ReadTwitterIds(path + "/" + fileName)
		ids = append(ids, newIds...)
	}
	return
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
