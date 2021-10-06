<template>
  <AppHeaderVue @open-login="isOpenLogin=true" @open-signup="isOpenSignup=true"/>
  <router-view/>
  <LoginModal @close-login="isOpenLogin=false" v-if="isOpenLogin"/>
  <SignupModal @close-signup="isOpenSignup=false" v-if="isOpenSignup"/>
</template>

<script>
import AppHeaderVue from './components/AppHeader.vue'
import LoginModal from './pages/LoginModel.vue'
import SignupModal from './pages/SignUpModel.vue'
export default {
 components:{
   AppHeaderVue,
   LoginModal,
   SignupModal
 },
 data(){
   return{
     isOpenLogin:false,
     isOpenSignup:false,
   }
 }
 
}
</script>

<style>
</style>
