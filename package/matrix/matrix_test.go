package matrix

import (
	"errors"
	"testing"
)

func TestNewEmptyMatrix(t *testing.T) {
	t.Run("creation of a Empty Matrix with Zero Values", func(t *testing.T) {
		m := NewEmptyMatrix(3, 3)
		if m.Rows != 3 {
			t.Fatal(errors.New("invalid row count expect 3"))
		}
		if m.Cols != 3 {
			t.Fatal(errors.New("invalid col count expect 3"))
		}
		if m.v[0][0] != 0 {
			t.Fatal(errors.New("invalid value at 0 0"))
		}
		if m.v[0][1] != 0 {
			t.Fatal(errors.New("invalid value at 0 1"))
		}
		if m.v[1][0] != 0 {
			t.Fatal(errors.New("invalid value at 1 0"))
		}
		if m.v[1][1] != 0 {
			t.Fatal(errors.New("invalid value at 1 1"))
		}
	})
}

func TestNewMatrix(t *testing.T) {
	m := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	t.Run("creation of a Matrix from 2D Array as Initial Value", func(t *testing.T) {
		if m.Cols != 2 {
			t.Fatal(errors.New("invalid col count expect 2"))
		}
		if m.Rows != 2 {
			t.Fatal(errors.New("invalid row count expect 2"))
		}
		if m.v[0][0] != 1 {
			t.Fatal(errors.New("invalid Value at 0 0 expect 1"))
		}
		if m.v[0][1] != 2 {
			t.Fatal(errors.New("invalid Value at 0 1 expect 2"))
		}
		if m.v[1][0] != 3 {
			t.Fatal(errors.New("invalid Value at 1 0 expect 3"))
		}
		if m.v[1][1] != 4 {
			t.Fatal(errors.New("invalid Value at 1 1 expect 4"))
		}
	})

	t.Run("return as 2D Array", func(t *testing.T) {
		tmp := m.As2DArray()
		if tmp[0][0] != 1 {
			t.Fatal(errors.New("invalid Value at 0 0 expect 1"))
		}
		if tmp[0][1] != 2 {
			t.Fatal(errors.New("invalid Value at 0 1 expect 2"))
		}
		if tmp[1][0] != 3 {
			t.Fatal(errors.New("invalid Value at 1 0 expect 3"))
		}
		if tmp[1][1] != 4 {
			t.Fatal(errors.New("invalid Value at 1 1 expect 4"))
		}
	})

	t.Run("return as 1D Array", func(t *testing.T) {
		tmp := m.As1DArray()
		if tmp[0] != 1 {
			t.Fatal(errors.New("invalid Value at 0 expect 1"))
		}
		if tmp[1] != 3 {
			t.Fatal(errors.New("invalid Value at 1 expect 3"))
		}
		if tmp[2] != 2 {
			t.Fatal(errors.New("invalid Value at 2 expect 2"))
		}
		if tmp[3] != 4 {
			t.Fatal(errors.New("invalid Value at 3 expect 4"))
		}
	})
}

func TestNewMatrix1D(t *testing.T) {
	t.Run("creation of a Matrix with 1D Array as Initial Value", func(t *testing.T) {
		m := NewMatrix1D([]float64{1, 2, 2, 3}, 2)
		if m.Cols != 2 {
			t.Fatal(errors.New("invalid col count expect 2"))
		}
		if m.Rows != 2 {
			t.Fatal(errors.New("invalid row count expect 2"))
		}
		if m.v[0][0] != 1 {
			t.Fatal(errors.New("invalid Value at 0 0 expect 1"))
		}
		if m.v[0][1] != 2 {
			t.Fatal(errors.New("invalid Value at 0 1 expect 2"))
		}
		if m.v[1][0] != 2 {
			t.Fatal(errors.New("invalid Value at 1 0 expect 2"))
		}
		if m.v[1][1] != 3 {
			t.Fatal(errors.New("invalid Value at 1 1 expect 3"))
		}
	})
}

func TestMatrix_IsSquare(t *testing.T) {
	msq := NewMatrix([][]float64{
		{1, 1},
		{1, 1},
	})
	mnsq := NewMatrix([][]float64{
		{1, 1, 1},
		{1, 1, 1},
	})
	if !msq.IsSquare() {
		t.Fatal(errors.New("the Matrix must be Square"))
	}
	if mnsq.IsSquare() {
		t.Fatal(errors.New("the Matrix can't be Square"))
	}
}

func TestMatrix_Fill1D(t *testing.T) {
	m := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	err := m.Fill1D([]float64{1, 1, 1, 1}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if m.v[0][0] != 1 {
		t.Fatal(errors.New("value at 0 0 expect 1"))
	}
	if m.v[0][1] != 1 {
		t.Fatal(errors.New("value at 0 1 expect 1"))
	}
	if m.v[1][0] != 1 {
		t.Fatal(errors.New("value at 1 0 expect 1"))
	}
	if m.v[1][1] != 1 {
		t.Fatal(errors.New("value at 1 1 expect 1"))
	}
}

func TestMatrix_Fill2D(t *testing.T) {
	m := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	err := m.Fill2D([][]float64{
		{1, 1},
		{1, 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	if m.v[0][0] != 1 {
		t.Fatal(errors.New("value at 0 0 expect 1"))
	}
	if m.v[0][1] != 1 {
		t.Fatal(errors.New("value at 0 1 expect 1"))
	}
	if m.v[1][0] != 1 {
		t.Fatal(errors.New("value at 1 0 expect 1"))
	}
	if m.v[1][1] != 1 {
		t.Fatal(errors.New("value at 1 1 expect 1"))
	}
}

func TestMatrix_Add(t *testing.T) {
	m := NewMatrix([][]float64{
		{3, 2},
		{0, 1},
	})
	ma := NewMatrix([][]float64{
		{1, 3},
		{2, 0},
	})
	err := m.Add(ma)
	if err != nil {
		t.Fatal(err)
	}
	if m.v[0][0] != 4 {
		t.Fatal(errors.New("invalid value in 0 0 expect 4"))
	}
	if m.v[0][1] != 5 {
		t.Fatal(errors.New("invalid value in 0 1 expect 5"))
	}
	if m.v[1][0] != 2 {
		t.Fatal(errors.New("invalid value in 1 0 expect 2"))
	}
	if m.v[1][1] != 1 {
		t.Fatal(errors.New("invalid value in 1 1 expect 1"))
	}
}

func TestMatrix_Subtract(t *testing.T) {
	m := NewMatrix([][]float64{
		{8, 7},
		{6, 5},
	})
	ma := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	err := m.Subtract(ma)
	if err != nil {
		t.Fatal(err)
	}
	if m.v[0][0] != 7 {
		t.Fatal(errors.New("invalid value in 0 0 expect 7"))
	}
	if m.v[0][1] != 5 {
		t.Fatal(errors.New("invalid value in 0 1 expect 5"))
	}
	if m.v[1][0] != 3 {
		t.Fatal(errors.New("invalid value in 1 0 expect 3"))
	}
	if m.v[1][1] != 1 {
		t.Fatal(errors.New("invalid value in 1 1 expect 1"))
	}
}

func TestMatrix_Inverse(t *testing.T) {
	base := NewMatrix([][]float64{
		{1, 2},
		{2, 3},
	})
	m1 := NewEmptyMatrix(2, 2)
	m2 := NewEmptyMatrix(2, 2)
	err := base.Inverse(m1, m2)
	if err != nil {
		t.Fatal(errors.New("err is not nil"))
	}
	if m1.v[0][0] != 1 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m1.v[0][1] != 0 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m1.v[1][0] != 0 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m1.v[1][1] != 1 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m2.v[0][0] != -3 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m2.v[0][1] != 2 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m2.v[1][0] != 2 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
	if m2.v[1][1] != -1 {
		t.Fatal(errors.New("invalid result for Matrix"))
	}
}
