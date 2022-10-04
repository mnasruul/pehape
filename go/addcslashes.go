package pehape

func Addcslashes(s string, c string) string {
	var tmpRune []rune
	strRune := []rune(s)
	list := []rune(c)
	for _, ch := range strRune {
		for _, v := range list {
			if ch == v {
				tmpRune = append(tmpRune, '\\')
			}
		}
		tmpRune = append(tmpRune, ch)
	}
	return string(tmpRune)
}
