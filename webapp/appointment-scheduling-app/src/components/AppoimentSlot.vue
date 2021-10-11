 <template>
   <tr class="bg-gray-100 text-center border-b text-sm text-gray-600">
                    <td class="p-2 border-r">{{time}}</td>
                    <td class="p-2 border-r">{{clientName}}</td>
                    <td class="p-2 border-r">{{clientEmail}}</td>
                    <td class="p-2 border-r">{{status}}</td>
                    <td>
                        <a href="#" class="bg-blue-500 p-2 text-white hover:shadow-lg text-xs font-thin">Approve</a>
                        <a href="#" class="bg-red-500 p-2 text-white hover:shadow-lg text-xs font-thin">Remove</a>
                    </td>
    </tr>
 </template>
 
 <script>
 export default {
  props:{
      time,
      clientName,
      clientEmail,
      status,
      edit,
      approveFunc,
      removeFunc
  }
 }
 </script>
 
 <style>
 
 </style>