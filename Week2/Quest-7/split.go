package piscine

func Split(s, sep string) []string {
	var mstrArr []string
	whereiam := 0
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:len(sep)+i] == sep {
			mstrArr = append(mstrArr, s[whereiam:i])
			whereiam = i + len(sep)
		}
	}
	if whereiam < len(s)-1 {
		if s[len(s)-len(sep):] == sep {
			mstrArr = append(mstrArr, s[whereiam:len(s)-len(sep)])
		} else {
			mstrArr = append(mstrArr, s[whereiam:])
		}
	}
	return mstrArr
}
