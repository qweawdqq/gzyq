package model

/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-06-12 16:05
 */

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func init()  {
	orm.RegisterModel(new(OneAlpha), new(OneBete), new(OneSubAction), new(OneResAction), new(RuleMenu))//注册表studentinfo 如果没有会自动创建
}


/**
 * @Description: 数据表
 * @author : 贾亮
 * @date : 2017-06-12 13:23
 */
type OneAlpha struct {
	Id       string `orm:"pk"`
	Name     string
	KeyName  string
	Symbol   string
	Value    string
	ActionId string
	BeteId   string
	NextId   string
	RuleId   string
}
type OneBete struct {
	Id      string `orm:"pk"`
	Type    string
	AlphaId string
	RuleId  string
	BeteId  string
	Sort    string
}
type OneSubAction struct {
	Id       string `orm:"pk"`
	PrientId string
	Name     string
	Text     string
	Type     string
	Sort     string
	RuleId   string
}
type OneResAction struct {
	Id       string `orm:"pk"`
	Name     string
	IsReturn string
	RuleId   string
}

type RuleMenu struct {
	Id         string `orm:"pk"`
	Name       string
	Info       string
	PrientId   string
	Sort       string
	Url        string
	IndexAlpha string
}

