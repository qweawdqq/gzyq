package main

import (
	"strings"
	"fmt"
	"gzyq/action"
)



/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-11-29 20:54
 */

func main() {
	b := "  1212  "
	b = strings.TrimSpace(b)
	b = ">" + b + "<"
	fmt.Println(b)
	fmt.Println(action.ComMap["+"].DoCom(2, 5))
	fmt.Println(action.ComMap["-"].DoCom(2, 5))
	fmt.Println(action.ComMap["*"].DoCom(2, 5))
	fmt.Println(action.ComMap["/"].DoCom(2, 5))
}