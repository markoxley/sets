package main

import (
	"fmt"
	"sets/sets"
	"time"
)

func main() {
	s1 := sets.New("C", "C++", "DBMS", "Java", true, 3, 5, time.Now())
	s2 := sets.New("DBMS", "Data Science", "Java", "Machine Learning", "Python")
	fmt.Println(s1.Difference(s2).Elements())
	fmt.Println(s2.Difference(s1).Elements())
	fmt.Println(s1.SymmetricDifference(s2))

	t := "hello"
	t2 := &t

	fmt.Printf("t is %T", t)
	fmt.Printf("t2 is %T", t2)
}
