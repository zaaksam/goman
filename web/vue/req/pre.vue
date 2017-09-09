<template>
    <Card style="width:100%;">
        <p slot="title">Response</p>
        <p slot="extra" v-html="status"></p>
        <p>
            <Tabs v-show="res" :animated="false" v-model="resMode">
                <Tab-pane label="Body">
                    <RadioGroup v-model="bodyMode" type="button" size="large">
                        <Radio label="pretty">pretty</Radio>
                        <Radio label="raw">raw</Radio>
                    </RadioGroup>
                    <span>&nbsp;&nbsp;&nbsp;&nbsp;language：{{ bodyLanguage }}</span>
                    <pre v-show="bodyMode==='raw'">{{ bodyRaw }}</pre>
                    <pre v-show="bodyMode==='pretty'" v-html="body"></pre>
                </Tab-pane>
                <Tab-pane label="Headers">
                    <Table stripe border :columns="headerColumns" :data="headers"></Table>
                </Tab-pane>
                <Tab-pane label="Cookies">
                    <Table stripe border :columns="cookieColumns" :data="cookies"></Table>
                </Tab-pane>
                <Tab-pane label="Code">
                    <pre>{{ code }}</pre>
                </Tab-pane>
            </Tabs>

        </p>
    </Card>
</template>

<style>
.hljs-ul {
    list-style: decimal;
    background-color: #fff;
    margin: 0px 0px 0 40px !important;
    padding: 0px;
}

.hljs-ul li {
    list-style: decimal-leading-zero;
    border-left: 1px solid #ddd !important;
    background: #fff;
    padding: 5px!important;
    margin: 0 !important;
    line-height: 14px;
    word-break: break-all;
    word-wrap: break-word;
}

.hljs-ul li:nth-of-type(even) {
    background-color: #f8f8f9;
    color: inherit;
}
</style>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import _ from 'lodash'
import hljs from 'highlight.js'

@Component
export default class Pre extends Vue {
    @Prop()
    res: Req.ResponseModel | null

    headerColumns = <any[]>[
        {
            title: 'Key',
            key: 'key'
        },
        {
            title: 'Value',
            key: 'value'
        }
    ]
    cookieColumns = <any[]>[
        {
            title: 'Name',
            key: 'Name'
        },
        {
            title: 'Value',
            key: 'Value'
        },
        {
            title: 'Path',
            key: 'Path'
        },
        {
            title: 'Domain',
            key: 'Domain'
        },
        {
            title: 'Expires',
            key: 'Expires'
        }
    ]

    bodyMode: 'pretty' | 'raw' = 'pretty'
    bodyLanguage: string = ''
    isPretty: boolean = true
    isRaw: boolean = false

    resMode: number = 0

    mounted() {
        this.resChange()
    }

    @Watch('res')
    resChange() {
        this.bodyMode = 'pretty'
        this.resMode = 0
    }

    get status(): string {
        if (!this.res) {
            return ''
        }

        let status: string = 'Proto：<span style="color:#3399ff;">' + this.res.proto + '</span>'
        status += '&nbsp;&nbsp;&nbsp;&nbsp;Status：<span style="color:#3399ff;">' + this.res.status + '</span>'
        status += '&nbsp;&nbsp;&nbsp;&nbsp;Time：<span style="color:#3399ff;">' + this.res.duration + '</span>'

        return status
    }

    get bodyRaw(): string {
        if (!this.res || !this.res.body) {
            return ''
        }

        return this.res.body
    }

    get body(): string {
        if (!this.res || !this.res.body) {
            return ''
        }

        let body = this.res.body
        let contentType = 'text/plain'

        _.forEach(this.res.headers, (v, k) => {
            k = k.toLowerCase()
            if (k == 'content-type') {
                contentType = v[0]
            }
            return
        })

        let language: string[] | undefined = undefined
        if (contentType.indexOf('application/json') != -1) {
            body = JSON.stringify(JSON.parse(body), undefined, 4)
            language = ['json']
        } else if (contentType.indexOf('application/xml') != -1) {
            language = ['xml']
        } else if (contentType.indexOf('text/html') != -1) {
            language = ['html']
        }

        let pretty = hljs.highlightAuto(body, language)
        this.bodyLanguage = pretty.language
        if (!language) {
            this.bodyLanguage += '(auto)'
        }
        body = '<ul class="hljs-ul"><li>' + _.replace(pretty.value, /\n/g, '\n</li><li>') + '</li></ul>'

        return body
    }

    get headers(): Req.HeaderModel[] {
        if (!this.res) {
            return []
        }

        let headers: Req.HeaderModel[] = []

        _.forIn(this.res.headers, (va, k) => {
            // headers += k + '：' + _.join(v, ';') + '\r\n'
            _.forEach(va, (v) => {
                headers.push({
                    key: k,
                    value: v
                })
            })
        })

        return headers
    }

    get cookies(): Req.CookieModel[] {
        return this.res ? this.res.cookies : []

        // if (!this.res) {
        //     return ''
        // }

        // let cookies: string = ''

        // _.forEach(this.res.Cookies, (v) => {
        //     cookies += JSON.stringify(v) + '\r\n'
        // })

        // return cookies
    }

    get code(): string {
        return this.res ? this.res.code + '[body ignore]' : ''
    }

}
</script>
