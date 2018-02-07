package views

const Welcome = `<!DOCTYPE html>
<html>
<head>
    <title>{{.appName}}</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link rel="stylesheet" href="/static/css/layui.css">
</head>
<body class="layui-layout-body">
    <div class="layui-layout layui-layout-admin">
        <div class="layui-header">
            <div data-bind="text: title" class="layui-logo"></div>
            <ul class="layui-nav layui-layout-right">
                <li class="layui-nav-item"><a data-bind="click: onGithub" href="javascript:;">Github</a></li>
                <li class="layui-nav-item"><a data-bind="click: onGitee" href="javascript:;">Gitee</a></li>
                <li class="layui-nav-item"><a data-bind="click: onClose" href="javascript:;">退出</a></li>
            </ul>
        </div>
        <div class="layui-body" style="left:145px;top:130px;">
            <input data-bind="css: { 'layui-btn-disabled': isConnect }, value: btnText, disable: isConnect, click: onOpen" type="button" class="layui-btn layui-btn-lg layui-btn-radius layui-btn-normal" />
        </div>
    </div>
</body>
{{.windowsDebug}}
<script src="/web/config.js{{.unix}}"></script>
<script src="/static/lib/knockout-3.4.2.min.js"></script>
<script src="/static/lib/welcome.js{{.unix}}"></script>
</html>`
