# 问题回顾

```
2023年3月18日15:13:44

    gorm hook函数的方法接收者 必须是以 *自定义类型的方式去使用，否则无法生效
        内在原因是结构体是值类型，使用指针传递才能改变到原来数据中的值
    
    TagId 这种其他表的主键，在SQL的后续优化中需要通过 `gorm:"index"`去添加索引
    
    内嵌的结构体，
        使用 db.Model(&Article).Related(&Tag)去进行关联查询
        Prel
    
2023年3月19日19:09:37
在windows下打包linux可执行二进制文件
注意：需要使用自带的cmd
set CGO_ENABLED=0  
set GOOS=linux
set GOARCH=amd64
go build main.go
```

first、last、take、find区别

first会进行排序，主键升序；检索单个对象，LIMIT 1;not found会报错
last会进行排序，主键降序；检索单个对象，LIMIT 1;not found会报错
take不会进行排序；检索单个对象，LIMIT 1;not found会报错
find 获取所有对象；not found 不会报错
