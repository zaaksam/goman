const message: i18nModel = {
    goman: {
        general: {
            noData: 'no data',
            noFilterData: 'no data'
        },
        list: {
            key: 'Key',
            value: 'Value',
            description: 'Description'
        },
        app: {
            modal: {
                title: 'Warning',
                content: 'Goman has closed! Do you want to leave this page?',
                okText: 'Yes',
                cancelText: 'No'
            }
        },
        req: {
            tab: {
                title: 'New Tab',
                params: 'Params',
                send: 'SEND',
                urlTips: 'Enter request URL',
                advanced: {
                    title: 'Advanced Request',
                    requests: 'Requests',
                    concurrency: 'Concurrency',
                    timeout: 'Timeout',
                    timeoutUnit: 'Seconds'
                },
                request: {
                    title: 'Request',
                    code: 'Code'
                },
                showPreview: 'Show Preview',
                modal: {
                    title: 'Warning',
                    content: "You can't access localhost(127.0.0.1) network in a docker container",
                    okText: 'OK'
                }
            },
            report: {
                title: 'Response Report',
                summary: {
                    total: 'Total',
                    slowest: 'Slowest',
                    fastest: 'Fastest',
                    average: 'Average',
                    rps: 'Requests/sec'
                },
                table: {
                    type: 'Detail',
                    average: 'Average',
                    slowest: 'Slowest',
                    fastest: 'Fastest',
                    dns: 'DNS',
                    conn: 'Conn',
                    req: 'Req Write',
                    delay: 'Res Wait',
                    res: 'Res Read',
                    finish: 'Total'
                },
                chart: {
                    title: 'Average Duration',
                    serieName: 'Duration'
                }
            },
            pre: {
                listTable: {
                    hasError: 'Error',
                    contentLength: 'Content Length',
                    duration: 'Duration'
                },
                title: 'Response',
                proto: 'Proto',
                status: 'Status',
                code: 'Code',
                report: 'Report',
                body: {
                    pretty: 'Pretty',
                    raw: 'Raw',
                    language: 'Language'
                },
                cookie: {
                    name: 'Name',
                    value: 'Value',
                    path: 'Path',
                    domain: 'Domain',
                    expires: 'Expires'
                }
            }
        }
    }
}

export default message