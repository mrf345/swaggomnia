package utils

import "regexp"

func GetPathParams(path string) [][]string {
	return regexp.
		MustCompile("{% (.*?), '(.*?)', (.*?) %}").
		FindAllStringSubmatch(path, -1)
}
