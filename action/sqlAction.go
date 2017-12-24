package action

import (
	"github.com/astaxie/beego/orm"
	"gzyq/logUtils"
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
func (action *SqlAction)SetId(id string)  {
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
func (action *SqlAction)DoAction(mMap map[string]string,oneOrm orm.Ormer,onelog logUtils.LogUtils ,sfkqOrm bool,sfkqLog bool) error {
	return nil
}
