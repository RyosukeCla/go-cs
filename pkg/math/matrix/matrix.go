package matrix

type Matrix struct {
	elements []float64
	cols     int
	rows     int
	size     int
}

func Zeros(rows, cols int) *Matrix {
	size := cols * rows
	return &Matrix{
		elements: make([]float64, size),
		cols:     cols,
		rows:     rows,
		size:     size,
	}
}

func Ones(rows, cols int) *Matrix {
	size := cols * rows
	elements := make([]float64, size)
	for i := 0; i < size; i++ {
		elements[i] = 1.0
	}
	return &Matrix{
		elements: elements,
		cols:     cols,
		rows:     rows,
		size:     size,
	}
}

func New(rows, cols int, elements []float64) *Matrix {
	size := cols * rows
	return &Matrix{
		elements: elements,
		cols:     cols,
		rows:     rows,
		size:     size,
	}
}

func (m *Matrix) Element(row, col int) float64 {
	return m.elements[row+m.rows*col]
}

func (m *Matrix) Elements() []float64 {
	return m.elements
}

func (m *Matrix) Add(mat *Matrix) *Matrix {
	res := Zeros(m.rows, m.cols)
	for i := 0; i < m.size; i++ {
		res.elements[i] = m.elements[i] + mat.elements[i]
	}
	return res
}

func (m *Matrix) Sub(mat *Matrix) *Matrix {
	res := Zeros(m.rows, m.cols)
	for i := 0; i < m.size; i++ {
		res.elements[i] = m.elements[i] - mat.elements[i]
	}
	return res
}

func (m *Matrix) Mul(mat *Matrix) *Matrix {
	res := Zeros(m.rows, m.cols)
	for i := 0; i < m.size; i++ {
		res.elements[i] = m.elements[i] * mat.elements[i]
	}
	return res
}

func (m *Matrix) Div(mat *Matrix) *Matrix {
	res := Zeros(m.rows, m.cols)
	for i := 0; i < m.size; i++ {
		res.elements[i] = m.elements[i] / mat.elements[i]
	}
	return res
}

func (m *Matrix) MatMul(mat *Matrix) *Matrix {
	res := Zeros(m.rows, mat.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			for k := 0; k < m.cols; k++ {
				res.elements[i+j*mat.rows] += m.elements[k+i*m.rows] + mat.elements[j+k*mat.rows]
			}
		}
	}

	return res
}
