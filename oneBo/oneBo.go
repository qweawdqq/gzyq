package oneBo

import (
	"gzyq/rule"
	"gzyq/onecache"
	"fmt"
	"time"
	"gzyq/oneDao"
	"sync"
	demo "github.com/demo/rpc"
	"encoding/json"
)
var cache *onecache.CacheTable
var lock sync.RWMutex
func init() {
	cache = onecache.Cache("RULE")
	fmt.Println("创建了缓存RULE")
}
/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-06-14 08:41
 */
func GetIndexAlphaByRuleId(ruleId string) (*rule.AlphaNode,error) {
	res, err := cache.Value(ruleId)
	if err != nil {
		lock.Lock()
		defer lock.Unlock()
		res, err = cache.Value(ruleId)
		if err == nil{
			return res.GetData().(*rule.AlphaNode),nil
		}
		fmt.Println("缓存内存地址",cache)
		fmt.Println("缓存错误",err)
		fmt.Println("向缓存添加rule数据",ruleId)
		alpha,err := oneDao.GetIndexAlphaByRuleId(ruleId)
		if err != nil{
			fmt.Println("查找rule失败")
			return nil,err
		}

		fmt.Println("向缓存中添加了数据",alpha)
		//缓存时间15分钟
		cache.Add(ruleId,onecache.CACHE_TIME*time.Minute,alpha)
		return alpha,nil
	}
	fmt.Println("缓存里面有数据",res.GetData().(*rule.AlphaNode))
	//
	//aaa,_:=json.Marshal(res.GetData().(*rule.AlphaNode))
	//fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa",string(aaa))
	return res.GetData().(*rule.AlphaNode),nil
}

func GetRuleJsonByRuleId(id string) (r string, err error) {
	fmt.Println("请求的id为="+id)
	alpha,err:=GetIndexAlphaByRuleId(id)
	fmt.Println(alpha)
	if err != nil{
		return "",err
	}
	json,err:=json.Marshal(alpha)
	fmt.Println("GetRuleJsonByRuleId",string(json))
	if err != nil{
		return "",err
	}
	return string(json),err
}


func GetAllRuleMenu()([]*demo.PrientMenu, error)  {

	//var allMenuList []oneDao.AllMenu
	allMenuList,err := oneDao.GetAllRuleMenu()
	if err != nil{
		return nil,err
	}
	var pMenuList []*demo.PrientMenu
	for index,v :=range allMenuList  {
		fmt.Println(index)
		subMenuList,err:=oneDao.GetAllRuleMenuList(v)
		var subList []*demo.SubMenu

		if err == nil{
			for _,oneSubMenu := range subMenuList{
				subList =  append(subList,&demo.SubMenu{oneSubMenu.Id,oneSubMenu.Name,oneSubMenu.Info,oneSubMenu.PrientId,oneSubMenu.Sort,oneSubMenu.Url,oneSubMenu.IndexAlpha})
			}
			pp:=&demo.SubMenu{v.Id,v.Name,v.Info,v.PrientId,v.Sort,v.Url,v.IndexAlpha}
			pMenu:= &demo.PrientMenu{pp,subList}
			pMenuList = append(pMenuList,pMenu)

		}else {
			pp:=&demo.SubMenu{v.Id,v.Name,v.Info,v.PrientId,v.Sort,v.Url,v.IndexAlpha}
			pMenu:= &demo.PrientMenu{pp,nil}
			pMenuList = append(pMenuList,pMenu)
		}

	}
	return pMenuList,nil
}