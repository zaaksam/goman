<template>
    <div>
        <Form>
            <Form-item>
                <Row type="flex">
                    <Col :xs="18" :sm="19" :md="20" :lg="21">
                    <Input v-model="req.url" placeholder="请输入..." style="width:100%;" :disabled="isSending" @on-change="onURLChange">
                    <Select v-model="req.method" slot="prepend" style="width:100px" :disabled="isSending">
                        <Option v-for="item in methodOptions" :value="item" :key="item">{{ item }}</Option>
                    </Select>
                    <Button slot="append" style="width:100px" :disabled="isSending" @click="onParamsShow">Params</Button>
                    </Input>
                    </Col>
                    <Col :xs="6" :sm="5" :md="4" :lg="3" style="display:flex;justify-content:flex-end;">
                    <Button-group>
                        <Button @click="onSend" type="primary" icon="android-send" :loading="isSending">SEND</Button>
                        <Button type="primary" icon="ios-arrow-down" :disabled="true"></Button>
                    </Button-group>
                    </Col>
                </Row>
            </Form-item>
            <Form-item v-show="showParams">
                <ReqList :listData="params" :isSending="isSending" @onChange="onParamsChange"></ReqList>
            </Form-item>
            <Form-item>
                <Card style="width:100%;">
                    <p slot="title">Request</p>
                    <p>
                        <Tabs :animated="false" v-model="reqMode">
                            <Tab-pane label="Body" name="Body" :disabled="!isBody">
                                <div style="padding-bottom:10px;">
                                    <RadioGroup v-model="bodyMode" size="large">
                                        <Radio label="normal">x-www-form-urlencoded</Radio>
                                        <Radio label="raw">raw</Radio>
                                    </RadioGroup>
                                    <Select v-model="rawContentType" :disabled="!isBodyRaw" style="width:220px;" :transfer="true">
                                        <Option value="text">Text</Option>
                                        <Option value="text/plain">Text(text/plain)</Option>
                                        <Option value="application/json">JSON(application/json)</Option>
                                        <Option value="application/javascript">Javascript(application/javascript)</Option>
                                        <Option value="application/xml">XML(application/xml)</Option>
                                        <Option value="text/xml">XML(text/xml)</Option>
                                        <Option value="text/html">HTML(text/html)</Option>
                                    </Select>
                                </div>
                                <ReqList v-show="!isBodyRaw" :listData="bodys" :isSending="isSending"></ReqList>
                                <Input v-show="isBodyRaw" v-model="body" type="textarea" :rows="10"></Input>
                            </Tab-pane>
                            <Tab-pane :label="headersLabel" name="Headers">
                                <ReqList :listData="headers" :isSending="isSending" :isHeaders="true" :contentType="contentType"></ReqList>
                            </Tab-pane>
                            <Tab-pane label="Code" name="Code">
                                <pre class="mypre">{{ code }}</pre>
                            </Tab-pane>
                        </Tabs>
                    </p>
                </Card>
            </Form-item>
            <Form-item>
                <ReqPre :res="res"></ReqPre>
            </Form-item>
        </Form>
    </div>
</template>

<style>
.mypre {
    overflow: auto;
    overflow-y: hidden;
}
</style>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import _ from 'lodash'
import moment from 'moment'
import API from '../../ts/api'
import ReqList from './list.vue'
import ReqPre from './pre.vue'

interface IOptions {
    label: string
    value: string

}

@Component({
    components: {
        ReqList,
        ReqPre
    }
})
export default class Tab extends Vue {
    @Prop()
    tab: Req.TabModel

    isSending: boolean = false
    reqMode: 'Body' | 'Headers' | 'Code' = 'Headers'
    contentType: string = ''
    headers: Req.ListModel[] = [
        {
            isDisable: false,
            id: 'ls0',
            key: 'User-Agent',
            value: C.userAgent,
            desc: ''
        }
    ]
    params: Req.ListModel[] = []
    bodys: Req.ListModel[] = []
    body: string = ''
    bodyMode: 'normal' | 'raw' = 'normal'

    req: Req.RequestModel = {
        method: 'GET',
        url: '',
        headers: {},
        body: ''
    }
    res: Req.ResponseModel | null = null

    methodOptions: string[] = ['GET', 'POST', 'HEAD', 'PUT', 'DELETE']
    rawContentType: string = 'text'
    showParams: boolean = true

    mounted() {
        this.onTabLabelChange(this.req.url)
        setTimeout(() => {
            this.onParamsShow()
        }, 1)
    }

    onParamsShow() {
        this.showParams = !this.showParams
    }

    /**
     * 根据params变动，更新url及tab.label
     */
    onParamsChange() {
        let url = this.req.url
        let index = url.indexOf('?')
        if (index != -1) {
            url = url.substr(0, index)
        }

        let pmStr = ''
        _.forEach(this.params, (pm, k) => {
            if (pm.isDisable) {
                return
            }

            if (pmStr != '') {
                pmStr += '&'
            }
            pmStr += pm.key + '=' + pm.value
        })

        this.req.url = pmStr == '' ? url : url + '?' + pmStr
        this.onTabLabelChange(this.req.url)
    }

    /**
     * 根据url变动，更新params及tab.label
     */
    onURLChange(event: any) {
        this.onTabLabelChange(this.req.url)

        if (this.req.url == '') {
            return
        }

        let url = this.req.url
        let index = url.indexOf('?')
        if (index == -1) {
            this.params = []
            return
        }

        let pmStr = url.substr(index + 1)
        if (pmStr == '') {
            return
        }

        let params: Req.ListModel[] = []
        let pms = pmStr.split('&')
        //遍历参数
        _.forEach(pms, (pm) => {
            if (pm == '') {
                return
            }

            let key = ''
            let value = ''
            index = pm.indexOf('=')
            if (index == -1) {
                key = pm
            } else {
                key = pm.substr(0, index)
                value = pm.substr(index + 1)
            }

            params.push({
                isDisable: false,
                id: 'ls' + moment().valueOf().toString(),
                key: key,
                value: value,
                desc: ''
            })
        })

        this.params = params
    }

    onTabLabelChange(label: string) {
        if (label.trim() == '') {
            label = 'New Tab'
        }

        this.tab.label = label.length > 18 ? label.substr(0, 18) : label
    }

    async onSend() {
        this.isSending = true
        this.res = null

        this.req.headers = {}
        _.forEach(this.headers, (v) => {
            if (!v.isDisable) {
                if (this.req.headers[v.key]) {
                    this.req.headers[v.key].push(v.value)
                } else {
                    this.req.headers[v.key] = [v.value]
                }
            }
        })

        this.req.body = ''
        if (this.isBody) {
            if (this.isBodyRaw) {
                this.req.body = this.body
            } else {
                _.forEach(this.bodys, (v) => {
                    if (!v.isDisable) {
                        if (this.req.body != '') {
                            this.req.body += '&'
                        }

                        this.req.body += v.key + '=' + encodeURIComponent(v.value)
                    }
                })
            }
        }

        let data = new URLSearchParams()
        data.append('req', JSON.stringify(this.req))

        let result = await API.post<Req.ResponseModel>('/req', data)

        if (result.code == 10000) {
            let res = result.data
            if (res) {
                this.res = res
            }
        } else {
            this.$Message.error({ duration: 5, content: result.msg + '(' + result.code.toString() + ')' })
        }

        this.isSending = false
    }

    get headersLabel(): string {
        let label = 'Headers'

        if (this.headers.length > 0) {
            label += '(' + this.headers.length.toString() + ')'
        }

        return label
    }

    get isBody(): boolean {
        let isBody = this.req.method != 'GET' && this.req.method != 'HEAD'
        if (!isBody && this.reqMode === 'Body') {
            this.reqMode = 'Headers'
        }

        return isBody
    }

    get isBodyRaw(): boolean {
        let isBodyRaw = this.bodyMode == 'raw'

        if (this.isBody) {
            this.contentType = isBodyRaw ? this.rawContentType : 'x-www-form-urlencoded'
        }

        return isBodyRaw
    }

    get code(): string {
        let urls: string[]
        let path: string
        let hosts = this.req.url.split('://')
        if (hosts.length > 1) {
            urls = hosts[1].split('/')
            path = hosts[1]
        } else {
            urls = hosts[0].split('/')
            path = hosts[0]
        }

        let host = urls[0]
        path = urls.length > 1 ? _.replace(path, host, '') : ''

        let code = this.req.method + ' ' + path + ' ' + 'HTTP/1.1'
        code += '\nHost: ' + host

        _.forEach(this.headers, (v) => {
            if (!v.isDisable) {
                code += '\n' + v.key + ': ' + v.value
            }
        })

        code += '\n'

        if (this.isBody) {
            if (this.isBodyRaw) {
                code += '\n' + this.body
            } else {
                let i = 0
                _.forEach(this.bodys, (v) => {
                    if (!v.isDisable) {
                        code += i == 0 ? '\n' : '&'
                        code += v.key + '=' + encodeURIComponent(v.value)
                        i++
                    }
                })
            }
        }

        return code
    }
}
</script>
