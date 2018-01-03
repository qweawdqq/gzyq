package logUtils

import "time"



type LogUtils interface {
 GetSfNeedRz() bool
 GetSfOpenLct() bool
 AddRz(info string)
 GetRz() []string
	SetSfNeedRz(sfOpen bool)
}
/** 
 * @Description:日志工具类
 * @author : 贾亮
 * @date : 2017-12-24 18:20
 */

type LogUtil struct {
	RequestFw       string //请求的服务
	RequestUser       string//请求的人
	RequestTime   time.Time //请求时间
	RuleID       string  //请求的规则ID
	RuleName        string//请求的规则名称
	Lct []string//流程图
	SfOpenLct bool //是否开启流程图
	SfNeedRz bool //是否记录日志
	Rz []string//详细的日志

}

func (this *LogUtil)SetSfNeedRz(sfOpen bool){
	this.SfNeedRz = sfOpen
}
/**
 * @Description:获取是否要记录日志
 * @author : 贾亮
 * @date : 2017-12-24 18:45
 */
func (this *LogUtil)GetSfNeedRz() bool {
	return this.SfNeedRz
}

/**
 * @Description:获取是否要开启流程图
 * @author : 贾亮
 * @date : 2017-12-24 18:45
 */
func (this *LogUtil)GetSfOpenLct() bool {
	return this.SfOpenLct
}
/**
 * @Description:添加日志
 * @author : 贾亮
 * @date : 2017-12-24 18:47
 */
func (this *LogUtil)AddRz(info string)  {
	this.Rz = append(this.Rz,info)
}
/**
 * @Description:获取日志
 * @author : 贾亮
 * @date : 2017-12-24 18:53
 */
func (this *LogUtil) GetRz() []string{
	return this.Rz

}