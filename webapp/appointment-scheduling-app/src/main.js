import { createApp} from 'vue'
import Notifications from 'notiwind'
import Loading from 'vue-loading-overlay';
import VueCookies from 'vue3-cookies'
import App from './App.vue'
import router from './router'
import store from './store/index'

// Import stylesheet
import 'vue-loading-overlay/dist/vue-loading.css';
import './index.css'
import './assets/tailwind.css'
import './assets/homepage.css'



const app=createApp(App)
app.use(router)
app.use(store)
app.use(Loading)
app.use(VueCookies,{
    expireTimes: 60 * 60 * 24 * 0.5,
    path: "/",
    domain: "",
})

app.use(Notifications)
app.mount('#app')
