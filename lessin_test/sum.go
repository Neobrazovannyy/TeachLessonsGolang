package sum

func sum(arr_numbers ...int) int{
	var return_sum int = 0;

	for _, num := range(arr_numbers){
		return_sum+=num
	}

	return return_sum
}