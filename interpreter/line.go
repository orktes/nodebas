package interpreter

type (
	LineNumber int64
	Lines      []Line
	Line       struct {
		Number     LineNumber
		Statements Statements
	}
)
