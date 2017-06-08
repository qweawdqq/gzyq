package action

import (
	"bytes"
	"strings"
	"fmt"
	"errors"
	"strconv"
	"log"
	"gzyq/oneConfig"
)

type Action interface {
	GetName() string
	SetName(name string)
	SetText(text string) //设置待解析字符串
	GetText() string
	DoAction(mMap map[string]string) error
}

type SubAction struct {
	Name string //名字
	Text string //待解析字符串
}
type ComAction struct {
	Name string //名字
	Text string //待解析字符串
}

func (action *SubAction)DoAction(mMap map[string]string) error {
	var str = ""
	if strings.Contains(action.GetText(), "$") {
		str1, err := getSubAction(action.GetText(), mMap)
		if err != nil {
			mMap[oneConfig.ONE_ERROR_MSG] = err.Error()
			mMap[oneConfig.ONE_ERROR_NAME] = oneConfig.ONE_ERROR_ACTION+action.Name
			return err
		}
		str = str1
	} else {
		str = action.GetText()
	}
	mMap[action.GetName()] = str
	return nil

}
func (action *SubAction)GetText() string {
	return action.Text
}
func (action *SubAction)GetName() string {
	return action.Name
}
func (action *SubAction)SetText(text string) {
	action.Text = text
}
func (action *SubAction)SetName(name string) {
	action.Name = name
}
/**
 * @Description: 拼接字符串的方法 （根据$）
 * @author : 贾亮
 * @date : 2017/6/6 14:01
 */
func getSubAction(str string, mMap map[string]string) (string, error) {
	var buffer bytes.Buffer
	if str != "" {
		array := strings.Split(str, "$")
		isOddNumber := true
		for _, v := range array {
			if isOddNumber {
				fmt.Println("奇数" + v)
				if v != "" {
					buffer.WriteString(v)
				}
			} else {
				fmt.Println("偶数" + v)
				if v != "" {
					if value, ok := mMap[v]; ok {
						buffer.WriteString(value)
					} else {
						fmt.Println("Key Not Found in ")
						err := errors.New("不能发现key值为<" + v + ">的值")
						return "", err
					}
				}
			}
			isOddNumber = !isOddNumber
		}
		return buffer.String(), nil
	}
	return "", errors.New("待解析字符串不允许为空")
}


//计算部分
func (action *ComAction)DoAction(mMap map[string]string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Runtime error caught: %v", r)
			mMap[oneConfig.ONE_ERROR_NAME] = oneConfig.ONE_ERROR_ACTION+action.Name
			mMap[oneConfig.ONE_ERROR_MSG] = oneConfig.ONE_ERROR_INFO1

		}
	}()
	var str = ""
	if strings.Contains(action.GetText(), "$") {
		str1, err := getSubAction(action.GetText(), mMap)
		if err != nil {
			return err
		}
		str = str1
	} else {
		str = action.GetText()
	}
	str = strconv.FormatFloat(Count(str), 'f', -1, 64)
	mMap[action.GetName()] = str
	return nil
}
func (action *ComAction)GetText() string {
	return action.Text
}
func (action *ComAction)GetName() string {
	return action.Name
}
func (action *ComAction)SetText(text string) {
	action.Text = text
}
func (action *ComAction)SetName(name string) {
	action.Name = name
}

type StackNode struct {
	Data interface{}
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	Count int
}

func (this *LinkStack) Init() {
	this.top = nil
	this.Count = 0
}

func (this *LinkStack) Push(data interface{}) {
	var node *StackNode = new(StackNode)
	node.Data = data
	node.next = this.top
	this.top = node
	this.Count++
}

func (this *LinkStack) Pop() interface{} {
	if this.top == nil {
		return nil
	}
	returnData := this.top.Data
	this.top = this.top.next
	this.Count--
	return returnData
}

//Look up the top element in the stack, but not pop.
func (this *LinkStack) LookTop() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.Data
}

func Count(data string) float64 {
	//TODO 检查字符串输入
	var arr []string = generateRPN(data)
	return calculateRPN(arr)
}

func calculateRPN(datas []string) float64 {
	var stack LinkStack
	stack.Init()
	for i := 0; i < len(datas); i++ {
		//fmt.Println("datas[i]=" + datas[i])
		if isNumberString(datas[i]) {
			if f, err := strconv.ParseFloat(datas[i], 64); err != nil {
				panic("operatin process go wrong.")
			} else {
				stack.Push(f)
			}
		} else {
			p1 := stack.Pop().(float64)
			p2 := stack.Pop().(float64)
			p3 := normalCalculate(p2, p1, datas[i])
			stack.Push(p3)
		}
	}
	res := stack.Pop().(float64)
	return res
}

func normalCalculate(a, b float64, operation string) float64 {
	switch operation{
	case "*":
		return a * b
	case "-":
		return a - b
	case "+":
		return a + b
	case "/":
		return a / b
	default:
		panic("invalid operator")
	}
}

func generateRPN(exp string) []string {

	var stack LinkStack
	stack.Init()

	var spiltedStr []string = convertToStrings(exp)
	var datas []string

	for i := 0; i < len(spiltedStr); i++ {
		// 遍历每一个字符
		tmp := spiltedStr[i] //当前字符
		//fmt.Println("tmp::" + tmp)
		if !isNumberString(tmp) {
			//是否是数字
			// 四种情况入栈
			// 1 左括号直接入栈
			// 2 栈内为空直接入栈
			// 3 栈顶为左括号，直接入栈
			// 4 当前元素不为右括号时，在比较栈顶元素与当前元素，如果当前元素大，直接入栈。
			if tmp == "(" ||
				stack.LookTop() == nil || stack.LookTop().(string) == "(" ||
				( compareOperator(tmp, stack.LookTop().(string)) == 1 && tmp != ")" ) {
				stack.Push(tmp)
			} else {
				// ) priority
				if tmp == ")" {
					//当前元素为右括号时，提取操作符，直到碰见左括号
					for {
						if pop := stack.Pop().(string); pop == "(" {
							break
						} else {
							datas = append(datas, pop)
						}
					}
				} else {
					//当前元素为操作符时，不断地与栈顶元素比较直到遇到比自己小的（或者栈空了），然后入栈。
					for {
						pop := stack.LookTop()
						if pop != nil && compareOperator(tmp, pop.(string)) != 1 {
							datas = append(datas, stack.Pop().(string))
						} else {
							stack.Push(tmp)
							break
						}
					}
				}
			}

		} else {
			datas = append(datas, tmp)
		}
		//fmt.Println(datas)
	}

	//将栈内剩余的操作符全部弹出。
	for {
		if pop := stack.Pop(); pop != nil {
			datas = append(datas, pop.(string))
		} else {
			break
		}
	}
	return datas
}

// if return 1, o1 > o2.
// if return 0, o1 = 02
// if return -1, o1 < o2
func compareOperator(o1, o2 string) int {
	// + - * /
	var o1Priority int
	if o1 == "+" || o1 == "-" {
		o1Priority = 1
	} else {
		o1Priority = 2
	}
	var o2Priority int
	if o2 == "+" || o2 == "-" {
		o2Priority = 1
	} else {
		o2Priority = 2
	}
	if o1Priority > o2Priority {
		return 1
	} else if o1Priority == o2Priority {
		return 0
	} else {
		return -1
	}
}

func isNumberString(o1 string) bool {
	if o1 == "+" || o1 == "-" || o1 == "*" || o1 == "/" || o1 == "(" || o1 == ")" {
		return false
	} else {
		return true
	}
}

func convertToStrings(s string) []string {
	var strs []string
	bys := []byte(s)
	var tmp string
	for i := 0; i < len(bys); i++ {
		if !isNumber(bys[i]) {
			if tmp != "" {
				strs = append(strs, tmp)
				tmp = ""
			}
			strs = append(strs, string(bys[i]))
		} else {
			tmp = tmp + string(bys[i])
		}
		//fmt.Println(strs)
	}
	strs = append(strs, tmp)
	fmt.Println(strs)
	return strs
}

func isNumber(o1 byte) bool {
	//fmt.Println(string(o1))
	if o1 == '+' || o1 == '-' || o1 == '*' || o1 == '/' || o1 == '(' || o1 == ')' {
		return false
	} else {
		return true
	}
}

//--------------------------------resultAction
/**
 * @Description: 主要用来返回结果的
 * @author : 贾亮
 * @date : 2017/6/7 13:01
 */
type ResultAction struct {
	Id         string
	Name       string
	ActionList []Action
	IsReturn   bool //是向下还是返回  返回为true
}

func (action *ResultAction)GetName() string {
	return action.Name
}
func (action *ResultAction)DoAction(mMap map[string]string) {
	for _, v := range action.ActionList {
		v.DoAction(mMap)
	}
}

func (action *ResultAction)SetName(name string) {
	action.Name = name
}

func (action *ResultAction)SetIsReturn(isReturn bool) {
	action.IsReturn = isReturn
}
func (action *ResultAction)SetList(actionList []Action) {
	action.ActionList = actionList
}
func (action *ResultAction)GetIsReturn() bool {
	return action.IsReturn
}

func (action *ResultAction) GetResultStr(mMap map[string]string) string {
	count := len(action.ActionList)
	if count > 0 {
		key := action.ActionList[count - 1].GetName()
		if v, ok := mMap[key]; ok {
			return v
		}
	}
	return ""
}

