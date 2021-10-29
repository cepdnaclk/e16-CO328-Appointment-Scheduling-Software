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
                            Client Name 
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />
                            </svg>
                        </div>
                    </th>
                    <th class="p-2 border-r cursor-pointer text-sm font-thin text-gray-500">
                        <div class="flex items-center justify-center">
                            Email
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
            <AppoimentSlot v-for="(ele,index) in slotListOfService" :approved="ele.approved" 
            :approveFunc="()=>approveSlot(ele.slotId)" :clientEmail="ele.clientEmail" :clientName="ele.clientName" :clientRequested="ele.clientRequested" :removeFunc="()=>removeSlot(ele.slotId)" :time="ele.time" :key="index" data-test="appoiment-card"/>
        </table>
    </div>
</div>
</template>
 
<script>
 import AppoimentSlot from '../components/AppoimentSlot.vue'
 import {getDataFromApiPOST} from '../Functions/dataApi'
 import {APPROVE_CLIENT_REQUEST,REMOVE_CLIENT_REQUEST} from '../constatns'
 export default {
    components:{
        AppoimentSlot
    },
    props:[
    "date",
    "slotList",
    "selecteServiceID",
    ],
    data(){
        return{
            slotListOfService:this.slotList
        }
    },
    methods:{
        removeSlot(id){ 
            getDataFromApiPOST(this,REMOVE_CLIENT_REQUEST,
            {slotId:id,serviceId:this.selecteServiceID,date:this.date},"Done ","Client removed!" ,'',false,true).then(data=>{
                if(data.message){
                 this.$emit("refresh-appoiment-model")
                }
            }).catch(err=>{
                console.log(err)
            })
            
        },
        approveSlot(id){
            getDataFromApiPOST(this,APPROVE_CLIENT_REQUEST,{slotId:id,serviceId:this.selecteServiceID,date:this.date},"Done ","Client Approved!" ,'',false,true).then(data=>{
                if(data.message){
                this.$emit("refresh-appoiment-model")
               }
            }).catch(err=>{
                console.log(err)
            })
        }
    }
 }
 </script>
 
 <style>
 
 </style>