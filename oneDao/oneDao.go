package oneDao

import (
	"gzyq/model"
	"gzyq/action"
	"github.com/astaxie/beego/orm"
	"gzyq/rule"
	"errors"
	"strconv"

	"fmt"
)

func GetAllRuleMenu() ([]model.RuleMenu, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	var ruleMenuList []model.RuleMenu
	_, err := o.QueryTable("RuleMenu").Filter("PrientId", "1").OrderBy("Sort").All(&ruleMenuList)
	if err != nil {
		return nil, err
	}

	return ruleMenuList, err
}
/**
 * @Description: 获取子菜单
 * @author : 贾亮
 * @date : 2017/6/19 10:41
 */
func GetAllRuleMenuList(menu model.RuleMenu) ([]model.RuleMenu, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	var rList []model.RuleMenu
	subMenuNum, _ := o.QueryTable("RuleMenu").Filter("PrientId", menu.Id).OrderBy("Sort").All(&rList)
	//下面有子菜单
	if subMenuNum > 0 {
		return rList, nil
	} else {
		//下面没有子菜单
		return nil, errors.New("不存在子菜单")
	}

}




/** 
 * @Description:根据ruleId  加载所有数据
 * @author : 贾亮
 * @date : 2017-06-13 10:50
 */

func GetIndexAlphaByRuleId(ruleId string) (*rule.AlphaNode, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//获取第一个Alpha 节点需要的Key
	var ruleMenu model.RuleMenu
	err := o.QueryTable("RuleMenu").Filter("Id", ruleId).One(&ruleMenu)

	//查所有alpha
	var lists []model.OneAlpha
	num, err := o.QueryTable("OneAlpha").Filter("RuleId", ruleId).All(&lists)
	var alphaMap map[string]model.OneAlpha
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		alphaMap, _ = GetAlphaMap([]model.OneAlpha(lists))
	}

	//查询原子Action
	var subActionlists []model.OneSubAction
	var subActionMap map[string][]action.Action
	_, err = o.QueryTable("OneSubAction").Filter("RuleId", ruleId).All(&subActionlists)
	if err == nil {
		subActionMap, _ = GetSubActionMap(subActionlists)
	}
	//查总action
	var resActionLists []model.OneResAction
	var resActionMap map[string]*action.ResultAction
	_, err = o.QueryTable("OneResAction").Filter("RuleId", ruleId).All(&resActionLists)
	if err == nil {
		resActionMap, _ = GetResActionMap(resActionLists, subActionMap)
	}

	//查询bete List
	var beteList []model.OneBete
	_, err = o.QueryTable("OneBete").Filter("RuleId", ruleId).All(&beteList)

	return getAlphaByMap(alphaMap, ruleMenu.IndexAlpha, resActionMap, beteList), nil
}

func getAlphaByMap(alphaMap map[string]model.OneAlpha, alphaId string, resActionMap map[string]*action.ResultAction, beteList []model.OneBete) *rule.AlphaNode {
	if alphaId == "" {
		return nil
	}
	if oneAlpha, ok := alphaMap[alphaId]; ok {
		symbol, _ := strconv.ParseInt(oneAlpha.Symbol, 10, 8)
		alphaNode := &rule.AlphaNode{oneAlpha.Id, oneAlpha.Name, true, oneAlpha.KeyName, symbol, oneAlpha.Value, resActionMap[oneAlpha.ActionId], getBeteList(beteList, oneAlpha.BeteId, alphaMap, alphaId, resActionMap), getAlphaByMap(alphaMap, oneAlpha.NextId, resActionMap, beteList)}
		return alphaNode
	} else {
		return nil
	}
}

func getActionByMap(resActionMap map[string]*action.ResultAction, id string) *action.ResultAction {
	if id == "" || id == "nil" {
		return nil
	}
	if value, ok := resActionMap[id]; ok {
		return value
	}
	return nil
}
/**
 * @Description: 将alpha 集合 转为map
 * @author : 贾亮
 * @date : 2017/6/12 15:20
 */
func GetAlphaMap(alphaList []model.OneAlpha) (map[string]model.OneAlpha, error) {
	mMap := make(map[string]model.OneAlpha)
	if alphaList != nil {
		for _, alpha := range alphaList {
			mMap[alpha.Id] = alpha
		}
		return mMap, nil
	}
	return nil, errors.New("传入集合为空")
}
func GetResActionMap(resList []model.OneResAction, subMap map[string][]action.Action) (map[string]*action.ResultAction, error) {
	if resList == nil || len(resList) < 1 {
		return nil, errors.New("OneResActionList为空")
	}
	mMap := make(map[string]*action.ResultAction)
	for _, value := range resList {
		mMap[value.Id] = &action.ResultAction{value.Id, value.Name, subMap[value.Id], getBoolByIsReturn(value.IsReturn)}
	}
	return mMap, nil
}
func getBeteList(beteList []model.OneBete, beteId string, alphaMap map[string]model.OneAlpha, alphaId string, resActionMap map[string]*action.ResultAction) ([]*rule.BeteNode) {
	beteNodeList := []*rule.BeteNode{}
	if beteList != nil {
		for _, oneBete := range beteList {
			if oneBete.BeteId == beteId {
				mType, _ := strconv.ParseUint(oneBete.Type, 10, 8)
				beteNodeList = append(beteNodeList, &rule.BeteNode{getAlphaByMap(alphaMap, oneBete.AlphaId, resActionMap, beteList), mType, oneBete.Sort})
			}
		}
		if len(beteNodeList) > 1 {
			SetBeteSort(beteNodeList)
		}
		return beteNodeList
	}
	return nil
}

func getBoolByIsReturn(isReturn string) bool {
	if isReturn == "1" {
		return true
	}
	return false
}
/**
 * @Description:获取原子Action 的map
 * @author : 贾亮
 * @date : 2017/6/12 16:56
 */
func GetSubActionMap(actionList []model.OneSubAction) (map[string][]action.Action, error) {
	mMap := make(map[string][]action.Action)

	if actionList != nil {
		for _, subAction := range actionList {
			if value, ok := mMap[subAction.PrientId]; ok {
				//存在的话就在后面添加 完了排序
				if subAction.Type == "2" {
					value = append(value, &action.ComAction{subAction.Id, subAction.Name, subAction.Text, subAction.Sort})
				} else {
					value = append(value, &action.SubAction{subAction.Id, subAction.Name, subAction.Text, subAction.Sort})
				}
				//如果长度大于1则需要排序
				if len(value) > 1 {
					SetActionSort(value)
				}
				mMap[subAction.PrientId] = value

			} else {
				//不存在则直接创建一个
				if subAction.Type == "2" {
					mMap[subAction.PrientId] = []action.Action{&action.ComAction{subAction.Id, subAction.Name, subAction.Text, subAction.Sort}}
				} else {
					mMap[subAction.PrientId] = []action.Action{&action.SubAction{subAction.Id, subAction.Name, subAction.Text, subAction.Sort}}
				}
			}
		}
		return mMap, nil
	}
	return nil, errors.New("传入集合为空")
}

func SetActionSort(actionList []action.Action) {
	for i := 0; i < len(actionList); i++ {
		for j := 0; j < len(actionList) - i - 1; j++ {
			if actionList[j].GetSort() > actionList[j + 1].GetSort() {
				actionList[j], actionList[j + 1] = actionList[j + 1], actionList[j]
			}
		}
	}
}
func SetBeteSort(nodes []*rule.BeteNode) {
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes) - i - 1; j++ {
			if nodes[j].Sort > nodes[j + 1].Sort {
				nodes[j], nodes[j + 1] = nodes[j + 1], nodes[j]
			}
		}
	}

}

