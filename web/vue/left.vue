<template>
    <Menu active-name="1" theme="dark" width="auto">
        <div class="layout-logo-left"></div>
        <Menu-item name="1">
            <Icon type="ios-navigate" :size="getIconSize"></Icon>
            <span class="layout-text">Option 1</span>
        </Menu-item>
        <Menu-item name="2">
            <Icon type="ios-keypad" :size="getIconSize"></Icon>
            <span class="layout-text">Option 2</span>
        </Menu-item>
        <Menu-item name="3">
            <Icon type="ios-analytics" :size="getIconSize"></Icon>
            <span class="layout-text">Option 3</span>
        </Menu-item>
    </Menu>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class Left extends Vue {
    @Prop()
    app: AppModel

    get getIconSize() {
        return this.app.spanLeft === 4 ? 14 : 24
    }
}
</script>
