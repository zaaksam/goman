# goman

a web app like postman

# 已编译平台(go1.9.x编译)

web版使用了新的运行形式，使用更便捷，推荐使用web版

[goman.v0.3.0.web-win.zip](https://github.com/zaaksam/goman/releases/download/v0.3.0/goman.v0.3.0.web-win.zip)

[goman.v0.3.0.web-mac.tar.gz](https://github.com/zaaksam/goman/releases/download/v0.3.0/goman.v0.3.0.web-mac.tar.gz)

[goman.v0.3.0.app-mac.tar.gz](https://github.com/zaaksam/goman/releases/download/v0.3.0/goman.v0.3.0.app-mac.tar.gz)

windows下app版本因IE内核兼容问题未发版

# Docker镜像

`docker pull zaaksam/goman` 或者指定版本 `docker pull zaaksam/goman:0.3.0`

`docker run -p 8080:8080 zaaksam/goman`

在浏览中打开：`http:127.0.0.1:8080`

注意：docker模式下，goman处在容器内，无法使用 localhost(127.0.0.1) 来请求宿主机的网络资源

# 界面预览

windows web版

![](https://static.oschina.net/uploads/img/201802/08120715_zvnn.jpg)

macos web版

![](https://static.oschina.net/uploads/img/201802/08120750_hnI4.jpg)

macos app版（实验）

![](https://static.oschina.net/uploads/img/201802/08120826_tMsb.jpg)

![](https://static.oschina.net/uploads/img/201802/08120851_rVD1.jpg)

# 技术资源

Backend Go + beego

Frontend Typescript + Vue + iView

web版引导界面使用了 knockout + layui样式

# 更新日志

2018-02-08 v0.3.0

* 更便捷的web模式运行方式
* web模式下增加了浏览器与应用的互相监测
* docker模式下请求localhost(127.0.01)时会收到警告提示
* app模式建议为实验性质使用
* 简化及增强了各个编译脚本内容

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