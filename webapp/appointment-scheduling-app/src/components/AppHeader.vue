<template>
  <nav class="sticky top-0 w-full bg-gradient-to-r from-green-900 to-blue-500 px-4 py-4 text-white flex" style="z-index:90;">
    <div>
      <a href="/" class=" text-black text-xl font-black rounded-lg border-2 px-2 border-solid border-black">AppOIMENTmE</a>
    </div>
      <div class=" px-9 py-2" v-if="isNotLogged">
        <button  class=" px-3 hover:text-gray-500 focus:outline-none " @click="$emit('open-login')">LOGIN</button>
        <button  class=" px-3 hover:text-gray-500 focus:outline-none"  @click="$emit('open-signup')">SIGN UP</button>
      </div>
      <div class=" px-9 py-2" v-if="enableSP">
        <router-link to="/dash-board">
          <button  class=" px-3 hover:text-gray-500 focus:outline-none ">DASHBOARD</button>
        </router-link>
        <router-link to="/add-service">
          <button  class=" px-3 hover:text-gray-500 focus:outline-none" >ADD SERVICE</button>
        </router-link>
        <button  class=" px-3 hover:text-gray-500 focus:outline-none" @click="logout">LOGOUT</button>
      </div>
      <div class=" px-9 py-2" v-if="enableClient">
        <router-link to="/find-service">
        <button  class=" px-3 hover:text-gray-500 focus:outline-none " >SEARCH</button>
        </router-link>
        <router-link to="/open-link">
        <button  class=" px-3 hover:text-gray-500 focus:outline-none" >OPEN LINK</button>
        </router-link>
        <button  class=" px-3 hover:text-gray-500 focus:outline-none" @click="logout">LOGOUT</button>
      </div>
  </nav>
</template>

<script>
export default {
  computed:{
    enableClient(){
      return this.$store.state.isClient && this.$store.state.isLogged
    },
    enableSP(){
      return !this.$store.state.isClient && this.$store.state.isLogged
    },
    isNotLogged(){
      return !this.$store.state.isLogged
    }
  },methods:{
    logout(){
      let loader = this.$loading.show();
      setTimeout(()=>{
        this.$store.commit('setLogedOut')
        loader.hide()
        this.$router.push('/')

      },2000)
      this.$cookies.remove("user_token")
      this.$cookies.remove("user_type")
    }
  }
}
</script>

<style>

</style>
