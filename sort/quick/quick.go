package quick
//快速排序：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另
//外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过
//程可以递归进行，以此达到整个数据变成有序序列
//PS:可以使用三数取中值或九数取中值选择比较合理的mid，之后交换到array[0]位置

func Sort(array []int) {
	if len(array) <= 1 {
		return
	}

	mid := array[0]
	i := 0
	head, tail := 0, len(array)-1
	for head < tail {
		if array[tail] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	Sort(array[:head])
	Sort(array[head+1:])
}