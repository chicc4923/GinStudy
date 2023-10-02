package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonDecode(t *testing.T) {
	str := `{"age":18}`                   // 模拟前端传来的某条参数；
	person := Person{Name: "老六", Age: 27} // 模拟数据库中已有的记录；
	fmt.Println(json.NewDecoder(strings.NewReader(str)))

	//if err == nil {
	fmt.Printf("解码结果: %+v", person)
	//} // 解码结果: {Name:老六 Age:18}
}
