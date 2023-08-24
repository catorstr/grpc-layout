package script

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func BuildDbModel(mysqlSource string) {
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlSource)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../internal/db/mysql/gorm/dal", //注意修改目录，使其生成到对应的地方
		ModelPkgPath: "../gorm/model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	//设置目标 db
	g.UseDB(db)
	//g.ApplyBasic(g.GenerateAllTable()...) //渲染所有表
	g.ApplyBasic( //指定表与字段类型
		g.GenerateModel("issue"),
		g.GenerateModel("like"),
		g.GenerateModel("comment"),
		g.GenerateModel("cart"),
		g.GenerateModel("label"),
		g.GenerateModel("shop_addr"),
		g.GenerateModel("order_forward", gen.FieldType("price", "decimal.Decimal")),
		g.GenerateModel("user", gen.FieldType("balance", "decimal.Decimal")),
		g.GenerateModel("trace_data", gen.FieldType("price", "decimal.Decimal")))
	g.Execute()
}
