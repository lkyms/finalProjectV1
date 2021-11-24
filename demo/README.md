# 接口文档

- [1.注册](#1注册)
- [2.注册发送验证码](#2注册发送验证码)
- [3.登录](#3登录)
- [4.找回密码](#4找回密码)
***
#### 1.注册
###### 接口功能
> 用户注册

###### URL
> http://localhost:114/api/register

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
> | 参数     | 必选  | 类型   | 说明             |
> | :------- | :---- | :----- | ---------------- |
> | nickname | false | string | 昵称     |
> | realname | false | string | 真实姓名         |
> | username | true  | string | 工号/学号/校友id |
> | password | true  | string | 密码             |
> | cap      | true  | string | 验证码           |

###### 返回字段
> | 返回字段 | 字段类型 | 说明                             |
> | :------- | :------- | :------------------------------- |
> | status   | int      | 返回结果状态。0：正常；1：错误。 |
> | msg      | string   | 返回消息                         |

###### 接口示例
> 地址：[http://localhost:114/api/register](http://localhost:114/api/register)
``` json
{
     "status": 0,
     "msg" : "注册成功"
}
```
***
#### 2.注册发送验证码

###### 接口功能
> 注册发送手机验证码

###### URL
> http://localhost:114/api/register/sendMessage

###### 支持格式
> application/json

###### HTTP请求方式
> POST

###### 请求参数
> | 参数     | 必选  | 类型   | 说明             |
> | :------- | :---- | :----- | ---------------- |
> | phone | true | string | 手机号码     |
> | type | false  | int | 类型,0为注册验证码,1为找回密码验证码，默认为0         |


###### 返回字段
> | 返回字段 | 字段类型 | 说明                             |
> | :------- | :------- | :------------------------------- |
> | status   | int      | 返回结果状态。0：正常；1：错误。 |
> | msg      | string   | 返回消息                         |

###### 接口示例
> 地址：[http://localhost:114/api/register/sendMessage](http://localhost:114/api/register/sendMessage)
``` json
{
    "status": 0,
    "msg" : "发送成功"
}
```
***
#### 3.登录

###### 接口功能
> 用户登录

###### URL
> http://localhost:114/api/login

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
> | 参数     | 必选 | 类型   | 说明             |
> | :------- | :--- | :----- | ---------------- |
> | username | true | string | 工号/学号/校友id |
> | password | true | string | 密码             |

###### 返回字段
> | 返回字段 | 字段类型 | 说明                             |
> | :------- | :------- | :------------------------------- |
> | status   | int      | 返回结果状态。0：正常；1：错误。 |
> | msg      | string   | 返回消息                         |

###### 接口示例
> 地址：[http://localhost:114/api/login](http://localhost:114/api/login)
``` json
{
     "status": 0,
     "msg" : "登录成功"
}
```
***
#### 4.找回密码

###### 接口功能
> 找回密码

###### URL
> http://localhost:114/api/findpassword

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
> | 参数     | 必选  | 类型   | 说明             |
> | :------- | :---- | :----- | ---------------- |
> | username | true  | string | 工号/学号/校友id |
> | password | true  | string | 密码             |
> | phone | true  | string | 手机号码             |
> | cap      | true  | string | 验证码           |

###### 返回字段
> | 返回字段 | 字段类型 | 说明                             |
> | :------- | :------- | :------------------------------- |
> | status   | int      | 返回结果状态。0：正常；1：错误。 |
> | msg      | string   | 返回消息                         |

###### 接口示例
> 地址：[http://localhost:114/api/findpassword](http://localhost:114/api/findpassword)
``` json
{
     "status": 0,
     "msg" : "重置成功"
}
```