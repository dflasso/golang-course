package main

func main() {
	ununsortList := []int{12, 13, 14}

	for i, v1 := range ununsortList {
		for j, v2 := range ununsortList {

			if v1 < v2 {

				ununsortList[i] = v2
				ununsortList[j] = v1
			}

		}

	}

}
