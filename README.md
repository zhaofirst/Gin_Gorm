# 基本介绍
网络链接使用gin框架
数据库映射使用gorm框架
实现了对数据库表格demo基于主键的增、删、改、查

# 文件系统
httpfile实现网络链接
dbfile实现数据库映射
models定义数据库表格（此处为User:   ID      int    `json:"id"`
                                Name    string `json:"name"`
                                Age     int    `json:"age"`
                                Email   string `json:"email"`
                                Phone   string `json:"phone"`
                                Address string `json:"address"`)


# test文件
测试了输入id 1 和 id 12的不同人物的增删改查询