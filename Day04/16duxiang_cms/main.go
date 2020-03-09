package main

import (
	"fmt"
	"os"
)

/**
通过面向对象的方式实现简单版管理系统
1 查看所有的学生
2 新增学生
3 删除学生

**/
type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

func text() {
	fmt.Println("--------------MENU---------------")
	fmt.Println(`
1 查看所有的学生
2 新增学生
3 删除学生 `)
	fmt.Println("等待输入.,按0退出系统 ......")
}
func showAllStudent(allStu []student) {
	for _, v := range allStu {
		fmt.Printf("ID:%d\t,Name:%v\n", v.id, v.name)
	}
	fmt.Println("所有的学生输出完成\n")
}
func addStudent(s []student, id int64, name string) []student {
	return append(s, newStudent(id, name))
}
func deleteStudent(s []student, id int64) []student {
	for k, v := range s {
		if v.id == id {
			fmt.Println("删除成功")
			return append(s[:k], s[k+1:]...)
		}
	}
	fmt.Println("删除失败")
	return s
}
func main() {
	allStudent := []student{}
	fmt.Println("系统初始化成功")
	text()

	for {
		var inputNumber int
		fmt.Scanln(&inputNumber)
		if inputNumber > 0 && inputNumber < 4 {
			switch inputNumber {
			case 1:
				showAllStudent(allStudent)
				text()
				continue
			case 2:
				fmt.Print("ID: ")
				var tempId int64
				var tempName string
				fmt.Scanln(&tempId)
				fmt.Print("Name: ")
				fmt.Scanln(&tempName)

				allStudent = addStudent(allStudent, tempId, tempName)
				fmt.Println("添加学生成功")
				showAllStudent(allStudent)
				text()
				continue
			case 3:
				var tempID int64
				fmt.Print("请输入ID： ")
				fmt.Scanln(&tempID)
				allStudent = deleteStudent(allStudent, tempID)
				text()
				continue
			case 0:
				fmt.Println("退出系统")
				break
			default:
				fmt.Println("Errror")
				os.Exit(1)
			}
		}
	}
}
