package bubble
//冒泡排序：两两相比每次把最大的气泡放到最后

func Sort(array []int) {
	len := len(array)
	flag := true
	for i := 0; i < len && flag; i++ {
		flag = false
		for j := 0; j < len-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
	}
}