package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMysqlConnection() rel.Adapter {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FJakarta",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	adapter, err := mysql.Open(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	return adapter
}
