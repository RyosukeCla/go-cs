package countingfilter

import (
	"testing"
)

func Test(t *testing.T) {
	bf := New(1000, 50)
	samples := []string{
		"Hello",
		"World!",
		"Hoge",
		"huga",
		"Super Flute",
		"ahoasjfoaijsdf",
		"ahoasjfoaijsdf12rsfasf",
	}

	for _, sample := range samples {
		bf.Add([]byte(sample))
	}

	for _, sample := range samples {
		if bf.Check([]byte(sample)) != true {
			t.Fatal("error")
		}
	}

	bf.Delete([]byte(samples[0]))
	if bf.Check([]byte(samples[0])) != false {
		t.Fatal("error")
	}

	if bf.Check([]byte("Super Flut3")) != false {
		t.Fatal("error")
	}
}
