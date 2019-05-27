# QLog

#### 功能
- [x] 设置日志输出级别debug/info/warn/error/fatal
- [x] 设置输出位置：文件/控制台
- [x] 打出日志相应文件及所在行数
- [x] 自定义日期格式（DateFormat）
- [x] 自定义单个文件大小(FileMaxSize，单位B，默认1G)，当日文件info_2018-06-30.log，大于设定文件大小后，info_2018-06-30.1.log，序号依次增加
- [x] 支持双写（终端与文件）
- [x] 打出当前机器IP
- [x] 设置日志打印最大个数
- [ ] 设置日志是否gip打包
- [ ] 异步打印日志

#### 任务
- [ ]

#### 使用示例
[`参见test/logger_test.go`](test/logger_test.go)