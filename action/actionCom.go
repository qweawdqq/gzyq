package action

import "github.com/shopspring/decimal"


/** 
 * @Description:具体负责计算的接口
 * @author : 贾亮
 * @date : 2017-12-03 16:18
 */


var(
	ComMap = map[string]ComInterface{
		"+": &AddCom{},
		"-": &DelCom{},
		"*": &MulCom{},
		"/": &DivCom{},
}

)

type ComInterface interface {
	DoCom(a, b float64) float64
}

type AddCom struct {
	a float64
	b float64
}
func (l *AddCom)DoCom(a, b float64) float64 {
	a1 := decimal.NewFromFloat(a)
	b1 := decimal.NewFromFloat(b)
	c1 ,_ :=a1.Add(b1).Float64()
	return c1
}

type DelCom struct {
	a float64
	b float64
}
func (l *DelCom)DoCom(a, b float64) float64 {
	a1 := decimal.NewFromFloat(a)
	b1 := decimal.NewFromFloat(b)
	c1 ,_ :=a1.Sub(b1).Float64()
	return c1
}

type MulCom struct {
	a float64
	b float64
}
func (l *MulCom)DoCom(a, b float64) float64 {
	a1 := decimal.NewFromFloat(a)
	b1 := decimal.NewFromFloat(b)
	c1 ,_ :=a1.Mul(b1).Float64()
	return c1
}

type DivCom struct {
	a float64
	b float64
}
func (l *DivCom)DoCom(a, b float64) float64 {
	a1 := decimal.NewFromFloat(a)
	b1 := decimal.NewFromFloat(b)
	c1 ,_ :=a1.Div(b1).Float64()
	return c1
}