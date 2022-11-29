package main

import "fmt"

func main() {
	list1 := List{}

	// list1.addValue(12)
	// list1.addValue(1)
	// list1.addValue(3)
	// list1.addValue(123)

	list1.addMultipleVal()

	list1.printAllVal()
}

type node struct {
	data int
	next *node
}

type List struct {
	head *node
	tail *node
}

func (l *List) addMultipleVal() {
	//get limit from user
	var limit int
	fmt.Print("Enter lmit for list: ")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		var val int
		fmt.Print("Enter value: ")
		fmt.Scanf("%d", &val)
		l.addValue(val)
	}
}

func (l *List) addValue(data int) {
	temp := node{
		data: data,
	}

	if l.head == nil {
		l.head = &temp
	} else {
		l.tail.next = &temp
	}
	l.tail = &temp
}

func (l *List) printAllVal() {

	temp := l.head

	for temp != nil {
		fmt.Print((*temp).data, " ")
		temp = temp.next
	}
	fmt.Println()
}
