package selectsort
//选择排序：每次选出一个最小的放到最前面，下一次从剩下的选最小的放第二个

func Sort(array []int) {
	len := len(array)
	for i := 0; i < len; i++ {
		min := i
		for j := i+1; j < len; j++ {
			if array[j] < array[min]{
				min = j
			}
		}
		if i != min {
			array[i], array[min] = array[min], array[i]
		}
	}
}
