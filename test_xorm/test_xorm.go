package main

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"log"
)

// 1.获取一个表中最新的id
// 	table := "acl_order"
//	has, _ := Orm.Limit(1).Table(&table).Desc("id").Get(&lastId)

// 2.gin的ctx.JSON()是直接返回还是执行到函数结束

// 3, 按id更新一个interface
//func Update(id int64, beanPtr interface{}) (int64, error) {
//		return Orm.ID(id).Update(beanPtr)

// 4.按id获取一个数据
//	table := "acl_order"
//		has, _ := Orm.Table(&table).Where("id=?", id).Get(&res)

type Article struct {
	Id 	int64
	Title	string
	Words	string
}

// 如果结构体有Tablename方法，则orm使用时默认使用return名的数据库表，否则使用.Table()指定的
func(this *Article) TableName() string {
	return "article"
}

var aRepo *Article = new(Article)

func (this *Article) Insert() (int64, error) {
	return Insert(this)
}




var orm *xorm.Engine

func Insert(beans *Article) (int64, error) {
	return orm.Insert(&beans)
}

 
//创建orm引擎
func init() {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:07~@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		fmt.Println("数据库连接失败:", err)
	}
	fmt.Println("[Connecting MySQL Success!]")

}

func main() {
	//title := "十年之间"
	//res := make([]*Article, 0)
	//var res []string
	//err := orm.Table("article").Cols("title").Find(&res)
	res, err := orm.Exist(&Article{Title:"十年之间1"})
	//err := orm.Find(&res)
	//title := "zhaoyi"
	//words := "8808测试数据"
	//edit := &Article{Words:"5555"}
	//res, err := orm.Where("title=? and words=?", title, words).Update(edit)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res", res)
//	for _, v := range res {
//		//fmt.Println(v.Title)
//		fmt.Println(v)
//
//	}

	//var article Article
	//table := "article"
	//article.Title = "测试数据，敬请期待，10.66.66.66"
	//article.Words = "测试数据，敬请期待，10.777.777.777"  // 这种格式，居然返回不到id

	//article := &Article{Title:"测试数据", Words:"测试数据"} // 这种格式，新增数据后，可以返回article.Id
	//res, err := Insert(article)

	//res, err := orm.Update(&Article{Title:"测试数据，敬请期待，10.66.66.66"}) // test ok，目前4条数据，res返回了4，全部数据修改
	//res, err := orm.ID(22).Update(&Article{Title:"测试数据，敬请期待，10.1.1.1"}) // test ok，一开始int类型不一致导致出错

	//res, err := orm.Insert(article) // test ok ，but 将数据类型修改后，

	//var lastid int64
	// 按id获取一个数据
	//var article []map[string]string
	//id := 21
	//article := new(Article)
	//res, err := orm.Table("article").Where("id=?", id).Get(&article) 

	// 获取符合某列条件
	//title := "zhaoyi"
	//res, err := orm.Table("article").Where("title=?", title).QueryString() // test ok

	// 模糊查询
	//title := "zhao"
	//res, err := orm.Table("article").QueryString("select * from article where title like ?", "%"+title+"%") // test ok
	//res, err := orm.Table("article").Where("title like ?", "%"+title+"%").QueryInterface() // test ok,输出的string是字节


	// 获取最新id
	//res, err := orm.Limit(1).Table(table).Desc("id").Cols("id").Get(&lastid) // test ok
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(lastid, "res", res)
	//fmt.Println("res", res, article.Title)
	//fmt.Println(article)
}
