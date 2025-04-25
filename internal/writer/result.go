package writer

type WriteResult struct {
	path  string
	count int
}

func NewResult(path string, count int) *WriteResult {
	return &WriteResult{path: path, count: count}
}

func (result *WriteResult) Path() string {
	return result.path
}

func (result *WriteResult) Count() int {
	return result.count
}
