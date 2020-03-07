package main

import (
	"fmt"
)

/**
函数版的学生管理系统
1 查看所有的学生
2 新增学生
3 删除学生
*/

type student struct {
	id   int64
	name string
}

//变量声明
var (
	allStudent map[int64]*student // * 表示指针类型的student
)

func text() {
	fmt.Println("--------------MENU---------------")
	fmt.Println(`
1 查看所有的学生
2 新增学生
3 删除学生`)
	fmt.Println("等待输入.,按0退出系统 ......")
}

//通过构造函数生成一个学生
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for _, v := range allStudent {
		fmt.Printf("学号：%d \t姓名：%v\n", v.id, v.name)
	}
	fmt.Println("-----显示完成-----")
}

func addStudent() {
	//创建一个学生
	var (
		id   int64
		name string
	)
	fmt.Print("学号：")
	fmt.Scanln(&id)

	fmt.Print("姓名")
	fmt.Scanln(&name)
	//添加到当前的班级
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	var (
		deleteID int64
	)
	// 1 先输入学生的序号
	fmt.Print("输入学号： ")
	// 2 输入学生的ID
	fmt.Scanln(&deleteID)
	// 3 删除
	delete(allStudent, deleteID)
	showAllStudent()
}
func main() {
	allStudent = make(map[int64]*student, 48)
	fmt.Println("系统初始化成功")
	///1 打印菜单
	text()
	var inputNumber int
	fmt.Scanln(&inputNumber)

	for {
		if inputNumber >= 0 && inputNumber <= 3 {
			switch inputNumber {
			case 1:
				showAllStudent()
				text()
				fmt.Scanln(&inputNumber)
				continue
			case 2:
				addStudent()
				text()
				fmt.Scanln(&inputNumber)
				continue
			case 3:
				deleteStudent()
				text()
				fmt.Scanln(&inputNumber)
				continue
			case 0:
				fmt.Println("退出系统")
				break
			}
		}
		fmt.Println("Error")
		break
	}
}
