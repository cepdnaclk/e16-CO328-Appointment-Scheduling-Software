
import {createRouter,createWebHistory} from 'vue-router'

import HomePage from './pages/HomePage.vue'
import aboutUs from './pages/aboutUs.vue'

const routes=[
    // dynamic segments start with a colon
    { path: '/', component: HomePage },
    { path: '/', component: aboutUs },
]

const router = new createRouter({
    history:createWebHistory(),
    routes,
})

export default router;