 <template>
 <div class="main-container bg-gradient-to-r from-gray-400 to-gray-200 ">
   <div class="w-1/2">
    <form @submit.prevent="addNewService">
        <div class="my-4">
          <label for="" class=" pb-20">Service Name</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Service Name" v-model.trim="serviceName">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Service Discription</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Service Discription" v-model.trim="serviceDiscription">
        </div>
        <div class=" mt-4">
          <label for="" class=" pb-20">Service Type</label>
          <select class="w-full p-2  text-base text-black transition duration-500 ease-in-out transform rounded-lg bg-gray-100 focus:border-blueGray-500 focus:bg-white focus:outline-none focus:shadow-outline focus:ring-2 ring-offset-current ring-offset-2 mb-3" v-model="serviceType">
            <option value="Once">Once</option>
            <option value="Weekly">Weekly</option>
            <option value="Monthly">Monthly</option>
          </select>
        </div>

          <div class="flex m-1 mb-5">
            <draggable class="dragArea list-group w-full"  v-model="myArray" @change="log">
              <transition-group>
              <div
                class="list-group-item bg-gray-300 m-3 p-3 rounded-md w-full"
                v-for="element in list"
                :key="element.id"
              >
                  <label for="" class=" pb-20 mx-4 fot-black text-sm">APPOIMENT SLOT  <span class="font-thin text-gray-600 text-xs"> Start Time : End Time</span> </label>
                  <div class=" flex">
                  <VueTimepicker class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent mx-2" :format="yourFormat" v-model="element.startTime"></VueTimepicker>
                  <VueTimepicker class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent mx-2" :format="yourFormat" v-model="element.endTime" ></VueTimepicker>
                  <div class="w-auto h-auto">
                    <div class="flex-1 h-full" @click="deleteSlot(element.id)">
                       <div class="flex items-center justify-center flex-1 h-full p-2 bg-red-500 text-white shadow rounded-lg hover:bg-red-300">
                     <div class="relative" >
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" >
                        <path d="M18.693,3.338h-1.35l0.323-1.834c0.046-0.262-0.027-0.536-0.198-0.739c-0.173-0.206-0.428-0.325-0.695-0.325
                        H3.434c-0.262,0-0.513,0.114-0.685,0.312c-0.173,0.197-0.25,0.46-0.215,0.721L2.79,3.338H1.307c-0.502,0-0.908,0.406-0.908,0.908
                        c0,0.502,0.406,0.908,0.908,0.908h1.683l1.721,13.613c0.057,0.454,0.444,0.795,0.901,0.795h8.722c0.458,0,0.845-0.34,0.902-0.795
                        l1.72-13.613h1.737c0.502,0,0.908-0.406,0.908-0.908C19.601,3.744,19.195,3.338,18.693,3.338z M15.69,2.255L15.5,3.334H4.623
                        L4.476,2.255H15.69z M13.535,17.745H6.413L4.826,5.193H15.12L13.535,17.745z" 
                        
                        />
              </svg>
            </div>
          </div>
        </div>
      </div>
                  </div>
              </div>
          </transition-group>
          </draggable>
      </div>
      <div class="flex justify-center w-full mb-6">
      <div class="w-auto h-auto">
        <div class="flex-1 h-full">
          <div class="flex items-center justify-center flex-1 h-full p-2 bg-blue-800 text-white shadow rounded-full hover:text-black  hover:bg-green-500" @click="addNewSlot">
            <div class="relative" >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" @click="addNewSlot"/>
              </svg>
            </div>
          </div>
        </div>
      </div>

      </div>

        <button type="submit" class=" w-full rounded  text-white bg-gradient-to-r from-gray-900 to-pink-800 hover:from-green-200 hover:to-gray-500 p-3 outline-none">
            ADD SERVICE
        </button>
    </form>
 </div>
  
 </div>

 </template>
 
 <script>
  import axios from 'axios'
  import VueTimepicker from 'vue3-timepicker/src/VueTimepicker.vue'
  import { defineComponent } from 'vue'
  import { VueDraggableNext } from 'vue-draggable-next'
  import {ADD_NEW_SERVICE_URL} from '../constatns'
  export default defineComponent({
    components: {
      draggable: VueDraggableNext,
      VueTimepicker
    },
    data(){
      return {
        enabled: true,
        dragging: false,
        list: [
          { 
          startTime: {
            hh: '03',
            mm: '05',
            ss: '00',
            a: 'am'
          },
          endTime:{
            hh: '03',
            mm: '05',
            ss: '00',
            a: 'am'
          }, id: 1 
        }
        ],
        serviceName:"",
        serviceDiscription:"",
        serviceType:"",
        id:1,
        yourFormat: 'hh:mm:ss a',
      }
    },
    methods:{
      log(){

      },
      addNewSlot(){
        this.id=this.id+1
        this.list.push({ 
          startTime: {
            hh: '03',
            mm: '05',
            ss: '00',
            a: 'am'
          },
          endTime:{
            hh: '03',
            mm: '05',
            ss: '00',
            a: 'am'
          }, id: this.id
        })
      },

      addNewService(){
                if(this.serviceName==="" || this.serviceDiscription===""){
                    this.$notify({
                    group: "generic",
                    title: "Required ",
                    text: "Service name and discription are required!"
                    }, 4000)
                    
                }else{
                    let loader = this.$loading.show();
                    axios.post(ADD_NEW_SERVICE_URL, {
                        timeSlots:this.list,
                        serviceName:this.serviceName,
                        serviceDiscription:this.serviceDiscription,
                        serviceType:this.serviceType

                    },{
                    headers: {
                      'Authorization': this.$cookies.get("user_token")
                    }
                    }).then( (response) =>{
                        loader.hide()
                        this.$notify({
                        group: "done",
                        title: "Done ",
                        text: "New Service Added!"
                        }, 4000);
                        this.$router.push('/dash-board')
                        console.log(response);
                    }).catch(error=>{
                        loader.hide()
                        
                        if (error.response && error.response.status==403) {
                            this.$cookies.remove("user_token")
                            this.$cookies.remove("user_type")
                            this.$notify({
                             group: "generic",
                            title: "Session Expires ",
                            text: "Session Expired pleace login"
                            }, 4000)
                            this.$store.commit('setLogedOut')
                            this.$router.push('/')
                        } else {
                            this.$notify({
                            group: "error",
                            title: "Error",
                            text: "Something happened in setting up the request that triggered a Error!",
                            }, 4000)
                        }

                    });
                  }
       },

      deleteSlot(id){
        console.log(id)
        this.list = this.list.filter((obj) => {return obj.id !== id;})
      }
    }
 })
 </script>
 
 <style>

   .main-container{
    padding: 5em;
    height: 110%;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    object-fit: contain;
}

 
 </style>