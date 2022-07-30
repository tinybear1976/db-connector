# 配置说明
## 优先级
如果存在连接器文件，比如 xxx.ctor,那么模块会先行解析并加载。如果配置文件中存在配置信息部分(明文结构体)，模块会继续分析，并覆盖相同key的连接信息。如果是在开发模式下，方便明文配置数据库连接，则可以在配置文件中保留配置信息部分，作为正式发行版本，应该删除该配置数据，只保留连接器文件即可