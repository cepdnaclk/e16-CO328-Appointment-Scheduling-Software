 <template>
   <div class="rounded overflow-hidden shadow-lg">
      <img class="w-full" src="../assets/back1.jpg" alt="Mountain">
      <div class="px-6 py-4">
        <div class="font-bold text-xl mb-2">{{serviceName}}</div>
        <h1 class="text-blue-800 text-base font-black">{{fname}} {{lname}} <span class="text-base text-blue-400">{{ownerEmail}}</span></h1>
        <p class="text-gray-400 text-base">
          {{serviceType}}
        </p>
        <p class="text-gray-700 text-base">
          {{serviceDiscription}}
        </p>
      </div>
      <div class="px-6 pt-4 pb-2">
        <button class="w-full mt-6 mb-3 py-2 text-white font-semibold bg-purple-400 hover:shadow-xl duration-200 hover:bg-purple-800" @click="checkNow">Check Now</button>
      </div>
    </div>
 </template>
 
 <script>


 export default {
    props:[
        "serviceName",
        "serviceDiscription",
        "serviceType",
        "serviceId",
        "fname",
        "lname",
        "checkNow",
        "ownerEmail"
    ],

 }
 </script>
 
 <style>
 
 </style>