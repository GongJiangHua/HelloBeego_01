package db_mysql

import (
	"BeegoProject0603/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	//mysql 驱动
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
/**
 * 连接mysql数据库
 */
func Connect(){

	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword)

	//连接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {// err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早解决，早解决
		panic("数据库连接错误，请检查配置")
	}
	Db = db
	fmt.Println(db)
}

/**
* 将用户信息保存到数据库中去的函数
 */
func AddUser(u models.User)(int64, error){
	//1、将密码进行hash计算，得到密码hash值，然后在存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	psswordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(psswordBytes)
	//execute， .exe
	result, err :=Db.Exec("insert into person(name,birthday,address,password)" +
		" values(?,?,?,?) ", u.Name,u.Birthday,u.Address,u.Password)
	if err != nil {
		return -1,err
	}
	row,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return row,nil
}
func QueryUser(u models.User)(int64,error) {
	rows,err :=Db.Query("select * from person where name = ?",u.Name )

	 //fmt.Println(err)
	 //fmt.Println(rows)

	for rows.Next(){
		var userid string
		var username string
		var birthday string
		var address string
		var password string
		rows.Columns()
		u.Name=userid
		u.Birthday=birthday
		u.Address=address
		u.Password=password
		fmt.Println(u.Name)
		fmt.Println(u.Password)
		fmt.Println(u.Address)
		fmt.Println(u.Password)
		err = rows.Scan(&userid,&username,&birthday,&address,&password)
		fmt.Println(userid,username,birthday)
	}


	if err!=nil {
		return -1,err
	}else {
		return 1,nil
	}

}
