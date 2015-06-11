package interpreter

func (lines Lines) Len() int {
	return len(lines)
}
func (lines Lines) Swap(i, j int) {
	lines[i], lines[j] = lines[j], lines[i]
}
func (lines Lines) Less(i, j int) bool {
	return lines[i].Number < lines[j].Number
}
