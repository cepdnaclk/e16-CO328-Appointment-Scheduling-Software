 <template>
   <tr class="bg-gray-100 text-center border-b text-sm text-gray-600">
                    <td class="p-2 border-r">{{time}}</td>
                    <td class="p-2 border-r">{{clientRequested ?clientName:""}}</td>
                    <td class="p-2 border-r">{{clientRequested ?clientEmail:""}}</td>
                    <td class="p-2 border-r">{{status}}</td>
                    <td>
                        <button  class="p-2 hover:shadow-lg text-xs font-thin m-1  outline-none " @click="approveBtnFunc" :class="approveDisable">
                            {{approveBtnText}}
                        </button>
                        <button href="#" class="bg-red-500 p-2 text-white hover:shadow-lg text-xs font-thin m-1 hover:bg-white hover:text-red-500 outline-none" @click="removeFunc">Remove</button>
                    </td>
    </tr>
 </template>
 
 <script>
 export default {
  props:[
      "time",
      "clientName",
      "clientEmail",
      "clientRequested",
      "approved",
      "approveFunc",
      "removeFunc",
  ],
  computed:{
      status(){
          if(this.clientRequested && !this.approved){
              return "Pending.."
          }else if(this.clientRequested && this.approved){
              return "Approved"
          }else{
              return "Not Requested Yet"
          }
      },
      approveDisable(){
          if(this.clientRequested && this.approved){
              return "bg-gray-500 text-black"
          }else{
              return "bg-blue-500  text-white hover:bg-white hover:text-blue-500"
          }
      },
      approveBtnText(){
          if(this.clientRequested && this.approved){
              return "Approved"
          }else{
              return "Approve"
          }
      },
      approveBtnFunc(){
         if(this.clientRequested && this.approved){
             return null
         } else{
             return this.approveFunc
         }
      }
  }
 }
 </script>
 
 <style>
 
 </style>