# goman

a web app like postman

# 已编译平台(go1.9.2编译)

[goman.v0.2.0.app-darwin64.zip](https://github.com/zaaksam/goman/files/1685050/goman.v0.2.0.app-darwin64.zip)

[goman.v0.2.0.web-darwin64.zip (需要在终端下运行)](https://github.com/zaaksam/goman/files/1685047/goman.v0.2.0.web-darwin64.zip)

[goman.v0.2.0.web-win64.zip](https://github.com/zaaksam/goman/files/1685038/goman.v0.2.0.web-win64.zip)

windows下app版本因兼容问题未发版

# Docker镜像

`docker pull zaaksam/goman`

`docker run -p 8080:8080 zaaksam/goman`

# 界面预览

![](https://static.oschina.net/uploads/space/2018/0203/091232_EVdr_2686168.jpg)

![](https://static.oschina.net/uploads/space/2018/0203/091306_FECg_2686168.jpg)

# 技术资源

Backend Go + beego

Frontend Typescript + Vue + iView

# 更新日志

2018-02-03 v0.2.2

* 增加docker镜像编译脚本
* 修改代码支持跨平台条件编译
* 现在访问首页会默认跳往webui地址

2018-02-02 v0.2.1

* 优化构建脚本，更好支持不同平台交叉编译
* 增加不同版本构建main入口

2018-02-01 v0.2.0

* 增加多语方支持，目前支持：简体中文（默认）、英文（不完全）
* 恢复app模式（webview），并默认编译
* 请求支持高级选项模式，可批量请求，并提供统计报告

2017-09-12 v0.1.3

* 拆分webpack为dev、prod环境
* 升级相关依赖项

---

2017-09-10 v0.1.2

* 项目初始化