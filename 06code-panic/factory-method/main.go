package main

import (
	"fmt"
	"log"
)

/*
	名字， 年龄， 身高， 体重， 颜值， 资产， 性别
	human, 使用 marray 方式 封装 幸福指数

	1 性取向有问题， 直接 panic
	2 颜值有问题， 使用 error
	3 返回 幸福指数

*/

// 封装 失败 error 接口
type BadSpouseError struct {
	why string
}

func (bse *BadSpouseError) Error() string {
	return bse.why
}

// BadSpouseError struct factory method, 失败的具体原因
func NewBadSpouseError(o *Human) *BadSpouseError {
	bse := new(BadSpouseError)

	if o.Rmb < 1e6 {
		bse.why = "太穷"
	} else if o.Age > 50 {
		bse.why = "age 太大"
	} else if o.Weight > 200 {
		bse.why = "太胖"
	} else if o.Looking < 70 {
		bse.why = "颜值过低"
	} else {
		return nil
	}

	return bse
}

type Gender int

func (g Gender) String() string {
	return []string{"Male", "Female", "Both"}[g]
}

// const 枚举
const (
	Male = iota
	Female
	Both
)

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
	Rmb    int
	// 自我 取向
	Sex Gender
	// other 去向
	TargetSex Gender
	// 自我 颜值
	Looking int
	// other 颜值
	TargetLooking int
}

func (h *Human) Marray(o *Human) (happiness int, err error) {
	/* 1 性取向有问题， 直接 panic */
	if h.Sex != o.TargetSex {
		// panic("sex fuck off")
		panic(&BadSpouseError{"性别不符合"})
		
	}
	/* 2 颜值有问题， 使用 error */
	// if h.TargetLooking < o.Looking {
	// 	// err = errors.New("颜值过低")
	// 	// err = &BadSpouseError{"颜值过低"}
	// 	err = NewBadSpouseError(h)
	// 	return
	// }

	if e := NewBadSpouseError(o); e != nil {
		err = e
		return
	}

	/* 3 返回 幸福指数 */
	happiness = (o.Height * o.Looking * o.Rmb) / (o.Weight * o.Age)
	return happiness, nil
}

// Human struct factory method
func NewHuman(name string, age, height, weight, rmb, looking, otherlooking int, sex, othersex Gender) *Human {
	h := new(Human)
	h.Name = name
	h.Age = age
	h.Height = height
	h.Weight = weight
	h.Rmb = rmb
	h.Looking = looking
	h.TargetLooking = otherlooking
	h.Sex = sex
	h.TargetSex = othersex

	return h
}

func main() {

	// 恢复恐慌， 阻止整个程序崩溃
	defer func(){
		if err := recover(); err != nil {
			log.Fatal(err)
		}

	}()

	// fmt.Println(time.Now().Month() == 10)

	// cook := NewHuman("cook", 60, 180, 160, 1.23e10, 60, 99, Male, Male)

	girl := NewHuman("girl", 20, 160, 120, 100, 99, 70, Female, Male)

	sicong := NewHuman("sicong", 30, 180, 201, 1.23e10, 80, 99, Male, Female)


	// happiness, err := cook.Marray(girl)
	// happiness, err := girl.Marray(cook)
	happiness, err := girl.Marray(sicong)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("girl happiness: ", happiness)
}
