package heihei

import (
	"fmt"
	"reflect"
)

type User struct {
	Name        string
	Age         int
	Description string
}

func LearnReflect() {
	u1 := &User{"tom", 5, "A loser indeed :-D"}

	o := reflect.ValueOf(u1)

	ref := o.Elem()

	t := ref.Type()

	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)

		fmt.Printf("%d. %s %s = %v \n", i, t.Field(i).Name, f.Type(), f.Interface())
	}
}
