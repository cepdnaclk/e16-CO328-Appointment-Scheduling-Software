 <template>
    <BlueModal eventName="close-signup">
   <div class="p-2">
    <h1 class="font-black text-center antialiased ">SIGN UP</h1>
    <form action="">
        <div class="my-4">
          <label for="" class=" pb-20">Frist Name</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Frist Name">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Last Name</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Last Name">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Email</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Email">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Password</label>
          <input type="password" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Password">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Confirm Password</label>
          <input type="password" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Re Enter Your Password">
        </div>
        <button type="submit" class=" w-full rounded  text-white bg-gradient-to-r from-red-800 to-pink-800 hover:from-pink-500 hover:to-yellow-500 p-3">
            SIGN UP
        </button>
    </form>
    </div>
 </BlueModal>
 </template>
 
 <script>
    import BlueModal from '../components/BlueModal.vue'
    export default {
        components:{
            BlueModal,
        }
    }
 </script>
 
 <style>
 
 </style>