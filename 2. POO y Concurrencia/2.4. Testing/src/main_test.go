package main

func TestSum(t *testing.T) {
	tableCases := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tableCases {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Failed: Incorrect sum, got %d expected %d", total, item.n)
		}
	}
}
