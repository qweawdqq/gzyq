package main

import (
	"fmt"
	"gzyq/operation"
	"gzyq/action"
	"gzyq/rule"
	"gzyq/oneConfig"
)

func main() {
	//0 不等于,1等于 ,2 小于,3 大于,4 小于等于,5 大于等于
	judgment := operation.JudgmentArray[operation.JUDG_INDEX_LESS]
	mBool := judgment.Judgment("123.30", "123.00")
	fmt.Println(mBool)
	fmt.Println(len(operation.JudgmentArray))

	//截取字符串
	//tesr1 := "$xxx$<file name=\"text\">$mm$</file>$dd$$fffsss$"
	//mMap := make(map[string]string)
	//mMap["xxx"] = "这是1"
	//mMap["mm"] = "这是2"
	//mMap["dd"] = "这是3"
	//mMap["fffsss"] = "这是4"
	//a := &action.SubAction{}
	//a.SetText(tesr1)
	//a.SetName("action1")
	//err :=a.DoAction(mMap)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if v, ok := mMap[a.GetName()]; ok {
	//	fmt.Println(v)
	//}else {
	//fmt.Println("没发现值")
	//}
	//测试计算
	//截取字符串
	//tesr1 := "$aaa$*5.55533333+($bbb$-$ccc$)/$ddd$"
	//mMap := make(map[string]string)
	//mMap["aaa"] = "10"
	//mMap["bbb"] = "20"
	//mMap["ccc"] = "50"
	//mMap["ddd"] = "5"
	//action := &action.ComAction{}
	//action.SetText(tesr1)
	//action.SetName("hehe")
	//err := action.DoAction(mMap)
	//if err != nil{
	//	fmt.Println(err)
	//}else {
	//	if v, ok := mMap[action.GetName()]; ok {
	//		fmt.Println(v)
	//	} else {
	//		fmt.Println("没发现值")
	//	}
	//}

	//name            string                //名字
	//value           bool                  //值
	//operationSymbol int8                  //运算符号
	//operationValue  string                //参与运算的值
	//reAction        action.ResultAction //返回值action
	//beteList  BeteList //bete节点集合
	//nextAlpha AlphaNode // 下一个alpha节点
	mMap := map[string]string{
		"k1": "4",
		"k2": "3000",
	}
	//id string
	//name string
	//actionList []Action
	//isReturn  bool //是向下还是返回  返回为true
	action1 := &action.ResultAction{"001", "action1", []action.Action{&action.SubAction{"sub1", "k1<5年"}}, true}
	action2 := &action.ResultAction{"002", "action2", []action.Action{&action.ComAction{"com1", "$k2$*1.2"}}, true}
	action3 := &action.ResultAction{"003", "action3", []action.Action{&action.ComAction{"com2", "$k2$*1.5"}}, true}
	action4 := &action.ResultAction{"004", "action4", []action.Action{&action.ComAction{"com4", "$k2$*1.3"}}, true}
	action5 := &action.ResultAction{"005", "action5", []action.Action{&action.ComAction{"com5", "$k2$*2.0"}}, true}
	alpha6 := &rule.AlphaNode{"alpha6", true, "k1", operation.JUDG_INDEX_MORE, "15", action5, nil, nil}
	alpha5 := &rule.AlphaNode{"alpha5", true, "k2", operation.JUDG_INDEX_LESS_EQUAL, "4000", action4, nil, nil}
	alpha4 := &rule.AlphaNode{"alpha4", true, "k2", operation.JUDG_INDEX_LESS_EQUAL, "2000", action3, nil, nil}
	rule1 := &rule.BeteNode{alpha4, 1}
	rule2 := &rule.BeteNode{alpha5, 1}
	beteList := []*rule.BeteNode{}
	beteList = append(beteList, rule1, rule2)
	alpha3 := &rule.AlphaNode{"alpha3", true, "k1", operation.JUDG_INDEX_LESS_EQUAL, "15", nil, beteList, alpha6}
	alpha2 := &rule.AlphaNode{"alpha2", true, "k1", operation.JUDG_INDEX_LESS, "10", action2, nil, alpha3}
	alpha1 := &rule.AlphaNode{"alpha1", true, "k1", operation.JUDG_INDEX_LESS, "5", action1, nil, alpha2}
	hehe := alpha1.DoAction(true, mMap)

	if v, ok := mMap[oneConfig.ONE_ERROR_MSG]; ok {
		fmt.Println("错误信息", v)
		fmt.Println("错误Name", mMap[oneConfig.ONE_ERROR_NAME])
	} else {
		fmt.Println("hehe", hehe)
	}

	fmt.Println(mMap)

}
