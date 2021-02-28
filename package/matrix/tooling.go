package matrix

// IsSquare is the current Matrix a Square Matrix
func (m *Matrix) IsSquare() bool {
	return m.Cols == m.Rows
}
