 <template>
   <div class="fixed bottom-0 right-0 w-1/6 m-2 mb-4">
    <div class="flex flex-col items-end">
    <transition name="bounce">
    <div class="flex flex-col w-full" v-if="show">
    <button class="shadow-2xl m-1 w-full rounded  text-white bg-gradient-to-r from-red-800 to-pink-800 hover:from-pink-500 hover:to-yellow-500 p-3 outline-none" :class="{'opacity-50':getClientStatus}" @click="setServiceProvider">Service Provider</button>
    <button class="  shadow-2xl m-1 w-full rounded  text-white bg-gradient-to-r from-red-800 to-pink-800 hover:from-pink-500 hover:to-yellow-500 p-3 outline-none" :class="{'opacity-50':!getClientStatus}" @click="setClient">Client</button>
    </div>
    </transition>
    <div class="rounded-full h-12 w-12 flex  border-none">
    <button class="shadow-2xl m-1 rounded-full h-12 w-12  text-white bg-gradient-to-r from-red-800 to-pink-800 hover:from-pink-500 hover:to-yellow-500 p-3 outline-none fab" @click="show=!show">{{getActivePerson}}</button>
     </div>
</div>
</div>
 </template>
 
 <script>
 
 export default {
    data(){
        return{
            show:false,
        }
    },
    computed:{
        getClientStatus(){
            return this.$store.state.isClient
        },
        getActivePerson(){
            if(this.$store.state.isClient){
                return "C"
            }else{
                return "SP"
            }
        }
    },methods:{
      setClient(){
        this.$store.commit('setClient')
        this.$cookies.set("user_type","Clint")
      },
      setServiceProvider(){
        this.$store.commit('setServiceProvider')
        this.$cookies.set("user_type","SP")
      }
    }
 }
 </script>
 
 <style scopped>
.bounce-enter-active {
animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
 </style>