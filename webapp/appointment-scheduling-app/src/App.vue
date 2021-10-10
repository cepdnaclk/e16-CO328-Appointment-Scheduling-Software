<template>
  <AppHeaderVue @open-login="isOpenLogin=true" @open-signup="isOpenSignup=true"/>
  <router-view/>
  <Footer/>
  <LoginModal @close-login="isOpenLogin=false" v-if="isOpenLogin"/>
  <SignupModal @close-signup="isOpenSignup=false" v-if="isOpenSignup"/>

  <NotificationGroup group="error">
    <div
      class="fixed inset-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none "
      style="z-index:111"
    >
      <div class="w-full max-w-sm">
        <Notification
          v-slot="{ notifications }"
          enter="transform ease-out duration-300 transition"
          enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0"
          leave="transition ease-in duration-500"
          leave-from="opacity-100"
          leave-to="opacity-0"
          move="transition duration-500"
          move-delay="delay-300"
        >
          <div
            class="flex w-full max-w-sm mx-auto mt-4 overflow-hidden bg-white rounded-lg shadow-md"
            v-for="notification in notifications"
            :key="notification.id"
          >
          <div class="flex items-center justify-center w-12 bg-red-500">
            <svg class="w-6 h-6 text-white fill-current" viewBox="0 0 40 40" xmlns="http://www.w3.org/2000/svg">
              <path d="M20 3.36667C10.8167 3.36667 3.3667 10.8167 3.3667 20C3.3667 29.1833 10.8167 36.6333 20 36.6333C29.1834 36.6333 36.6334 29.1833 36.6334 20C36.6334 10.8167 29.1834 3.36667 20 3.36667ZM19.1334 33.3333V22.9H13.3334L21.6667 6.66667V17.1H27.25L19.1334 33.3333Z"/>
            </svg>
          </div>

          <div class="px-4 py-2 -mx-3">
            <div class="mx-3">
              <span class="font-semibold text-red-500">{{ notification.title }}</span>
              <p class="text-sm text-gray-600">{{ notification.text }}</p>
            </div>
          </div>
          </div>
        </Notification>
      </div>
    </div>
  </NotificationGroup>

  <NotificationGroup group="generic">
    <div
      class="fixed inset-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none  "
      style="z-index:111"
    >
      <div class="w-full max-w-sm">
        <Notification
          v-slot="{ notifications }"
          enter="transform ease-out duration-300 transition"
          enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0"
          leave="transition ease-in duration-500"
          leave-from="opacity-100"
          leave-to="opacity-0"
          move="transition duration-500"
          move-delay="delay-300"
        >
          <div
            class="flex w-full max-w-sm mx-auto mt-4 overflow-hidden bg-white rounded-lg shadow-md"
            v-for="notification in notifications"
            :key="notification.id"
          >
            <div class="flex items-center justify-center w-12 bg-blue-500">
              <svg class="w-6 h-6 text-white fill-current" viewBox="0 0 40 40" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 3.33331C10.8 3.33331 3.33337 10.8 3.33337 20C3.33337 29.2 10.8 36.6666 20 36.6666C29.2 36.6666 36.6667 29.2 36.6667 20C36.6667 10.8 29.2 3.33331 20 3.33331ZM21.6667 28.3333H18.3334V25H21.6667V28.3333ZM21.6667 21.6666H18.3334V11.6666H21.6667V21.6666Z"/>
              </svg>
            </div>

            <div class="px-4 py-2 -mx-3">
              <div class="mx-3">
                <span class="font-semibold text-blue-500">{{ notification.title }}Info</span>
                <p class="text-sm text-gray-600">{{ notification.text }}</p>
              </div>
            </div>
          </div>
        </Notification>
      </div>
    </div>
</NotificationGroup>
  <NotificationGroup group="done">
    <div
      class="fixed inset-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none  "
      style="z-index:111"
    >
      <div class="w-full max-w-sm">
        <Notification
          v-slot="{ notifications }"
          enter="transform ease-out duration-300 transition"
          enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0"
          leave="transition ease-in duration-500"
          leave-from="opacity-100"
          leave-to="opacity-0"
          move="transition duration-500"
          move-delay="delay-300"
        >
          <div
            class="flex w-full max-w-sm mx-auto mt-4 overflow-hidden bg-white rounded-lg shadow-md"
            v-for="notification in notifications"
            :key="notification.id"
          >
            <div class="flex items-center justify-center w-12 bg-green-500">
              <svg class="w-6 h-6 text-white fill-current" viewBox="0 0 40 40" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 3.33331C10.8 3.33331 3.33337 10.8 3.33337 20C3.33337 29.2 10.8 36.6666 20 36.6666C29.2 36.6666 36.6667 29.2 36.6667 20C36.6667 10.8 29.2 3.33331 20 3.33331ZM16.6667 28.3333L8.33337 20L10.6834 17.65L16.6667 23.6166L29.3167 10.9666L31.6667 13.3333L16.6667 28.3333Z"/>
              </svg>
            </div>

            <div class="px-4 py-2 -mx-3">
              <div class="mx-3">
                <span class="font-semibold text-green-500">{{ notification.title }}Info</span>
                <p class="text-sm text-gray-600">{{ notification.text }}</p>
              </div>
            </div>
          </div>
        </Notification>
      </div>
    </div>
</NotificationGroup>
<FloatingButton v-if="getLoginStatus"/>
</template>

<script>
import AppHeaderVue from './components/AppHeader.vue'
import LoginModal from './pages/LoginModel.vue'
import SignupModal from './pages/SignUpModel.vue'
import FloatingButton from './components/FloatingButton.vue'
import Footer from './components/footer.vue'


export default {
 components:{
   AppHeaderVue,
   LoginModal,
   SignupModal,
   FloatingButton,
   Footer,
 },
 data(){
   return{
     isOpenLogin:false,
     isOpenSignup:false,
   }
 },
 computed:{
   getLoginStatus(){
     return this.$store.state.isLogged
   }
 },
 beforeCreate(){
   if (this.$cookies.isKey("user_token")) {
     this.$store.commit('setLogedIn')
     if (this.$cookies.get("user_type")==="SP") {
       this.$store.commit('setServiceProvider')
       
     }else{
       this.$store.commit('setClient')
     }
   }
 }
 
}
</script>

<style>
</style>
