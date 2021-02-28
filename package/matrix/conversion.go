package matrix

import "errors"

func (m *Matrix) As2DArray() [][]float64 {
	return m.v
}

func (m *Matrix) As1DArray() []float64 {
	res := make([]float64, m.Cols*m.Rows)
	for x := range m.v {
		for y := range m.v[x] {
			i := x + m.Rows*y
			res[i] = m.v[x][y]
		}
	}
	return res
}

func (m *Matrix) Fill1D(values []float64, n int) error {
	tmp := NewMatrix1D(values, n)
	return m.Fill2D(tmp.v)
}

func (m *Matrix) Fill2D(values [][]float64) error {
	if len(values) < m.Cols {
		return errors.New("not enough columns")
	}
	if len(values[0]) < m.Rows {
		return errors.New("not enough rows")
	}
	for i := 0; i < m.Cols; i++ {
		for y := 0; y < m.Rows; y++ {
			m.v[i][y] = values[i][y]
		}
	}
	return nil
}
