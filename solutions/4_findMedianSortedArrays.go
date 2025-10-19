package solutions

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//res := FindMedianSortedArraysByChan(nums1, nums2)
	res := FindMedianSortedArraysOptimize(nums1, nums2)
	return res
}

func FindMedianSortedArraysOptimize(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	mid := total/2 + 1

	lenI := len(nums1)
	lenJ := len(nums2)
	currentValue := 0
	lastValue := 0

	j := 0
	i := 0
	for i < lenI || j < lenJ {
		if i+j == mid {
			break
		}

		if i >= lenI {
			//out <- nums2[j]
			lastValue = currentValue
			currentValue = nums2[j]
			j++
			continue
		}

		if j >= lenJ {
			//out <- nums1[i]
			lastValue = currentValue
			currentValue = nums1[i]
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			//out <- nums1[i]
			lastValue = currentValue
			currentValue = nums1[i]
			i++
		} else {
			//out <- nums2[j]
			lastValue = currentValue
			currentValue = nums2[j]
			j++
		}

	}

	if total%2 != 0 {
		return float64(currentValue)
	}
	return float64(lastValue+currentValue) / 2
}

func FindMedianSortedArraysByChan(nums1 []int, nums2 []int) float64 {
	reader := arrToChan(nums1, nums2)
	count := 0
	total := len(nums1) + len(nums2)
	mid := total/2 + 1
	num1 := 0
	num2 := 0
	for item := range reader {
		count++
		if count == mid {
			num1 = item
			break
		}
		num2 = item
	}
	if total%2 != 0 {
		return float64(num1)
	}
	return float64(num2+num1) / 2
}

func arrToChan(nums1 []int, nums2 []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		i := 0
		j := 0
		lenI := len(nums1)
		lenJ := len(nums2)
		for i < lenI || j < lenJ {
			if i >= lenI {
				out <- nums2[j]
				j++
				continue
			}

			if j >= lenJ {
				out <- nums1[i]
				i++
				continue
			}

			if nums1[i] < nums2[j] {
				out <- nums1[i]
				i++
			} else {
				out <- nums2[j]
				j++
			}
		}
	}()
	return out
}
