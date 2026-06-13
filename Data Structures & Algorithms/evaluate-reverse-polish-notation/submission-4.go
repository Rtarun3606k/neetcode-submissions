
func evalRPN(tokens []string) int {
var stack []int

	for _ , ch :=  range tokens{
		if ch == "*"|| ch == "+" || ch == "/"|| ch == "-" {
			val1 := pop(&stack)
			val2 := pop(&stack)
			eval(ch ,val2,val1,&stack)
		}else{
			num,err :=  strconv.Atoi(ch)
			if err==nil{

			push(num,&stack)
			}
		}

	}
	return stack[len(stack)-1]
}

func push(val int,stack *[]int){
	*stack = append(*stack, val)
}


func pop(stack *[]int) int{
	top1:= (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return top1


}


func eval(token string, a,b int,stack *[]int){

	switch token {
case "+":
    *stack = append(*stack, a+b)

case "-":
    *stack = append(*stack, a-b)

case "*":
    *stack = append(*stack, a*b)

case "/":
    *stack = append(*stack, a/b)
}
}