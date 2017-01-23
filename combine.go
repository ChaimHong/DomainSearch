package ds

type Combines struct {
	c []string
}

func combineloop(arr, now []rune, i, n int, output *Combines) {
	for _, v := range arr {
		var newnow []rune
		if i > 0 {
			newnow = make([]rune, len(now))
			copy(newnow, now)
		}

		newnow = append(newnow, v)
		if len(newnow) < n {
			combineloop(arr, newnow, i+1, n, output)
		} else {
			output.c = append(output.c, string(newnow))
		}
	}
}

func GetCombineMatch(arr []rune, n int) []string {
	output := &Combines{}
	combineloop(arr, []rune{}, 0, n, output)
	return output.c
}
