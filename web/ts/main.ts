import Vue from 'vue'
import iView from 'iview'
import Router from './router'

import App from '../vue/app.vue'

Vue.use(iView)

new Vue({
    el: '#app',
    // delimiters: ['${', '}'],
    router: Router,
    render: h => h(App)
})