package quick
//快速排序：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另
//外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过
//程可以递归进行，以此达到整个数据变成有序序列
//PS:可以使用三数取中值或九数取中值选择比较合理的mid，之后交换到array[0]位置

func Sort(data []int, done chan struct{}) {
	fmt.Println(data)
	if len(data) <= 1 {
		done <- struct{}{}
		return data
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	childDone := make(chan struct{}, 2)
	// sortArray(data[:head], childDone)
	// sortArray(data[head+1:], childDone)
	go sortArray(data[:head], childDone)
	go sortArray(data[head+1:], childDone)
	<-childDone
	<-childDone

	done <- struct{}{}
	return data
}

func main() {
	a := []int{5, 1, 1, 2, 0, 0}
	done := make(chan struct{})
	go sortArray(a, done)
	<-done
	fmt.Println(a)
