package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"internetbar_echo/handle"
	"internetbar_echo/service/impl"
)

func main() {
	dsn := "root:124567@tcp(localhost:3306)/blog_echo?loc=Local&parseTime=true&charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping() //验证与数据库的连接是否仍然存在
	if err != nil {
		panic(err)
	}
	svr := impl.NewArticleService(db)
	tagSvr := impl.NewTagService(db)
	admin := impl.NewAdminService(db)
	comment := impl.NewCommentService(db)
	r := handle.New(svr, tagSvr, admin, comment)
	r.Logger.Fatal(r.Start(":1323"))
	//r.Start(":1323")

	/*e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello word")
	})
	e.Logger.Fatal(e.Start(":1323")) //启动一个服务*/
}

//Socket hang up  链接被挂断
