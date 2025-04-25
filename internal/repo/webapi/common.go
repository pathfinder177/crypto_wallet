package webapi

func parseAnswer(joined string) []string {
	strSliceOutput := []string{}
	lastPos := 0
	for i := range joined {
		if joined[i] == '\n' {
			s := string(joined[lastPos : i+1])
			strSliceOutput = append(strSliceOutput, s)
			lastPos = i + 1
		}
	}

	return strSliceOutput
}
