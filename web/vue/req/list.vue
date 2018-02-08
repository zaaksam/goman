<template>
    <div>
        <Table @on-selection-change="onSelect" border :columns="tableColumns" :data="tableData" :no-data-text="$t('goman.general.noData')" :no-filtered-data-text="$t('goman.general.noFilterData')"></Table>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import { CreateElement } from 'vue'
import _ from 'lodash'
import moment from 'moment'

@Component
export default class List extends Vue {
    @Prop() isHeaders: boolean
    @Prop() isSending: boolean = false
    @Prop() contentType: string
    @Prop() listData: Req.ListModel[]

    tableData: any[] = []
    tableColumns: any[] = [
        {
            type: 'selection',
            width: 60,
            align: 'center'
        },
        {
            title: '',
            key: 'key',
            width: 250,
            render: (h: CreateElement, params: any) => {
                return h('AutoComplete', {
                    props: {
                        disabled: this.getDisabled(params.index),
                        value: params.row.key,
                        transfer: true,
                        data: this.getOptions('key'),
                        'filter-method': this.onFilterMethod
                    },
                    on: {
                        'on-change': (value: string) => {
                            this.onChange(params.index, 'key', value)
                        }
                    }
                })
            }
        },
        {
            title: '',
            key: 'value',
            render: (h: CreateElement, params: any) => {
                return h('AutoComplete', {
                    props: {
                        disabled: this.getDisabled(params.index),
                        value: params.row.value,
                        transfer: true,
                        data: this.getOptions('value'),
                        'filter-method': this.onFilterMethod
                    },
                    on: {
                        'on-change': (value: string) => {
                            this.onChange(params.index, 'value', value)
                        }
                    }
                })
            }
        },
        {
            title: '',
            key: 'desc',
            width: 200,
            render: (h: CreateElement, params: any) => {
                return h('Input', {
                    props: {
                        disabled: this.getDisabled(params.index),
                        value: params.row.desc
                    },
                    on: {
                        'on-change': (event: any) => {
                            this.onChange(
                                params.index,
                                'desc',
                                event.target.value
                            )
                        }
                    }
                })
            }
        },
        {
            width: 80,
            renderHeader: (h: CreateElement, params: any) => {
                return h('Button', {
                    props: {
                        disabled: this.isSending,
                        size: 'small',
                        type: 'primary',
                        icon: 'plus-round'
                    },
                    on: {
                        click: () => {
                            this.onAdd()
                        }
                    }
                })
            },
            render: (h: CreateElement, params: any) => {
                return h('Button', {
                    props: {
                        disabled: this.isSending,
                        size: 'small',
                        icon: 'close-round'
                    },
                    on: {
                        click: () => {
                            this.onDel(params.index)
                        }
                    }
                })
            }
        }
    ]

    keyOptions: string[] = []
    valueOptions: string[] = []

    created() {
        this.onLocale()
    }

    @Watch('$i18n.locale')
    onLocale() {
        //注意 tableColumns 的处理位置需要与上面初始化定义一致
        this.tableColumns[1].title = this.$t('goman.list.key')
        this.tableColumns[2].title = this.$t('goman.list.value')
        this.tableColumns[3].title = this.$t('goman.list.description')
    }

    mounted() {
        this.syncTableData()
        this.onContentTypeChange()

        if (this.isHeaders) {
            this.keyOptions = ['Content-Type', 'User-Agent']
            this.valueOptions = [
                'text/plain',
                'text/html',
                'application/json',
                'application/javascript',
                'application/xml',
                'application/x-www-form-urlencoded'
            ]
        }
    }

    @Watch('listData')
    onListDataChange() {
        this.syncTableData()
    }

    @Watch('contentType')
    onContentTypeChange() {
        if (!this.isHeaders) {
            return
        }

        if (this.contentType == '') {
            return
        }

        if (this.contentType == 'text') {
            //text状态下，不加任何content-type
            for (let i = 0; i < this.listData.length; i++) {
                if (this.listData[i].key.toLowerCase() === 'content-type') {
                    this.onDel(i)
                    break
                }
            }

            return
        }

        let isFound = false

        _.forEach(this.listData, v => {
            if (v.key.toLowerCase() === 'content-type') {
                v.value = this.contentType
                isFound = true
                return
            }
        })

        if (!isFound) {
            this.listData.push({
                isDisable: false,
                id:
                    'ls' +
                    moment()
                        .valueOf()
                        .toString(),
                key: 'Content-Type',
                value: this.contentType,
                desc: ''
            })
        }

        this.syncTableData()
    }

    onFilterMethod(value: string, option: string) {
        return option.toLowerCase().indexOf(value.toLowerCase()) != -1
    }

    onSelect(selection: Req.ListModel[]) {
        _.forEach(this.listData, p => {
            p.isDisable = true

            let s = _.find(selection, function(s) {
                return p.id === s.id
            })
            if (s) {
                p.isDisable = false
            }
        })
        this.$emit('onChange')
    }

    onAdd() {
        this.listData.push({
            isDisable: false,
            id:
                'ls' +
                moment()
                    .valueOf()
                    .toString(),
            key: '',
            value: '',
            desc: ''
        })
        this.$emit('onChange')

        this.syncTableData()
    }

    onDel(index: number) {
        // _.remove(this.listData, (n, i) => {
        //     return i == index
        // })

        //删除指定的成员
        this.listData.splice(index, 1)
        this.$emit('onChange')

        this.syncTableData()
    }

    onChange(index: number, key: string, value: string) {
        if (key === 'key') {
            this.listData[index].key = value
        } else if (key === 'value') {
            this.listData[index].value = value
        } else if (key === 'desc') {
            this.listData[index].desc = value
        }
        this.$emit('onChange')
    }

    getOptions(type: 'key' | 'value') {
        return type === 'key' ? this.keyOptions : this.valueOptions
    }

    getDisabled(index: number) {
        if (this.isSending) {
            return true
        }

        return this.listData[index].isDisable
    }

    syncTableData() {
        //全量复制
        this.tableData = this.listData.map((n): any => {
            return {
                key: n.key,
                value: n.value,
                desc: n.desc,
                id: n.id,
                _checked: !n.isDisable
            }
        })
    }
}
</script>

