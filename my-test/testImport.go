package my_test

import (
	"base/constant"
	"base/tool"
	"context"
	"fmt"
	"framework/conf"
	"orm"
	"orm/model"

	"gorm.io/gorm"
)

func collTable() {
	// ====== DEMO 开始 ======
	var tbl = tool.NewTable[string, string, int]() // 1. 写入
	tbl.Put("alice", "projectA", 40)
	tbl.Put("alice", "projectB", 20)
	tbl.Put("bob", "projectA", 30)
	tbl.Put("charlie", "projectC", 50)

	// 2. 查询
	var v, _ = tbl.Get("alice", "projectA")
	fmt.Println("alice/projectA =", v)

	// 3. 整行/整列
	fmt.Println("alice row:", tbl.Row("alice"))
	fmt.Println("projectA col:", tbl.Column("projectA"))

	// 4. 遍历 cellSet
	for _, c := range tbl.CellSet() {
		fmt.Printf("%s -> %s : %d\n", c.Row, c.Col, c.Val)
	}

	// 5. 删除
	tbl.Remove("alice", "projectB")
	fmt.Println("after remove alice/projectB:", tbl.Row("alice"))
}

func myPrint() {
	fmt.Printf("Hello World %s 名称", constant.AppName)
	conf.InitConfig()
	fmt.Printf("配置端口 %d", conf.Config.Http.Port)
}

func myOrm() {
	orm.Conn()
	user := model.User{Name: "Jinzhu", Age: 18}
	// Create a single record
	ctx := context.Background()

	err := gorm.G[model.User](orm.MyDb).Create(ctx, &user) // pass pointer of data to Create
	if err != nil {
		panic(err)
	}

	//// Create with result
	//result := gorm.WithResult()
	//err := gorm.G[User](db, result).Create(ctx, &user)
	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count
}
