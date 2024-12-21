package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 定义数据库连接字符串，格式为：用户名:密码@tcp(主机地址:端口号)/数据库名
	db, err := sql.Open("mysql", "u982684694_airrock91:Justin576@tcp(154.56.47.1:3306)/test")
	if err != nil {
		log.Println("数据库连接失败:", err)
		return
	}
	// 延迟关闭数据库连接，确保资源释放
	defer db.Close()
	// 测试连接是否可用
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接:", err)
		return
	}
	log.Println("数据库连接成功")
}
