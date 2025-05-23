package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// os.Args demo
func main() {
	myFlag := "flag"
	switch myFlag {
	case "os.Args":
		if len(os.Args) > 0 {
			for index, arg := range os.Args {
				fmt.Printf("args[%d] = %v\n", index, arg)
			}
		}
	case "flag":
		// falg名 默认值 帮助信息
		var name string
		var age int
		var married bool
		var delay time.Duration
		flag.StringVar(&name, "name", "张三", "姓名")
		flag.IntVar(&age, "age", 18, "年龄")
		flag.BoolVar(&married, "married", false, "婚否")
		flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

		// 解析命令行参数
		flag.Parse()
		fmt.Println(name, age, married, delay)
		// 返回命令行参数后的其他参数
		fmt.Println("命令行参数后的其他额外参数", flag.Args())
		// 返回命令行参数后的其他参数个数
		fmt.Println("命令行参数后的其他额外参数个数", flag.NArg())
		// 返回使用的命令行参数个数
		fmt.Println("使用的命令行参数个数", flag.NFlag())
	}
}
