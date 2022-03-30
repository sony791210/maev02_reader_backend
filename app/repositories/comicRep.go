package repositories

import (
	"fmt"
	"io/ioutil"
)

func GetComicPageList(comicname string, page string) interface{} {
	var filePath []string
	files, err := ioutil.ReadDir("./public/" + comicname + "/" + page)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		filePath = append(filePath, "/static/"+comicname+"/"+page+"/"+file.Name())
	}

	return filePath
}
