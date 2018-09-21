package search

import (
	"encoding/json"
	"os"
)

const dataFile = "github/golearning/chapter2/data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"` //对应json中的site字段，其它同理
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile) //打开json文件
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()  //等待文件中的内容都读到内存中然后就关闭文件引用

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed  //Feed指针数组的切片，切片和数组最大的不同是数组是固定的，但是切片可以添加、删除等操作
	err = json.NewDecoder(file).Decode(&feeds) //解析json并填入feeds切片中

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
