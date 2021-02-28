package matrix

type Matrix struct {
	Cols int
	Rows int
	v    [][]float64
}

// NewEmptyMatrix creates a Matrix with no data based on column and row numbers
func NewEmptyMatrix(cols, rows int) *Matrix {
	tmp := Matrix{
		Cols: cols,
		Rows: rows,
		v:    [][]float64{},
	}
	tmp.v = make([][]float64, cols)
	for idx := range tmp.v {
		tmp.v[idx] = make([]float64, rows)
	}
	return &tmp
}

// NewMatrix creates a Matrix from a 2D Array
func NewMatrix(values [][]float64) *Matrix {
	if len(values) < 1 {
		panic("Matrix must have min. 1 Column")
	}
	return &Matrix{
		Cols: len(values),
		Rows: len(values[0]),
		v:    values,
	}
}

// NewMatrix1D creates a Matrix from a 1D Array
func NewMatrix1D(values []float64, n int) *Matrix {
	tmp := make([][]float64, n)
	col := -1
	for i := range values {
		if i%n == 0 {
			col++
		}
		if col == 0 {
			tmp[0] = append(tmp[0], values[i])
		} else {
			tmp[col] = append(tmp[col], values[i])
		}
	}
	return &Matrix{
		Cols: n,
		Rows: col + 1,
		v:    tmp,
	}
}
