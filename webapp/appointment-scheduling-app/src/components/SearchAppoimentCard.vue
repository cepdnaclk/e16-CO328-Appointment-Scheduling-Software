<template>
   <div >
      <h2 class=" text-center bg-white ">{{date}}</h2>
      <div class="table w-1/12 p-2">
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-50 border-b">
                    <th class="p-2 border-r cursor-pointer text-sm font-thin text-gray-500">
                        <div class="flex items-center justify-center">
                            Time
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />
                            </svg>
                        </div>
                    </th>
                    <th class="p-2 border-r cursor-pointer text-sm font-thin text-gray-500">
                        <div class="flex items-center justify-center">
                            Status
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />
                            </svg>
                        </div>
                    </th>
                    <th class="p-2 border-r cursor-pointer text-sm font-thin text-gray-500">
                        <div class="flex items-center justify-center">
                            Action
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />
                            </svg>
                        </div>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr class="bg-gray-50 text-center">
                    <td class="p-2 border-r">
                        <input type="text" class="border p-1" disabled>
                    </td>
                    <td class="p-2 border-r">
                        <input type="text" class="border p-1" disabled>
                    </td>
                    <td class="p-2">
                        <input type="text" class="border p-1" disabled>
                    </td>   
                </tr>
            </tbody>
            <SearchAppoimentSlot v-for="(ele,index) in slotListOfService"  
            :requestFunc="()=>requestSlot(ele.slotId)"  :clientRequested="ele.clientRequested" :time="ele.time" :key="index" />
        </table>
    </div>
</div>
</template>
 
 <script>
 import {getDataFromApiPOST} from '../Functions/dataApi'
 import {CLIENT_REQUEST_SERVICE_SLIOT_URL} from '../constatns'
 import SearchAppoimentSlot from '../components/SearchAppoimentSlot.vue'
 export default {
    components:{
        SearchAppoimentSlot
    },
    props:[
    "date",
    "slotList",
    "selecteServiceID",
    "ownerEmail"
    ],
    computed:{
        slotListOfService(){
            return this.slotList
        }
    },
    methods:{
        requestSlot(slotid){
            getDataFromApiPOST(this,CLIENT_REQUEST_SERVICE_SLIOT_URL,{serviceOwnerEmail:this.ownerEmail,date:this.date,slotId:slotid,serviceId:this.selecteServiceID},"","",'',false,false).then(data=>{
                if (!data.message) {
                    this.$emit("request-failed-alert")
                }else{
                    this.$emit("done")
                }
            }).catch(err=>{
                
            })
            
        }
    }
}
 </script>
 
 <style>
 
 </style>