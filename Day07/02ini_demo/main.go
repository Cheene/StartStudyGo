package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MySqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MySqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

//加载一个配置文件
func loadInit(fileName string, data interface{}) (err error) {
	//0 参数的校验
	//0.1 data 参数必须是指针类型，因为再函数中对其赋值
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data parm is not ptr type: %v", t.Kind())
		return
	}
	//0.2 传递的指针必须是一个结构体类型指针，是因为配置文件中各种键值对需要赋值给键值对
	// t.Elem() 获取指针的值
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data is not a struct,err: %v", t.Elem().Kind())
		return
	}
	//1 读文件得到字节数据
	//fileObj,err := os.Open(fileName)
	//if err != nil {
	//	fmt.Printf("The file is not exist, err: %v",err)
	//	return
	//}
	//var []byteData = byte(fileObj)
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("The file is not exist, err: %v", err)
		return
	}
	lineSilce := strings.Split(string(b), "\r\n")
	// %#v 让slice输出的格式更形象一下
	fmt.Printf("%#v\n", lineSilce)
	var structName string
	//2 一行一行的数据
	for idx, line := range lineSilce {
		line = strings.TrimSpace(line) //去掉首尾空格
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就忽略
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[ 开头，代表是新的节(section) ==> 创建相关的结构体
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error\n", idx+1)
				return
			}
			if len(strings.TrimSpace(line)) == 1 {
				err = fmt.Errorf("line: %d syntax error\n", idx+1)
				return
			}
			//获取节的名字
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error\n", idx+1)
				return
			}
			//根据字符串去寻找对应的结构体
			//1根据结构体获取所有的变量名
			for i := 0; i < t.Elem().NumField(); i++ {
				filed := t.Elem().Field(i)
				if sectionName == filed.Tag.Get("ini") {
					//获取字段名成功，需要记录下来
					structName = filed.Name
					fmt.Printf("找到%s 对应的结构体%s\n", sectionName, structName)
					break
				}
			}
		} else {
			/*
				// 2.3 否则就是 以 = 分割的键值对，并赋值给相关的结构体
				// 2.3.1 去掉首尾空格并判断是否有 =
				index := strings.Index(strings.TrimSpace(line),"=")
				if index == -1 || strings.HasPrefix(strings.TrimSpace(line),"="){
					err = fmt.Errorf("line: %d syntax error\n",idx+1)
					return
				}
				// 2.3.2 根据 = ,分割成两个部分
				key := strings.TrimSpace(line[:index])
				value := strings.TrimSpace(line[index+1:])

				// 2.3.3 根据 structname,去 data 里面取出来
				v := reflect.ValueOf(data)
				sValue := v.Elem().FieldByName(structName)
				sType := sValue.Type()
				//判断进一步里面的字段是不是结构体
				if sValue.Kind() != reflect.Struct{
					err = fmt.Errorf("data中的%s字段应该是一个结构体",sValue.Kind())
					return
				}
				// 2.3.4 变例每一个结构体的字段，判断Tag 是否等于 key
				var fieldName string
				var filedType reflect.StructField
				for i:=0; i < sValue.NumField(); i++{
					filed := sType.Field(i)
					filedType = filed
					if filed.Tag.Get("ini") == key{
						//找到了对应的字段
						fieldName = filed.Name
						break
					}
				}
				if len(fileName) == 0{
					continue
				}
				// 2.3.5 如果 key = tag,给字段赋值
				// 根据 fieldName取出这个歌字段对其赋值
				fileObj := sValue.FieldByName(fileName)
				fmt.Println(fieldName,filedType.Type.Kind())
				fmt.Println("value: ",value)
				switch filedType.Type.Kind() {
				case reflect.String:
					fileObj.SetString(value)
				case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
					var valueInt int64
					valueInt,err := strconv.ParseInt(value,10,64)
					if err != nil {
						err = fmt.Errorf("line: %d value type errpr",idx+1)
					}
					fileObj.SetInt(valueInt)
				}
			*/
		}
	}
	return
}

func main() {
	var config Config
	err := loadInit("./config.ini", &config)
	if err != nil {
		fmt.Printf("load ini failed,err: %v\n", err)
	}
	fmt.Println(config)
}
