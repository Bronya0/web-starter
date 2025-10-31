# web-starter

#### 介绍
脚手架，集成常用web开发工具

注意，该仓库有多个branch，不同分支实现不同，比如：
- simple分支是极简设计，仅保留必要功能（如db/log），适合简单API
- 默认分支集成度较高



- http框架：echo（不太喜欢gin）
- 主协程崩溃恢复
- RequestID
- CSRF防御
- XSS防御
- 日志：支持压缩、按时间轮转、按体积轮转
- kafka客户端
- es客户端
- redis客户端
- 配置管理
- 支持定时任务
- gorm数据库
- jwt
- 路由区分公开api和认证api
- 邮件发送
- AES加密：CBC、GCM

认证授权需要自己根据业务完善。
