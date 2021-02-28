package matrix

func (m *Matrix) Multiply(R *Matrix) {
	result := [][]float64{}
	for i := 0; i < len(m.v); i++ {
		result = append(result, make([]float64, len(R.v[0])))
		for j := 0; j < len(R.v[0]); j++ {
			sum := 0.0
			for k := 0; k < len(m.v[0]); k++ {
				sum += m.v[i][k] * R.v[k][j]
			}
			result[i][j] = sum
		}
	}
	m.v = result
}
