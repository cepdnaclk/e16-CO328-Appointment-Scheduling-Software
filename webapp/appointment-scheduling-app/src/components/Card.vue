 <template>
   <div class="flex flex-col sm:flex-col lg:flex-row xl:flex-row md:flex-row justify-center items center  container"  >
                <div class="py-12 sm:py-12 md:py-6 lg:py-6 xl:py-6 px-8 w-full md:max-w-min sm:w-full bg-purple-500 transform scale-1 sm:scale-1 md:scale-105 lg:scale-105 xl:scale-105 z-40  shadow-none sm:shadow-none md:shadow-xl lg:shadow-xl xl:shadow-xl">
                    <h1 class="text-purple-200 font-semibold text-l ">{{serviceType}}</h1>
                    <div class="text-center py-4 px-7">
                        <h1 class="text-white text-xl font-black">{{serviceName}}</h1>
                    </div>
                    <div class="h-px bg-purple-400"></div>
                    <div class="text-center mt-3">
                        <p class="text-sm text-white text-opacity-80">
                       {{serviceDiscription}}
                        </p>
                    </div>
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
        "checkNow"
    ],

 }
 </script>
 
 <style>
 
 </style>