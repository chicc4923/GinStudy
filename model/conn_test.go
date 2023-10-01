package model

import (
	"log"
	"testing"
)

// TestSetupDB 测试初始化数据库
func TestSetupDB(t *testing.T) {
	err := SetupDB()
	if err != nil {
		log.Fatal(err)
	}
}
