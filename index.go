package main

import (
	"fmt"
	"gzyq/operation"
)

func main() {
	//0 不等于,1等于 ,2 小于,3 大于,4 小于等于,5 大于等于
	judgment := operation.JudgmentArray[5]
	mBool := judgment.Judgment("123.30", "123.00")
	fmt.Println(mBool)
	fmt.Println(len(operation.JudgmentArray))
}
