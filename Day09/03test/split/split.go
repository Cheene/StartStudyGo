// split/split.go

package split

import (
	"strings"
)

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) (ret []string) {
	// sep 在那个位置
	// 按 b 进行切割
	ret = make([]string, 0, strings.Count(s, sep)+1)
	index := strings.Index(s, sep)
	for index > -1 {
		ret = append(ret, s[:index])
		s = s[index+len(sep):] // sep 的长度不一定为1
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
