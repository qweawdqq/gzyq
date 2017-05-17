package operation
/**
 *@Description: 负责判断运算的 接口与实体
 *@author:JiaLiang
 *@date:2017-05-17 20:12
 */
//0 不等于,1等于 ,2 小于,3 大于,4 小于等于,5 大于等于
var JudgmentArray []Judgment = []Judgment{&NotEqualJudgment{}, &EqualJudgment{}, &LessJudgment{}, &MoreJudgment{}, &LessEqualJudgment{}, &MoreEqualJudgment{}}

type Judgment interface {
	Judgment(a, b string) bool
}


//小于
type LessJudgment struct {
}

func (l *LessJudgment)Judgment(a, b string) bool {
	if a < b {
		return true
	}
	return false
}
//大于
type MoreJudgment struct {
}

func (l *MoreJudgment)Judgment(a, b string) bool {
	if a > b {
		return true
	}
	return false
}
//小于等于
type LessEqualJudgment struct {
}

func (l *LessEqualJudgment)Judgment(a, b string) bool {
	if a <= b {
		return true
	}
	return false
}
//大于等于
type MoreEqualJudgment struct {
}

func (l *MoreEqualJudgment)Judgment(a, b string) bool {
	if a >= b {
		return true
	}
	return false
}
//不等于
type NotEqualJudgment struct {
}

func (l *NotEqualJudgment)Judgment(a, b string) bool {
	if a != b {
		return true
	}
	return false
}
//等于
type EqualJudgment  struct {
}

func (l *EqualJudgment)Judgment(a, b string) bool {
	if a == b {
		return true
	}
	return false
}






