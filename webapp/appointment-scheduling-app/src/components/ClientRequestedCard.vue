 <template>
<div class="container mx-auto my-1 ">
    <div class="relative rounded-lg flex flex-col md:flex-row items-center md:shadow-xl md:h-72 mx-2">
        <div class="z-0 order-1 md:order-2 relative w-full md:w-2/5 h-20 md:h-full overflow-hidden rounded-lg md:rounded-none md:rounded-r-lg">
            <div class="absolute inset-0 w-full h-full object-fill object-center bg-blue-500 bg-opacity-30 bg-cover bg-bottom" style="background-image: url( https://images.unsplash.com/photo-1525302220185-c387a117886e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1350&q=80 ); background-blend-mode: multiply;"></div>
            <div class="md:hidden absolute inset-0 h-full p-6 pb-6 flex flex-col-reverse justify-start items-start bg-gradient-to-b from-transparent via-transparent to-gray-900">
            </div>
            <svg class="hidden md:block absolute inset-y-0 h-full w-24 fill-current text-white -ml-12" viewBox="0 0 100 100" preserveAspectRatio="none">
                <polygon points="50,0 100,0 50,100 0,100" />
            </svg>
        </div>
        <div class="z-10 order-2 md:order-1 w-full h-full md:w-3/5 flex items-center -mt-6 md:mt-0">
            <div class="p-8 md:pr-18 md:pl-14 md:py-12 mx-2 md:mx-0 h-full bg-white rounded-lg md:rounded-none md:rounded-l-lg shadow-xl md:shadow-none">
                <h4 class="hidden md:block text-xl text-gray-400">{{date}}</h4>
                <h4 class="hidden md:block text-xl text-gray-400">{{time}}</h4>
                <h3 class="hidden md:block font-bold text-2xl text-gray-700">{{serviceName}}</h3>
                <p class="text-gray-600 text-justify">{{servicOwnerEmail}}</p>
                <p class="text-gray-600 text-justify">{{discription}}</p>
                <p class="text-green-600 text-justify">{{getStatus}}</p>
                <button class="flex items-baseline mt-3 text-red-600 hover:text-red-900 focus:text-blue-900 outline-none" @click="onPressCancel">
                    <span>CANCEL REQUEST</span>
                    <span class="text-xs ml-1">&#x279c;</span>
                </button>
            </div>
        </div>
    </div>
</div>
 </template>
 
 <script>
 export default {
    props:{
        "time":String,
        "serviceName":String,
        "approved": Boolean,
        "discription":String,
        "servicOwnerEmail":String,
        "onPressCancel": Function,
        "date":String
    },
    computed:{
        getStatus(){
            if(this.approved){
                return "Approved"
            }else{
                return "Pending"
            }
        }
    }
 }
 </script>
 
 <style>
 
 </style>