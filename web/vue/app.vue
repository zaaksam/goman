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
    mounted() {
        if (C.runMode === 'web') {
            this.heartbeat()
        }
    }

    heartbeat() {
        //发送心跳包
        setTimeout(async () => {
            await API.get<any>('/app')
            this.heartbeat()
        }, 1000)
    }
}
</script>
