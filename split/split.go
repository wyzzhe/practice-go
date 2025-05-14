package split

import "strings"

// split.go

// Split 将s按照sep进行分割，返回一个字符串的切片
// Split("我爱你", "爱") => ["我"， "你"]
func Split(s, sep string) (rec []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		rec = append(rec, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	rec = append(rec, s)
	return rec
}
