package oneRpc

import (
	"fmt"
	"time"
	"gzyq/oneConfig"
	"gzyq/oneBo"
	demo "github.com/demo/rpc"
	"runtime"
	"strings"
	"strconv"
	"github.com/astaxie/beego/orm"
	"gzyq/logUtils"
)

/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-06-16 12:26
 */

type OneServiceImpl struct {

}

//  - ParamMap
func (service *OneServiceImpl)GetResultByRuleId(paramMap map[string]string) (r map[string]string, err error) {
	//--获取协程ID
	getGoroutineId();
	//--------------------------
	start := time.Now()
	fmt.Println("-->FunCall request:", paramMap)
	ruleId := paramMap["ruleId"]
	//是否开启事务
	oneOrm :=getOrmByBool(true)
	//是否开启日志
	log:= getLogByBool(true)
	alpha, _ := oneBo.GetIndexAlphaByRuleId(ruleId)
	str := alpha.DoAction(true, paramMap,oneOrm,log,true,true)
	setErrReturn(paramMap, str)
	r = paramMap
	fmt.Println("str==", str)
	//提交事务
	setOrmCommit(oneOrm,r)
	dis := time.Since(start)
	fmt.Println("耗时", dis)
	return
}

func getLogByBool(sfkqrz bool)logUtils.LogUtils {
	oneLog := new(logUtils.LogUtil)

	if sfkqrz {
		oneLog.SfNeedRz = sfkqrz
	}
	return oneLog
}

//如发生错误事务回滚 否则提交
func setOrmCommit(oneOrm orm.Ormer,paramMap map[string]string)  {
	if _, ok := paramMap[oneConfig.ONE_ERROR_MSG]; ok {
		if  paramMap[oneConfig.ONE_MSG] == oneConfig.ONE_MSG_ERR{
			if oneOrm != nil{
				oneOrm.Rollback()
				fmt.Println("回滚了事务")
			}
		}
	}
	if oneOrm != nil{
		oneOrm.Commit()
		fmt.Println("提交了事务")

	}

}
/**
 * @Description:是否开启事务的方法
 * @author : 贾亮
 * @date : 2017-12-24 19:30
 */
func getOrmByBool(sfkq bool) orm.Ormer {
	if sfkq {
		o := orm.NewOrm()
		o.Begin()
		fmt.Println("开启了事务")
		return o
	}
	return nil

}

/**
 * @Description:设置返回或异常信息
 * @author : 贾亮
 * @date : 2017-12-24 19:00
 */
func setErrReturn(paramMap map[string]string, str string) {
	if _, ok := paramMap[oneConfig.ONE_ERROR_MSG]; ok {
		paramMap[oneConfig.ONE_MSG] = oneConfig.ONE_MSG_ERR
	} else {
		paramMap[oneConfig.ONE_MSG] = oneConfig.ONE_MSG_OK
		if str == "" {
			paramMap[oneConfig.ONE_MSG_INFO] = "未能获取返回值，有可能未发现满足条件"
		} else {
			paramMap[oneConfig.ONE_MSG_INFO] = str
		}
	}
}

func getGoroutineId() {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	} else {
		fmt.Println("goroutine ID", id)
	}
}

// Parameters:
//  - JSON
func (service *OneServiceImpl)UpDateRule(json string) (r map[string]string, err error) {
	fmt.Println("update==", json)
	r = make(map[string]string)
	r["1111"] = "1111"
	return
}

func (service *OneServiceImpl) GetRuleMenu(json string) (r []*demo.PrientMenu, err error) {

	r, err = oneBo.GetAllRuleMenu()
	fmt.Println("菜单里面的东西", r)

	return
}

func (service *OneServiceImpl)GetRuleJsonByRuleId(id string) (r string, err error) {

	r, err = oneBo.GetRuleJsonByRuleId(id)
	return
}