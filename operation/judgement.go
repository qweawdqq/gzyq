package operation

import (
	"fmt"
	"strconv"
	"strings"
)
/**
 *@Description: 负责判断运算的 接口与实体
 *@author:JiaLiang
 *@date:2017-05-17 20:12
 */

const JUDG_INDEX_NO_EQUAL int64 = 0//不等于
const JUDG_INDEX_EQUAL int64 = 1//等于
const JUDG_INDEX_LESS int64 = 2//小于
const JUDG_INDEX_MORE int64 = 3//大于
const JUDG_INDEX_LESS_EQUAL int64 = 4//大于等于
const JUDG_INDEX_MORE_EQUAL int64 = 5//小于等于
const JUDG_INDEX_WHERE_IN int64 = 6//where in



//0 不等于,1等于 ,2 小于,3 大于,4 小于等于,5 大于等于
var JudgmentArray []Judgment = []Judgment{&NotEqualJudgment{}, &EqualJudgment{}, &LessJudgment{}, &MoreJudgment{}, &LessEqualJudgment{}, &MoreEqualJudgment{},&WhereInJudgment{}}

type Judgment interface {
	Judgment(a, b string) bool
}


//小于
type LessJudgment struct {
}

func (l *LessJudgment)Judgment(a, b string) bool {
	fmt.Println("a=", a, "b=", b, "<")
	fa, erra := strconv.ParseFloat(a, 64)
	fb, errb := strconv.ParseFloat(b, 64)
	if erra != nil||errb != nil {
		//panic("alpha节点待比较参数异常")
		fmt.Println(erra,errb)
	}
	if  fa< fb{
	return true
	}
	return false
}
//大于
type MoreJudgment struct {
}

func (l *MoreJudgment)Judgment(a, b string) bool {

	fmt.Println("a=", a, "b=", b, ">")
	fa, erra := strconv.ParseFloat(a, 64)
	fb, errb := strconv.ParseFloat(b, 64)
	if erra != nil||errb != nil {
		//panic("alpha节点待比较参数异常")
		fmt.Println(erra,errb)
	}
	if  fa> fb{
		return true
	}
	return false
}
//小于等于
type LessEqualJudgment struct {
}

func (l *LessEqualJudgment)Judgment(a, b string) bool {

	fmt.Println("a=", a, "b=", b, "<=")
	fa, erra := strconv.ParseFloat(a, 64)
	fb, errb := strconv.ParseFloat(b, 64)
	if erra != nil||errb != nil {
		//panic("alpha节点待比较参数异常")
		fmt.Println(erra,errb)
	}
	if  fa<= fb{
		return true
	}
	return false
}
//大于等于
type MoreEqualJudgment struct {
}

func (l *MoreEqualJudgment)Judgment(a, b string) bool {
	fmt.Println("a=", a, "b=", b, ">=")
	fa, erra := strconv.ParseFloat(a, 64)
	fb, errb := strconv.ParseFloat(b, 64)
	if erra != nil||errb != nil {
		//panic("alpha节点待比较参数异常")
		fmt.Println(erra,errb)
	}
	if  fa>= fb{
		return true
	}
	return false
}
//不等于
type NotEqualJudgment struct {
}

func (l *NotEqualJudgment)Judgment(a, b string) bool {
	if a != b && a != "" {
		return true
	}
	return false
}
//等于
type EqualJudgment  struct {
}

func (l *EqualJudgment)Judgment(a, b string) bool {
	if a == b || a == "" {
		return true
	}
	return false
}
//where in 的情况
type WhereInJudgment  struct {
}

func (l *WhereInJudgment)Judgment(a, b string) bool {
	if b != "" && a !=""&& strings.Contains(b, ","){
		b =  strings.TrimSpace(b)
		array := strings.Split(b, ",")
                mMap := make(map[string]bool)
		for _, value := range array {
			mMap[value]=true
		}
		if _, ok := mMap[a]; ok {
			return true
		}
		return false
	}

	return false
}







