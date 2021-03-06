package action

import (
	"strings"
	"gzyq/oneConfig"
	"gzyq/logUtils"
	"github.com/astaxie/beego/orm"
	"fmt"
)

/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-12-23 23:23
 */

//摘要的
type SubAction struct {
	Id   string
	Name string //名字
	Text string //待解析字符串
	Sort string
}

func (action *SubAction)SetId(id string)  {
	action.Id = id
}
func (action *SubAction)SetText(text string) {
	action.Text = text
}
func (action *SubAction)SetName(name string) {
	action.Name = name
}
func (action *SubAction)SetSort(sort string) {
	action.Sort = sort
}
func (action *SubAction)GetSort() string {
	return action.Sort
}
func (action *SubAction)GetText() string {
	return action.Text
}
func (action *SubAction)GetName() string {
	return action.Name
}
func (action *SubAction)DoAction(mMap map[string]string,oneOrm orm.Ormer,onelog logUtils.LogUtils ,sfkqOrm bool) error {
	//var str = ""
	//if strings.Contains(action.GetText(), "$") {
	//	str1, err := getSubAction(action.GetText(), mMap)
	//	if err != nil {
	//		mMap[oneConfig.ONE_ERROR_MSG] = err.Error()
	//		mMap[oneConfig.ONE_ERROR_NAME] = oneConfig.ONE_ERROR_ACTION + action.Name
	//		return err
	//	}
	//	fmt.Println("str1==",str1)
	//	str = str1
	//} else {
	//	str = action.GetText()
	//}
	str , _ := GetSubText(action.Text,mMap)
	mMap[action.GetName()] = str
	return nil

}

func GetSubText(strx string,mMap map[string]string)(string, error) {
	var str = ""
	if strings.Contains(strx, SubKey) {
		str1, err := getSubAction(str, mMap)
		if err != nil {
			mMap[oneConfig.ONE_ERROR_MSG] = err.Error()
			mMap[oneConfig.ONE_ERROR_NAME] = oneConfig.ONE_ERROR_ACTION+ strx
			return str,err
		}
		fmt.Println("str1==", str1)
		str = str1
	} else {
		str = strx
	}
return str,nil
}

