package device42

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func stringChecksum(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func idsToString(ids []int) string {
	s := make([]string, len(ids))
	for _, i := range ids {
		s = append(s, strconv.FormatInt(int64(i), 10))
	}

	return strings.Join(s, "")
}

func idsChecksum(ids []int) string {
	return stringChecksum(idsToString(ids))
}

func interfaceSliceToStringSlice(i []interface{}) []string {
	var s []string = make([]string, len(i))
	for n, d := range i {
		s[n] = fmt.Sprintf("%v", d)
	}
	return s
}
