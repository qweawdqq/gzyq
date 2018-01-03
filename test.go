package main

import (
	"strings"
	"fmt"

	"errors"
	"bytes"
	"gzyq/action"
)



/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-11-29 20:54
 */

func main() {
	//b := "  1212  "
	//b = strings.TrimSpace(b)
	//b = ">" + b + "<"
	//fmt.Println(b)
	//fmt.Println(action.ComMap["+"].DoCom(2, 5))
	//fmt.Println(action.ComMap["-"].DoCom(2, 5))
	//fmt.Println(action.ComMap["*"].DoCom(2, 5))
	//fmt.Println(action.ComMap["/"].DoCom(2, 5))
	sql := "select * from User t where t.id = $hehe$ and name = $333$"
	mMap := make(map[string]string)
mMap["hehe"] = "124534634634636436"
	mMap["333"] = "name"
	str,list,err :=getSubAction(sql,mMap)
	fmt.Println(err)
	fmt.Println(str)
	fmt.Println(list)
}


func getSubAction(str string, mMap map[string]string) (string,[]string, error) {
	var buffer bytes.Buffer
	var keyList []string
	if str != "" {
		array := strings.Split(str, action.SubKey)
			fmt.Println("长度",len(array))
			if len(array)>1 {
				fmt.Println(array)
			isOddNumber := true
			for _, v := range array {
				fmt.Println(v)
				if isOddNumber {
					//if v != "" {
						buffer.WriteString(v)
					//}
				} else {
					//if v != "" {
						if value, ok := mMap[v]; ok {
							buffer.WriteString(action.SubSqlKey)
							keyList=  append(keyList,value)
							fmt.Println("list里面的key值",keyList)
						} else {
							err := errors.New("不能发现key值为<" + v + ">的值")
							return "",nil, err
						}
					//}
				}
				isOddNumber = !isOddNumber
			}
			return buffer.String(),keyList, nil
		}
		return array[0],nil,nil

	}
	return "", nil,errors.New("待解析字符串不允许为空")
}