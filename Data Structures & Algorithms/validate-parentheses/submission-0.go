func isValid(s string) bool {
    stack := []rune{}

	pairs := map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}

	for _,ch := range s{
		if ch == '('|| ch == '{' || ch == '['{
			stack = append(stack,ch)
		}else{
			if len(stack)==0{
				return false
			}

			if pairs[ch] != stack[len(stack)-1]{
				return false
			} 
			stack = stack[:len(stack)-1]
		}



	}

	return len(stack)==0
}
