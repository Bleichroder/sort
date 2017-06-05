package insert
//插入排序：每次得到一个新数值，依次与之前序列中的值比大小，完后插入

func Sort(array []int) {
	len := len(array)
	for i, j := 1, 0; i < len; i++ {
		tmp := array[i]
		for j = i; j > 0 && array[j-1] > tmp; j-- {
			array[j] = array[j-1]
		}
		array[j] = tmp
	}
}