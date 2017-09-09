import Vue from 'vue'
import VueRouter from 'vue-router'

import Req from '../vue/req/index.vue'

Vue.use(VueRouter)

const routes = <VueRouter.RouteConfig[]>[
    { path: '/web', name: 'req', component: Req }
]

const Router = new VueRouter({
    mode: 'history',
    routes: routes
})

export default Router