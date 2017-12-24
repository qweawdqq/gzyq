package action

import (
	"gzyq/oneConfig"
	"fmt"
)


/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-12-23 23:12
 */
func Factory(actionType string) Action {
	fmt.Println("action类型",actionType)
	if  oneConfig.ACTION_TYPE_COM ==  actionType{
		return  &ComAction{}
	}else if  oneConfig.ACTION_TYPE_SUB ==  actionType {
		return  &SubAction{}
	}else if  oneConfig.ACTION_TYPE_SQL ==  actionType{

		return  &SqlAction{}
	}else {
		return nil
	}
}
