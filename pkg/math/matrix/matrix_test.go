package matrix

import "testing"

func Test(t *testing.T) {
	a := Ones(3, 4)
	b := New(3, 4, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		10, 11, 12,
	})
	actual := a.Add(b)
	
	for i, element := range actual.Elements() {
		if element != float64(i + 2) {
			t.Fatal("error") 
		}
	}
}
