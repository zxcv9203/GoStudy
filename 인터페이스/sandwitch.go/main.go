package main

import "fmt"

//"reflect"

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	getOneSpoon() SpoonOfJam
}

type Bread struct {
	value string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange"
}

func (s *SpoonOfStrawberryJam) String() string {
	//fmt.Println("Strawberry")
	return " + strawberry"
}

func (b *Bread) String() string {
	//fmt.Println("Bread")
	return "bread" + b.value
}

//func (j *StrawberryJam) getOneSpoon() *SpoonOfStrawberryJam {
//	return &SpoonOfStrawberryJam{}
//}

//func (j *OrangeJam) getOneSpoon() *SpoonOfOrangeJam {
//	return &SpoonOfOrangeJam{}
//}

func (j *StrawberryJam) getOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) getOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.getOneSpoon()
	b.value += spoon.String()
}

//func (b *Bread) PutJamStraw(jam *StrawberryJam) {
//	spoon := jam.getOneSpoon()
//	b.value += spoon.String()
//}

//func (b *Bread) PutJamOrange(jam *OrangeJam) {
//	spoon := jam.getOneSpoon()
//	b.value += spoon.String()
//}

func main() {
	bread := &Bread{}
	//jam := StrawberryJam{}
	//bread.PutJamStraw(&jam)
	//fmt.Println(bread)

	////jam = OrangeJam{} // ERROR
	jam2 := OrangeJam{}
	bread.PutJam(&jam2)
	fmt.Println(bread)

	//newBread := &Bread{}
	//newBread.PutJam(&jam2)
	//fmt.Println(newBread)

	//var jam *SpoonOfOrangeJam
	//jam = &SpoonOfOrangeJam{}

	//fmt.Println((jam))
	//jam.String()

}
