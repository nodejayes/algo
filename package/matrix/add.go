package matrix

import "errors"

func (m *Matrix) Add(R *Matrix) error {
	if m.Cols != R.Cols || m.Rows != R.Rows {
		return errors.New("the input Matrix has not the same Rows and Cols as the current Matrix")
	}
	for i := range m.v {
		for j := range m.v[i] {
			m.v[i][j] += R.v[i][j]
		}
	}
	return nil
}
