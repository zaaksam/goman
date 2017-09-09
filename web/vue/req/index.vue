<template>
    <div>
        <Tabs v-model="tabID" type="card" closable @on-tab-remove="onTabDel" :animated="false" style="height:100%;">
            <Tab-pane v-for="tab in tabs" v-show="tab.isShow" :name="tab.id" :key="tab.id" :label="tab.label">
                <ReqTab :tab="tab"></ReqTab>
            </Tab-pane>
            <ButtonGroup slot="extra">
                <Button type="ghost" icon="plus-round" @click="onTabAdd" size="large"></Button>
            </ButtonGroup>
        </Tabs>
        <BackTop></BackTop>
    </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import _ from 'lodash'
import moment from 'moment'
import ReqTab from './tab.vue'

@Component({
    components: {
        ReqTab
    }
})
export default class Index extends Vue {
    tabs: Req.TabModel[] = []
    tabID: string = ''

    mounted() {
        this.onTabAdd()
    }

    /**
     * 获取指定的tab对象
     */
    getTab(id: string): Req.TabModel | undefined {
        return _.find(this.tabs, function(v) {
            return v.id === id
        })
    }

    /**
     * 添加tab
     */
    onTabAdd() {
        let id = 'tab' + moment().valueOf().toString()
        let tab: Req.TabModel = {
            id: id,
            isShow: true,
            label: ''
        }
        this.tabs.push(tab)
        this.tabID = id
    }

    /**
     * 移除tab
     */
    onTabDel(id: string) {
        _.remove(this.tabs, (tab) => {
            let has = tab.id === id
            if (has && this.tabs.length == 1) {
                this.onTabAdd()
            }

            return has
        })

        // let tab = this.getTab(id)
        // if (tab) {
        //     let c = _.countBy(this.tabs, (tab) => {
        //         return tab.isShow
        //     })
        //     if (c.true == 1) {
        //         //当前是最后一个tab，新建一个空白的替代
        //         this.onTabAdd()
        //     }

        //     // 只做隐藏，后续可做恢复列表
        //     tab.isShow = false
        // }
    }
}
</script>

