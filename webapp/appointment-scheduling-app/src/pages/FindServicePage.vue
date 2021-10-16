<template>
   <div class="container-bar">
        <div class="relative">
            <div class="input-tab"> <i class="fa fa-search text-gray-400 z-20 hover:text-gray-500"></i> </div> <input type="text" class="h-14 w-96 pl-10 pr-20 rounded-lg z-0 focus:shadow focus:outline-none" placeholder="Search anything...">
            <div class="searchbutton"> <button class="button-s" type='submit'>Search</button> </div>
        </div>
    </div>
    <div class="p-10 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-3 lg:grid-cols-3 xl:grid-cols-3 gap-5">
    <card serviceDiscription="bbbbbbbb" serviceId="000" serviceName="aseddddmm" serviceType="Monthly"/>
    <card />
    <card />
    <card />
    <card />
    <card />
    </div>
  
</template>

<script>
import card from '../components/CardClient.vue'
export default {
  components:{
    card
  }
}
</script>

<style>
@import url('../assets/findservice.css');
</style>