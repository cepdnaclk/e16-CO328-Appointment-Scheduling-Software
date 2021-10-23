import {createRouter,createWebHistory} from 'vue-router'
import store from './store/index'

import HomePage from './pages/HomePage.vue'
import AboutUs from './pages/aboutUs.vue'
import AddServicePage from './pages/AddServicePage'
import DashboardPage from './pages/DashboardPage'
import FindServicePage from './pages/FindServicePage'
import RequestedServicePage from './pages/RequestedServicePage'
//import SheduledAppoimentsPage from './pages/SheduledAppoimentsPage'


function authCheck(to, from,next){

    if(store.state.isLogged){
        next()
    }else{
        next('/')
    }
}

const routes=[
    // dynamic segments start with a colon
    { path: '/', component: HomePage,},
    { path: '/add-service', component: AddServicePage,beforeEnter: authCheck},
    { path: '/dash-board', component: DashboardPage,beforeEnter: authCheck},
    { path: '/find-service', component: FindServicePage,},
    { path: '/requested-service', component: RequestedServicePage,beforeEnter: authCheck},
    //{ path: '/sheduled-appoiments', component: SheduledAppoimentsPage,beforeEnter: authCheck},
    { path: '/about', component: AboutUs },
    
]

const router = new createRouter({
    history:createWebHistory(),
    routes,
})





export default router;