package main

import (
	"os"
	"io"
	"log"
	"strconv"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher 是协程安全的
	searcher = riot.Engine{}
	db *sql.DB
)

func init_db() bool {
	var err error
	db, err = sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/CanEat?charset=utf8")
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func init_searcher() bool {
	rows, err := db.Query("SELECT id,name FROM foods")
	if err != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()

	// 初始化
	searcher.Init(types.EngineOpts{
		Using: 3,
		GseDict: "zh",
	})

	// 将文档加入索引，docId 从1开始
	for rows.Next() {
		var id uint64
		var name string
		rows.Scan(&id, &name)
		searcher.Index(id, types.DocData{Content: name})
	}

	// 等待索引刷新完毕
	searcher.Flush()
	return true
}

func main() {

	if init_db() {
		defer db.Close()
	} else {
		return
	}

	if init_searcher() {
		defer searcher.Close()
	} else {
		return
	}

	gin.DisableConsoleColor()
	f, _ := os.Create("log/run.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.GET("/foods/info_list", foods_info_list)
	r.Run() // default on 0.0.0.0:8080
}

func foods_info_list(c *gin.Context) {
	key := c.Query("key")

	// 搜索输出格式见 types.SearchResp 结构体
	resp := searcher.SearchDoc(types.SearchReq{Text: key})

	tableData := make([]map[string]interface{}, 0)
	for _, doc := range resp.Docs {
		rows, err := db.Query("SELECT * FROM foods WHERE id=" + strconv.FormatUint(doc.DocId, 10))
		if err != nil {
			log.Println(err)
			continue
		}
		defer rows.Close()
		
		columns, err := rows.Columns()
		if err != nil {
			log.Println(err)
			continue
		}
		count := len(columns)

		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for rows.Next() {
			for i := 0; i < count; i++ {
				valuePtrs[i] = &values[i]
			}

			rows.Scan(valuePtrs...)
			entry := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
			}
			tableData = append(tableData, entry)
		}
	}

	c.JSON(200, gin.H{
		"num": len(tableData),
		"data": tableData,
	})
}
