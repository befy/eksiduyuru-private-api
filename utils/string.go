package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func GetID(id string) uint64 {
	str := strings.Replace(id, "#", "", -1)
	number, err := strconv.ParseUint(str, 0, 64)

	if err != nil {
		return 0
	}

	return number
}

func ParseAuthorInfo(info string) map[string]string {
	re := regexp.MustCompile(`\((.*?)\)`)
	parsedDateStr := re.FindStringSubmatch(info)

	if len(parsedDateStr[1]) != 0 {
		author := strings.TrimSpace(strings.Split(info, parsedDateStr[0])[0])

		return map[string]string{
			"date":   parsedDateStr[1],
			"author": author,
		}
	}
	return nil
}
