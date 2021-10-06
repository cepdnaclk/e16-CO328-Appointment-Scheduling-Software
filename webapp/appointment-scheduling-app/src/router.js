
import {createRouter,createWebHistory} from 'vue-router'

import HomePage from './pages/HomePage.vue'

const routes=[
    // dynamic segments start with a colon
    { path: '/', component: HomePage },
]

const router = new createRouter({
    history:createWebHistory(),
    routes,
})

export default router;