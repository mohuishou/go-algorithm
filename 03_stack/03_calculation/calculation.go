package calculation

import (
	"fmt"
	"strconv"
)

// 操作符的优先级
var operatorPriority = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"(": 2,
	")": 2,
}

// Calculator 计算器
type Calculator struct {
	nums      *StackInt
	operators *Stack
	exp       string
}

// NewCalculator NewCalculator
func NewCalculator(exp string) *Calculator {
	return &Calculator{
		nums:      NewStackInt(),
		operators: NewStack(),
		exp:       exp,
	}
}

// Calculate 获取计算结果
func (c *Calculator) Calculate() int {
	l := len(c.exp)
	for i := 0; i < l; i++ {
		switch e := (c.exp[i]); e {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// 一直往后获取数字，如果下一个还是数字说明这一个数还不完整
			j := i
			for j < l && c.exp[j] <= '9' && c.exp[j] >= '0' {
				j++
			}
			n, _ := strconv.Atoi(c.exp[i:j])
			i = j - 1
			c.nums.Push(n)
		case '+', '-', '*', '/':
			// 从计算符栈中获取栈顶元素，如果当前操作符的优先级低于栈顶元素的优先级
			// 并且栈顶元素不为空，和括号
			// 那么从数据栈中取两个数据和栈顶操作符进行计算
			pre := c.operators.Pop()
			for pre != "" && pre != "(" && operatorPriority[string(e)] <= operatorPriority[pre] {
				c.nums.Push(c.calc(pre))
				pre = c.operators.Pop()
			}
			if pre != "" {
				c.operators.Push(pre)
			}
			c.operators.Push(string(e))
		case '(':
			c.operators.Push(string(e))
		case ')':
			// 碰到右括号之后就一直不断操作符栈中弹出元素，并且取两个数据进行计算
			// 直到碰到左括号为止
			for o := c.operators.Pop(); o != "(" && o != ""; o = c.operators.Pop() {
				c.nums.Push(c.calc(o))
			}
		default:
			panic("invalid exp")
		}
	}
	// 最后如果不存在操作符，说明数据栈中的栈顶元素就是最后结果
	o := c.operators.Pop()
	if o == "" {
		return c.nums.Pop()
	}
	// 如果存在，就把最后的数据进行计算后返回
	return c.calc(o)
}

// calc 单次计算操作，o: 计算符
func (c *Calculator) calc(o string) int {
	b := c.nums.Pop()
	a := c.nums.Pop()

	fmt.Printf("%d %s %d\n", a, o, b)

	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return 0
}

// calculate 计算器，支持加减乘除
func calculate(s string) int {
	return NewCalculator(s).Calculate()
}
