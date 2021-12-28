//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Mgo struct {
	Session *mgo.Session
	DB      *mgo.Database
	MgoHost string
	MgoDB   string
}

type Url struct {
	Id  bson.ObjectId `bson:"_id"`
	Url string        `json:"url"`
}

type Item struct {
	Id       bson.ObjectId `bson:"_id"`
	Votes    int           `bson:"votes"`
	Answers  int           `bson:"answers"`
	Views    int           `bson:"views"`
	Url      string        `bson:"url"`
	Question string        `bson:"question"`
	//define by yourself...
}

func InitMgoDB(ConnStr, DBName string) *Mgo {
	mgoClient := &Mgo{
		MgoHost: ConnStr,
		MgoDB:   DBName,
	}
	session, err := mgo.Dial(ConnStr)
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InitMgoDB error: " + err.Error())
	}
	//defer session.Close()

	//设置模式
	session.SetMode(mgo.Monotonic, true)
	//获取文档集
	c := session.DB(DBName).C("urls")
	// 创建索引
	index := mgo.Index{
		Key:        []string{"url"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     true,            // 唯一索引 同mysql唯一索引
		DropDups:   true,            // 索引重复替换旧文档,Unique为true时失效
		Background: true,            // 后台创建索引
	}
	if err := c.EnsureIndex(index); err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InitMgoDB error: " + err.Error())
		return nil
	}

	c2 := session.DB(DBName).C("questions")
	// 创建索引
	index2 := mgo.Index{
		Key:        []string{"question"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     true,                 // 唯一索引 同mysql唯一索引
		DropDups:   true,                 // 索引重复替换旧文档,Unique为true时失效
		Background: true,                 // 后台创建索引
	}
	if err := c2.EnsureIndex(index2); err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InitMgoDB error: " + err.Error())
		return nil
	}

	mgoClient.DB = session.DB(DBName)
	mgoClient.Session = session
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InitMgoDB success")
	return mgoClient
}

func (mgo *Mgo) Close() {
	mgo.Session.Close()
}

func (mgo *Mgo) InsertUrls(urls []string) (err error) {
	c := mgo.DB.C("urls")
	for _, url := range urls {
		tmp := &Url{
			Id:  bson.NewObjectId(),
			Url: url,
		}
		err = c.Insert(tmp)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InsertUrls error: " + err.Error())
			break
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InsertUrls success url: " + url)
	}
	return err
}

func (mgo *Mgo) InsertItems(items []Item) (err error) {
	c := mgo.DB.C("questions")
	for _, item := range items {
		tmp := &Item{
			Id:       bson.NewObjectId(),
			Votes:    item.Votes,
			Answers:  item.Answers,
			Views:    item.Views,
			Url:      item.Url,
			Question: item.Question,
		}
		err = c.Insert(tmp)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InsertItems error: " + err.Error())
			break
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go InsertItems success")

	return err
}

func (mgo *Mgo) QueryUrls(topN int) ([]Url, error) {
	//*****查询多条数据*******
	var urls []Url //存放结果
	c := mgo.DB.C("urls")
	iter := c.Find(nil).Limit(topN).Iter()
	err := iter.All(&urls)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go QueryUrls ")
	return urls, err
}

func (mgo *Mgo) DeleteUrl(url Url) error {
	c := mgo.DB.C("urls")
	c.Remove(bson.M{"url": url.Url})
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " mongodb.go DeleteUrl url: " + url.Url)
	return nil
}
