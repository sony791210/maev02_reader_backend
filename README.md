## 先決條件，共同說明  

1. 前端 :react   閱讀器  
2. 後端 :go      API 開發  
3. 資料庫 :mysql   
4. 資料來源 :python  爬蟲  



# maev02_reader_api
使用go filber 開發  
資料夾結構參考laravel 實作  


## 使用 
cp .env.example .env  
修改相關內容  

go mod vender  

go run main.go  

會產生相關 swager  文件   
[http://127.0.0.1:2000/swagger/index.html](http://127.0.0.1:2000/swagger/index.html)

使用swag前可以除錯 
swag init 


