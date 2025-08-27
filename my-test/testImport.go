package my_test

import (
	"base/constant"
	"context"
	"fmt"
	"framework/conf"
	"orm"
	"orm/model"

	"gorm.io/gorm"
)

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
