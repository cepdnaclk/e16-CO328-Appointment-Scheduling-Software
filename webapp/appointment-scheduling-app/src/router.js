import {createRouter,createWebHistory} from 'vue-router'

import HomePage from './pages/HomePage.vue'
import AboutUs from './pages/aboutUs.vue'

import AddServicePage from './pages/AddServicePage'
import DashboardPage from './pages/DashboardPage'
import FindServicePage from './pages/FindServicePage'
import OpenLinkPage from './pages/OpenLinkPage'
import SheduledAppoimentsPage from './pages/SheduledAppoimentsPage'

const routes=[
    // dynamic segments start with a colon
    { path: '/', component: HomePage,},
    { path: '/add-service', component: AddServicePage,},
    { path: '/dash-board', component: DashboardPage,},
    { path: '/find-service', component: FindServicePage,},
    { path: '/open-link', component: OpenLinkPage,},
    { path: '/sheduled-appoiments', component: SheduledAppoimentsPage,},

    { path: '/about', component: AboutUs },


]

const router = new createRouter({
    history:createWebHistory(),
    routes,
})


export default router;