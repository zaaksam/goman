const message: i18nModel = {
    goman: {
        general: {
            noData: '暂无数据',
            noFilterData: '暂无筛选结果'
        },
        list: {
            key: '键',
            value: '值',
            description: '备注'
        },
        app: {
            modal: {
                title: '操作确认',
                content: 'goman已退出，是否关闭本页面？',
                okText: '确定',
                cancelText: '取消'
            }
        },
        req: {
            tab: {
                title: '新标签',
                params: '地址栏参数',
                send: '发送',
                urlTips: '请输入请求地址...',
                advanced: {
                    title: '高级请求内容',
                    requests: '请求次数',
                    concurrency: '并发数量',
                    timeout: '请求超时',
                    timeoutUnit: '秒'
                },
                request: {
                    title: '请求内容',
                    code: '请求原文'
                },
                showPreview: '显示预览',
                modal: {
                    title: '错误请求',
                    content: 'Docker模式下，不支持 127.0.0.1 或 localhost 地址的请求',
                    okText: '确定'
                }
            },
            report: {
                title: '响应统计',
                summary: {
                    total: '总耗时',
                    slowest: '最慢响应',
                    fastest: '最快响应',
                    average: '平均响应',
                    rps: '每秒请求数'
                },
                table: {
                    type: '明细',
                    average: '平均',
                    slowest: '最慢',
                    fastest: '最快',
                    dns: 'DNS解析',
                    conn: '建立连接',
                    req: '请求发送',
                    delay: '响应等待',
                    res: '响应读取',
                    finish: '总耗时'
                },
                chart: {
                    title: '平均请求时间',
                    serieName: '耗时'
                }
            },
            pre: {
                listTable: {
                    hasError: '是否错误',
                    contentLength: '内容长度',
                    duration: '耗时'
                },
                title: '响应内容',
                proto: '协议',
                status: '状态',
                code: '响应原文',
                report: '统计报告',
                body: {
                    pretty: '格式化',
                    raw: '原文',
                    language: '语言'
                },
                cookie: {
                    name: '键',
                    value: '值',
                    path: '路径',
                    domain: '作用域',
                    expires: '有效期'
                }
            }
        }
    }
}

export default message