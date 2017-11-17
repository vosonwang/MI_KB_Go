package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

func AddItem(title string, expand bool, level int) {

	u := uuid.NewV4()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO catalog(id,title,expand,type) VALUES($1,$2,$3,$4) RETURNING id")
	checkErr(err)

	res, err := stmt.Exec(u, title, expand, level)

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlDelete() {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
	}
}
func SqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficows", 16)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
