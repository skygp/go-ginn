## Go-Ginn简介:
**开箱即用的Golang脚手架，简称:ginn**

**Gin框架 + 彪悍库**

#### 现有库:
- gorm    
- viper   
- zap        
- casbin  
- 七牛云   
- jwt库   
- swagger   
- 雪花算法生成UUID
- bcrypt加密
- air调试神器
- 优雅重启
- 优雅关机
- utils公共库

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
- Ginn-QQ群: 458314077
- 作者QQ:    1209371783
- 加群与加好友备注: ginn，谢谢！

### 支持开源

<img src="temp/zfb.jpeg" width="20%">    <img src="temp/wx.jpeg" width="20%">
