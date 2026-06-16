func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)

	for _ , i := range nums{
		freq[i]++
	} 

	type Pair struct{
		Num int
		Freq int
	}

	pair := []Pair{}

	for num , count := range freq{
		pair = append(pair,Pair{
			Num : num,
			Freq : count,
		})
	}

	sort.Slice(pair,func(i,j int)bool{
		return pair[i].Freq > pair[j].Freq
	})

	

	result := []int{}

	for i:=0 ; i<k;i++{
		result = append(result,pair[i].Num)
	}

return result
}
