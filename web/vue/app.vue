<template>
    <div class="layout">
        <Layout>
            <Header :style="{position:'fixed',width:'100%','z-index':99}">
                <MyTop></MyTop>
            </Header>
            <Content :style="{margin:'50px 0px 0','z-index':0}">
                <div class="layout-content">
                    <div class="layout-content-main">
                        <router-view></router-view>
                    </div>
                </div>
            </Content>
        </Layout>
        <Modal v-model="showClose" :title="$t('goman.app.modal.title')" :ok-text="$t('goman.app.modal.okText')" :cancel-text="$t('goman.app.modal.cancelText')" @on-ok="onClose">
            <p>{{this.$t('goman.app.modal.content')}}</p>
        </Modal>
    </div>
</template>

<style>
.layout {
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    overflow: hidden;
}

.layout-content {
    min-height: 200px;
    margin: 15px;
    overflow: hidden;
    background: #fff;
    border-radius: 4px;
}

.layout-content-main {
    padding: 10px;
}
</style>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MyTop from './top.vue'
import API from '../ts/api'

@Component({
    components: {
        MyTop
    }
})
export default class App extends Vue {
    showClose = false

    mounted() {
        if (C.runMode === 'web') {
            this.heartbeat()
        }
    }

    onClose() {
        window.opener = null
        window.open('', '_self')
        window.close()
    }

    heartbeat() {
        //发送心跳包
        setTimeout(async () => {
            let result = await API.get<any>(
                '/app?t=' + Math.round(new Date().getTime() / 1000).toString()
            )
            if (result.code == -1) {
                this.showClose = true
                return
            }

            this.heartbeat()
        }, 1000)
    }
}
</script>
