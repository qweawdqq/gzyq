package rule

import (
	"gzyq/operation"
	"gzyq/action"
	_"strconv"

	"gzyq/logUtils"
	"github.com/astaxie/beego/orm"
)

type RuleNode struct {
	Id    string
	Alpha *AlphaNode

}



/** 
 * @Description:rule  alphaNode beteNode
 * @author : 贾亮
 * @date : 2017-05-23 10:19
 */

type AlphaNode struct {
	Id              string
	Name            string               //名字
	Value           bool                 //值
	OpertionKey     string               //参与运算的值
	OperationSymbol int64                //运算符号
	OperationValue  string               //参与运算的值
	ReAction        *action.ResultAction //返回值action
	BeteList        []*BeteNode          //bete节点集合
	NextAlpha       *AlphaNode           // 下一个alpha节点
}
/**
 * @Description:alpha 的执行总过程
 * isNeedCompute  是否需要重新计算
 * @author : 贾亮
 * @date : 2017/6/7 15:55
 */
func (alpha *AlphaNode)DoAction(isNeedCompute bool, mMap map[string]string,oneOrm orm.Ormer,log logUtils.LogUtils ,sfkqOrm bool,sfkqLog bool) string {
	//fmt.Println("alpha.name=", alpha.Name)
	//if alpha.ReAction != nil {
	//fmt.Println("alpha.ReAction.Name", alpha.ReAction.GetName())
	//}

	//计算 赋值
	if isNeedCompute {
		alpha.GetNewValue(mMap)
		//b1 := alpha.GetNewValue(mMap)
		//fmt.Println("b1==", strconv.FormatBool(b1))
	}
	alphaValue := alpha.GetValue()

	// 走action
	if alphaValue == true&&alpha.ReAction != nil {
		alpha.ReAction.DoAction(mMap,oneOrm,log,sfkqOrm)
		if alpha.ReAction.GetIsReturn() {
			return alpha.ReAction.GetResultStr(mMap)
		}
	}
	//走bete
	//bList := alpha.BeteList
	if alpha.BeteList != nil {
		//fmt.Println("bete节点不为空")
		for _, v := range alpha.BeteList {
			//fmt.Println("v.Alpha.Name-----", v.Alpha.Name)
			v.Alpha.GetNewValue(mMap)
			//如果为true 则向下走
			alphaBool := v.GetConnBool(alpha)
			//fmt.Println("alphaBool", strconv.FormatBool(alphaBool))
			if alphaBool {
				v.Alpha.SetValue(alphaBool)
				reStr := v.Alpha.DoAction(false, mMap,oneOrm,log,sfkqOrm,sfkqLog)
				if reStr != "" {
					return reStr
				}
			}

		}
	}
	if alpha.NextAlpha == nil {
		return ""
	}
	//走alpha
	return alpha.NextAlpha.DoAction(true, mMap,oneOrm,log,sfkqOrm,sfkqLog)
}

func (alpha *AlphaNode)SetBeteList() {

}

func (alpha *AlphaNode) GetValue() bool {
	return alpha.Value
}
func (alpha *AlphaNode) SetValue(value bool) {
	alpha.Value = value
}
func (alpha *AlphaNode)  SetValueByStr(mMap map[string]string) {
	op := mMap[alpha.OpertionKey]
	//fmt.Println("op", op)
	//fmt.Println("Symbol", alpha.OperationSymbol)
	//fmt.Println("Value", alpha.OperationValue)
	alpha.Value = operation.JudgmentArray[alpha.OperationSymbol].Judgment(op, alpha.OperationValue)
}
func (alpha *AlphaNode) GetNewValue(mMap map[string]string) bool {
	alpha.SetValueByStr(mMap)
	return alpha.GetValue()
}

type BeteNode struct {
	Alpha    *AlphaNode
	BeteType uint64 //连接符号
	Sort     string
}

const BETE_CONN_ADD uint64 = 1
const BETE_CONN_OR uint8 = 2
/**
 * @Description:连接两个Alpha节点  获取值
 * @author : 贾亮
 * @date : 2017/6/7 16:20
 */
func (beteNode *BeteNode)GetConnBool(alpha *AlphaNode) bool {
	//fmt.Println("运算alpha名称", alpha.Name)
	//fmt.Println("连接类型", beteNode.BeteType)
	//fmt.Println(beteNode.GetAlpha().Name)
	if beteNode.BeteType == BETE_CONN_ADD {
		return beteNode.Alpha.GetValue()&& alpha.Value
	} else {
		return beteNode.Alpha.GetValue() || alpha.Value
	}
}
func (beteNode *BeteNode)GetAlpha() *AlphaNode {
	return beteNode.Alpha
}

func (beteNode *BeteNode)GetIsResult() bool {
	return beteNode.Alpha.ReAction.GetIsReturn()
}

func (beteNode *BeteNode) DoAction(mMap map[string]string,oneOrm orm.Ormer,onelog logUtils.LogUtils ,sfkqOrm bool) {
	beteNode.Alpha.ReAction.DoAction(mMap,oneOrm,onelog,sfkqOrm)
}


