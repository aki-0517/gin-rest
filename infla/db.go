package infla

import (
	"fmt"
    "log"
    "os"
    "time"

    gormmysql "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"

)

func DSN() string {
	return fmt.Sprintf(
	  "%s:%s@tcp(%s:%s)/%s",
	  os.Getenv("DB_USER"),
	  os.Getenv("DB_PASSWORD"),
	  os.Getenv("DB_HOST"),
	  os.Getenv("DB_PORT"),
	  os.Getenv("DB_NAME"),
	) + "?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True&loc=Asia%2FTokyo&multiStatements=true"
  }


func NewDB(count uint) (*gorm.DB, error) {
    dsn := DSN()
    return dbConnect(gormmysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    }, count)
}

func dbConnect(dialector gorm.Dialector, config *gorm.Config, count uint) (*gorm.DB, error) {
    var db *gorm.DB
    var err error

    // countで指定した回数リトライする
    for ; count > 0; count-- {
        db, err = gorm.Open(dialector, config)
        if err != nil {
            if count == 1 { // Last attempt
                log.Println(err.Error())
                return nil, err
            }
            time.Sleep(time.Second * 2)
            log.Printf("retry... count:%v\n", count-1)
        } else {
            break
        }
    }
    return db, err
}


