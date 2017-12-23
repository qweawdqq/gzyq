package main

import (
	"github.com/astaxie/beego/orm"
	_"gzyq/model"

	"fmt"
	"os"
	"github.com/thrift"
	oneRpcs "github.com/demo/rpc"

	"gzyq/oneRpc"
)
//数据库连对象需要的信息
var (
	dbuser string = "qweawdqq"
	dbpwd string = "qweawdqq1"
	dbname string = "onework"
)

func init() {
	conn := dbuser + ":" + dbpwd + "@/" + dbname + "?charset=utf8"//组合成连接串
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", conn) //设置conn中的数据库为默认使用数据库
	orm.SetMaxIdleConns("default", 30)                                       //设置数据库最大空闲连接
	orm.SetMaxOpenConns("default", 30)                                       //设置数据库最大连接数
	//orm.RunSyncdb("default", false, true)//后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
	orm.Debug = true
}
func main() {
	//alpha, err := oneBo.GetIndexAlphaByRuleId("c9028fdadba3420996161781e8f20258")
	//fmt.Println("alpha11111", alpha)
	//
	//alpha, err = oneBo.GetIndexAlphaByRuleId("c9028fdadba3420996161781e8f20258")
	//fmt.Println("alpha222", alpha)
	//
	//alpha, err = oneBo.GetIndexAlphaByRuleId("c9028fdadba3420996161781e8f20258")
	//fmt.Println("alpha333", alpha)
	//
	//if alpha == nil {
	//	return
	//}
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(alpha)
	//}
	//mMap := map[string]string{
	//	"k1": "8",
	//	"k2": "3000",
	//}
	//hehe := alpha.DoAction(true, mMap)
	////0 不等于,1等于 ,2 小于,3 大于,4 小于等于,5 大于等于
	//judgment := operation.JudgmentArray[operation.JUDG_INDEX_LESS]
	//mBool := judgment.Judgment("123.30", "123.00")
	//fmt.Println(mBool)
	//fmt.Println(len(operation.JudgmentArray))
	////nextAlpha AlphaNode // 下一个alpha节点
	////mMap := map[string]string{
	////	"k1": "18",
	////	"k2": "8000",
	////}
	////id string
	////name string
	////actionList []Action
	////isReturn  bool //是向下还是返回  返回为true
	////action1 := &action.ResultAction{"001", "action1", []action.Action{&action.SubAction{"id1","sub1", "k1<5年","1"}}, true}
	////action2 := &action.ResultAction{"002", "action2", []action.Action{&action.ComAction{"id1","com1", "$k2$*1.2","1"}}, true}
	////action3 := &action.ResultAction{"003", "action3", []action.Action{&action.ComAction{"id1","com2", "$k2$*1.5","1"}}, true}
	////action4 := &action.ResultAction{"004", "action4", []action.Action{&action.ComAction{"id1","com4", "$k2$*1.3","1"}}, true}
	////action5 := &action.ResultAction{"005", "action5", []action.Action{&action.ComAction{"id1","com5", "$k2$*2.0","1"}}, true}
	////alpha6 := &rule.AlphaNode{"qweqweqwe","alpha6", true, "k1", operation.JUDG_INDEX_MORE, "15", action5, nil, nil}
	////alpha5 := &rule.AlphaNode{"qweqweqwe","alpha5", true, "k2", operation.JUDG_INDEX_LESS_EQUAL, "4000", action4, nil, nil}
	////alpha4 := &rule.AlphaNode{"qweqweqwe","alpha4", true, "k2", operation.JUDG_INDEX_LESS_EQUAL, "2000", action3, nil, nil}
	////rule1 := &rule.BeteNode{alpha4, 1,"1"}
	////rule2 := &rule.BeteNode{alpha5, 1,"2"}
	////beteList := []*rule.BeteNode{}
	////beteList = append(beteList, rule1, rule2)
	////alpha3 := &rule.AlphaNode{"qweqweqwe","alpha3", true, "k1", operation.JUDG_INDEX_LESS_EQUAL, "15", nil, beteList, alpha6}
	////alpha2 := &rule.AlphaNode{"qweqweqwe","alpha2", true, "k1", operation.JUDG_INDEX_LESS, "10", action2, nil, alpha3}
	////alpha1 := &rule.AlphaNode{"qweqweqwe","alpha1", true, "k1", operation.JUDG_INDEX_LESS, "5", action1, nil, alpha2}
	////hehe := alpha1.DoAction(true, mMap)
	//////
	//if v, ok := mMap[oneConfig.ONE_ERROR_MSG]; ok {
	//	fmt.Println("错误信息", v)
	//	fmt.Println("错误Name", mMap[oneConfig.ONE_ERROR_NAME])
	//} else {
	//	if hehe == "" {
	//		fmt.Println("hehe", "未能获取返回值，有可能未发现满足条件")
	//	} else {
	//		fmt.Println("hehe", hehe)
	//	}
	//}
	//
	//fmt.Println(mMap)

	//测试Thrif--------------------------------------------------

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)

	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &oneRpc.OneServiceImpl{}
	processor := oneRpcs.NewOneServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()

}

const (

	NetworkAddr = "127.0.0.1:8082"
)

//type RpcServiceImpl struct {
//}
//
//func (this *RpcServiceImpl) FunCall(callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
//	start := time.Now()
//	fmt.Println("-->FunCall:", callTime, funCode, paramMap)
//	ruleId := paramMap["ruleId"]
//	alpha, _ := oneBo.GetIndexAlphaByRuleId(ruleId)
//	hehe := alpha.DoAction(true, paramMap)
//	fmt.Println("hehe==", hehe)
//	r= append(r,hehe)
//	dis := time.Since(start)
//	fmt.Println("耗时", dis)
//	return
//}