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
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}else {
	fmt.Println("goroutine ID",id)
	}

	//--------------------------
	start := time.Now()
	fmt.Println("-->FunCall request:", paramMap)
	ruleId := paramMap["ruleId"]
	alpha, _ := oneBo.GetIndexAlphaByRuleId(ruleId)
	str := alpha.DoAction(true, paramMap)
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
	r = paramMap
	fmt.Println("str==", str)
	dis := time.Since(start)
	fmt.Println("耗时", dis)
	return

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

	r,err = oneBo.GetAllRuleMenu()
	fmt.Println("菜单里面的东西",r)

	return
}

func (service *OneServiceImpl)GetRuleJsonByRuleId(id string) (r string, err error){

	r,err =oneBo.GetRuleJsonByRuleId(id)
	return
}