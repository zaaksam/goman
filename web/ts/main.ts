import Polyfill from './polyfill'
Polyfill()

import Vue from 'vue'
import iView from 'iview'
import VueI18n from 'vue-i18n'
import Router from './router'

import gomanEN from './i18n/en-US'
import gomanZH from './i18n/zh-CN'

import App from '../vue/app.vue'

Vue.use(VueI18n)
Vue.use(iView)

const i18n = new VueI18n({
    // locale: 'en-US',
    locale: 'zh-CN',
    messages: {
        'en-US': <any>gomanEN,
        'zh-CN': <any>gomanZH
    }
})

new Vue({
    el: '#app',
    // delimiters: ['${', '}'],
    i18n,
    router: Router,
    render: h => h(App)
})