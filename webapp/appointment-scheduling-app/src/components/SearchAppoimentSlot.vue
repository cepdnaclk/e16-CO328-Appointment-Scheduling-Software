 <template>
   <tr class="bg-gray-100 text-center border-b text-sm text-gray-600">
                    <td class="p-2 border-r">{{time}}</td>
                    <td class="p-2 border-r" data-test="status">{{status}}</td>
                    <td>
                        <button  class="p-2 hover:shadow-lg text-xs font-thin m-1  outline-none " @click="requestBtnFunc" :class="requestBtnDisable" data-test="btn-text">
                            {{requestBtnText}}
                        </button>
                    </td>
    </tr>
 </template>
 
 <script>
 export default {
 props:[
      "time",
      "clientRequested",
      "requestFunc",
  ],
  computed:{
      status(){
          if(this.clientRequested){
              return "Already Requested.."
          }else{
              return "Not Requested Yet"
          }
      },
      requestBtnDisable(){
          if(this.clientRequested){
              return "bg-gray-500 text-black"
          }else{
              return "bg-blue-500  text-white hover:bg-white hover:text-blue-500"
          }
      },
      requestBtnText(){
          if(this.clientRequested){
              return "Not Available"
          }else{
              return "Request"
          }
      },
      requestBtnFunc(){
         if(this.clientRequested){
             return null
         } else{
             return this.requestFunc
         }
      },
  }
 }
 </script>
 
 <style>
 
 </style>