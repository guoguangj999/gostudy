package main

import (
	"encoding/json"
	"fmt"

	"com.ggj.com/ggj/app"
)

func main() {

	mod := app.App{Name: "zhangsan", Version: 2}

	fmt.Printf("mod name is %s", mod.Name)
	//json序列化
	js, _ := json.Marshal(mod)

	fmt.Printf("json string = %s", js)

}
