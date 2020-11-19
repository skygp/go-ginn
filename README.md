## Go-Ginn简介:
**开箱即用的Golang脚手架，简称:ginn**

**Gin框架 + 彪悍库**

#### 现有库:
- gorm    操作数据库的库
- viper    读取配置文件库
- zap 日志库    打印日志    
- casbin 库    支持RBAC(角色的访问控制)
- 七牛云     文件上传
- jwt库      生成token与校验
- swagger   api文档生成
- 雪花算法生成UUID
- bcrypt加密
- air调试神器
- 优雅重启
- 优雅关机

### 环境搭建调试和发布:
+ **初始环境**:

  * > 设置go modules    Proxy 建议使用: https://goproxy.cn

  + > go mod tidy

* **调试与发布**:

     * debug环节，在终端输入air即可调试

        > air  或 air -c  ./.air.conf

        

     * build环节，在终端输入make即可构建go二进制文件

        > make 或 make all

* **other**: 支持直接使用swagger生成API, 需要在api接口附近添加
  
  * swag init



## 广告：

- Ginn均已成型，剩下业务部分，如果对项目有兴趣，邀请您加入开发。
- QQ交流群: 458314077
- 加群备注: ginn。 否则拒绝加入哦，谢谢关注。

********************************************

### 假期合作：

- **接受外包或产品  | QQ: 1209371783 | 加好友备注: 合作**



### 捐款 | 开源不易 - 您的支持 - 将是种动力

<img src="temp/zfb.jpeg" width="20%">    <img src="temp/wx.jpeg" width="20%">
