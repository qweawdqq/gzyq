package action

import (
	"strconv"
	"fmt"
	"gzyq/oneConfig"
	"strings"
	"github.com/astaxie/beego/orm"
	"gzyq/logUtils"
)

/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-12-23 23:23
 */

//计算的
type ComAction struct {
	Id   string
	Name string //名字
	Text string //待解析字符串
	Sort string
}
func (action *ComAction)SetId(id string)  {
	action.Id = id
}
func (action *ComAction)SetText(text string) {
	action.Text = text
}
func (action *ComAction)SetName(name string) {
	action.Name = name
}
func (action *ComAction)SetSort(sort string) {
	action.Sort = sort
}
func (action *ComAction)GetSort() string {
	return action.Sort
}
func (action *ComAction)GetText() string {
	return action.Text
}
func (action *ComAction)GetName() string {
	return action.Name
}

//计算部分
func (action *ComAction)DoAction(mMap map[string]string,oneOrm orm.Ormer,onelog logUtils.LogUtils ,sfkqOrm bool) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime error caught: %v", r)
			mMap[oneConfig.ONE_ERROR_NAME] = oneConfig.ONE_ERROR_ACTION + action.Name
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
	fmt.Println("action125,str==", str)
	str = strconv.FormatFloat(Count(str), 'f', -1, 64)
	mMap[action.GetName()] = str
	return nil
}
