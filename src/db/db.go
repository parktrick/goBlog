package db

import (
	"engine"
	"log"
	"os"
	"path"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// 设置数据库路径
	_DB_NAME = "data/goBlog.db"
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

var (
	orm *xorm.Engine
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `xorm:"index"`
	Views           int64     `xorm:"index"`
	TopicTime       time.Time `xorm:"index"`
	TopicCount      int
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
	Content         string `xorm:"size(5000)"`
	Attachment      string
	Created         time.Time `xorm:"index"`
	Updated         time.Time `xorm:"index"`
	Views           int64     `xorm:"index"`
	Author          string
	ReplyTime       time.Time `xorm:"index"`
	ReplyCount      int
	ReplyLastUserId int64
}

// 评论
type Reply struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

// 账号
type Account struct {
	Id   int64
	Name string
	Pwd  string
}

func RegisterDB() {
	// 检查数据库文件
	if !engine.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	var err error
	orm, err = xorm.NewEngine(_SQLITE3_DRIVER, _DB_NAME)
	if err != nil {
		log.Fatalf("fail to create xorm Engine: %v", err)
	}
	err = orm.Sync(
		new(Category),
		new(Topic),
		new(Reply),
		new(Account),
	)
	//	orm.ShowSQL(true)
}

func GetOrm() *xorm.Engine {
	return orm
}
