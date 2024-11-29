# Gorm(Go语言的orm，orm（Object Relational Mapping），对象关系映射)

## 映射关系

| 数据表  | <->  | 结构体  |
|---|---|---|
|  数据行 |  <-> | 结构体实例  |
| 字段  |   <-> |  结构体字段 |

## 其它问题
1. 怎么传空串给数据库？
   1. Name *string,传递的时候使用new(string)
   2. Name sql.NullString,传递的时候使用
   
      Name: sql.NullString{
         String: "",
         Valid:  true,
      }
