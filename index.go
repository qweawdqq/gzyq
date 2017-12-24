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