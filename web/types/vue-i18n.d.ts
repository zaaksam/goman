/// <reference types='vue-i18n' />

declare interface i18nModel {
    goman: {
        general: {
            noData: string
            noFilterData: string
        }
        list: {
            key: string
            value: string
            description: string
        }
        app: {
            modal: {
                title: string
                content: string
                okText: string
                cancelText: string
            }
        }
        req: {
            tab: {
                title: string
                params: string
                send: string
                urlTips: string
                advanced: {
                    title: string
                    requests: string
                    concurrency: string
                    timeout: string
                    timeoutUnit: string
                }
                request: {
                    title: string
                    code: string
                }
                showPreview: string
                modal: {
                    title: string
                    content: string
                    okText: string
                }
            }
            report: {
                title: string
                summary: {
                    total: string
                    slowest: string
                    fastest: string
                    average: string
                    rps: string
                }
                table: {
                    type: string
                    average: string
                    slowest: string
                    fastest: string
                    dns: string
                    conn: string
                    req: string
                    delay: string
                    res: string
                    finish: string
                }
                chart: {
                    title: string
                    serieName: string
                }
            }
            pre: {
                listTable: {
                    hasError: string
                    contentLength: string
                    duration: string
                }
                title: string
                proto: string
                status: string
                code: string
                report: string
                body: {
                    pretty: string
                    raw: string
                    language: string
                }
                cookie: {
                    name: string
                    value: string
                    path: string
                    domain: string
                    expires: string
                }
            }
        }
    }
}