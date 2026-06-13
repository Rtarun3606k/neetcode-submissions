func dailyTemperatures(temperatures []int) []int {

	result := make([]int , len(temperatures))

	for i:=0 ; i<len(temperatures);i++{

		for j:=i+1 ; j<len(temperatures);j++{
		if temperatures[i]<temperatures[j]{
			result[i]=j-i
			break
		}
		}
	}
	return result
}
