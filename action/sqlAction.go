package action

import (
	"github.com/astaxie/beego/orm"
	"gzyq/logUtils"
	"fmt"
	"strings"
	"errors"
	"bytes"
)


/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-12-23 23:22
 */
type SqlAction struct {
	Id   string
	Name string //名字
	Text string //待解析字符串
	Sort string
}

func (action *SqlAction)SetId(id string) {
	action.Id = id
}
func (action *SqlAction)GetSort() string {
	return action.Sort
}
func (action *SqlAction)GetText() string {
	return action.Text
}
func (action *SqlAction)GetName() string {
	return action.Name
}
func (action *SqlAction)SetText(text string) {
	action.Text = text
}
func (action *SqlAction)SetName(name string) {
	action.Name = name
}
func (action *SqlAction)SetSort(sort string) {
	action.Sort = sort
}
func (action *SqlAction)DoAction(mMap map[string]string, oneOrm orm.Ormer, onelog logUtils.LogUtils, sfkqOrm bool) error {
	sql, keyList, err := getSubSqlAction(action.GetText(), mMap)
	fmt.Println(keyList)
	if err != nil {
		return err
	}
	if sfkqOrm {
		fmt.Println(sql)
		var mapList []orm.Params
		num, err := oneOrm.Raw(sql, keyList).Values(&mapList)
		if err!= nil && num > 0 {
			return err
		}
		mMap[action.GetName()] = fmt.Sprint(mapList)
		fmt.Println("数据库查询结果",mapList)
		//fmt.Println(maps[0]["user_name"]) // slene

	}

	return nil
}
func getSubSqlAction(str string, mMap map[string]string) (string, []string, error) {
	var buffer bytes.Buffer
	var keyList []string
	if str != "" {
		array := strings.Split(str, SubKey)
		fmt.Println("长度", len(array))
		if len(array) > 1 {
			//如果长度大于1 说明是需要截取的  把截取的KEY值改成“？”
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
						buffer.WriteString(SubSqlKey)
						keyList = append(keyList, value)
						fmt.Println("list里面的key值", keyList)
					} else {
						err := errors.New("不能发现key值为<" + v + ">的值")
						return "", nil, err
					}
					//}
				}
				isOddNumber = !isOddNumber
			}
			return buffer.String(), keyList, nil
		}
		return array[0], nil, nil
	}
	return "", nil, errors.New("待解析字符串不允许为空")
}