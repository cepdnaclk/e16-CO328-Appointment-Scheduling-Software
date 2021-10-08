import { createApp } from 'vue'
import Notifications from 'notiwind'
import Loading from 'vue-loading-overlay';
// Import stylesheet
import 'vue-loading-overlay/dist/vue-loading.css';
import App from './App.vue'
import './index.css'
import './assets/tailwind.css'
import './assets/homepage.css'
import './assets/about.css'
import router from './router'

const app=createApp(App)
app.use(router)
app.use(Loading)
app.use(Notifications)
app.mount('#app')
