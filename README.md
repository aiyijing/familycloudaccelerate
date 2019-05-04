# 天翼家庭云APP 破解提速脚本

## change log
二进制程序已经编译并发布

## 说明

* 思路参考:[Ruter's Journal](http://blog.ruterfu.com/2019/02/09/faster-upload-using-tianyicloud/)
破解思路,以及签名算法算法提取  
* 原理: 拦截session_key,session_secret,通过HTTP协议定时发送心跳包维持加速  

目前程序有两大分支

* shell 感谢 [vcheckzen](https://github.com/vcheckzen/FamilyCloudSpeederInShell.git) 贡献代码
* python (支持python2,python3)

## 使用方法

### 抓包  

1. 请确认当前是否支持天翼家庭云APP 天翼网盘提速,否则无法进行下一步骤  
2. 提取session_key与session_secret  
被抓包客户端必须处于光猫下局域网内。抓包方式大同小异,根本原理就是让客户端信任CA证书,进行中间人劫持攻击。我们的目的是为了提取session_key与session_secret  

* 电脑端 Charless抓包 [Ruter's Journal](http://blog.ruterfu.com/2019/02/09/faster-upload-using-tianyicloud/)

* Android HttpCanary抓包 [wiki](https://github.com/aiyijing/familycloudaccelerate/wiki/%E5%AE%B6%E5%BA%AD%E4%BA%91%E6%89%8B%E6%9C%BA%E7%AB%AF%E6%8A%93%E5%8C%85%E6%96%B9%E6%B3%95)  建议采用Android端抓包或者Android模拟器

* 更多方式请自由发挥: 安卓模拟器+Charless,安卓模拟器+HttpCanary ...

### python 版本使用
使用之前: 请先确认python版本,python-pip 是否安装,然后下载相应python 脚本

* 安装依赖  

```shell
pip install requests
```

* 配置config  

将 session_key,session_secret 写入文件

```python
{
    "session_key":"session_key",    # 必填 session_key
    "session_secret":"session_secret",# 必填 session_secret
    "setting":{
        "method":"POST",        # 可选:POST|GET
        "rate":600              # 心跳包频率 单位秒 建议修改为600
    },
    "send_data":{
          "prodCode": "76",     # 默认
          "version": "2.0.10",  # app 版本
          "channelId": "web"    # 默认参数 与用户登录方式有关
    },
    "extra_header":{
        "User-Agent": "Apache-HttpClient/UNAVAILABLE (java 1.4)"    #附加HTTP Header
    }
}
```  

* 启动程序

```shell
# 前台执行
python FamilySpeedUp.py
```

```shell
# 后台执行
nohup python FamilySpeedUp.py
```
### 二进制程序使用
* 下载相应平台程序,请移步release
* 配置config.json参数:session_key session_secret 
```shell
chmod a+x FamilySpeedUp
#config.json与程序在同一路径下
./FamilySpeedUp
#config.json与程序不在同一路径下,请提交config.json路径
./FamilySpeedup ${dir}/config.json
```

### Shell 版本使用

使用之前请确认: unixlike 环境已经存在curl,openssl
使用方法请参考: [vcheckzen](https://github.com/vcheckzen/FamilyCloudSpeederInShell.git)

## TODO

Progress: shell版本依赖:openssl curl.目前Go语言版本已经完成, 当前没有测试环境,测试完备后发出.

* 能正常提速,但是无法获取提速结果,需要修改相关接口  
* Python版本程序不便于移植嵌入式平台如:openwrt,正在编写GO语言版本以便于移植

欢迎大家提 ISSUE 本人定当竭力相助