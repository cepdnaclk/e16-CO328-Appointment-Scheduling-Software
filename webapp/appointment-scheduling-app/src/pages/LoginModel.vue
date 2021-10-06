 <template>
 <BlueModal eventName="close-login">
   <div class="p-2">
    <h1 class="font-black text-center antialiased ">LOG IN</h1>
    <form @submit.prevent="login">
        <div class="my-4">
          <label for="" class=" pb-20">Email</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Email" v-model.trim="email">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Password</label>
          <input type="password" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Password" v-model.trim="password">
        </div>
        <button type="submit" class=" w-full rounded  text-white bg-gradient-to-r from-red-800 to-pink-800 hover:from-pink-500 hover:to-yellow-500 p-3">
            LOG IN
        </button>
    </form>
    </div>
 </BlueModal>
 </template>
 
<script>
    import BlueModal from '../components/BlueModal.vue'
    import axios from 'axios'
    export default {
        components:{
            BlueModal,
        },
        data(){
            return{
                email:"",
                password:"",
            }
        },
        methods:{
            login(){ 
                if(this.email==="" || this.password===""){
                    this.$notify({
                    group: "generic",
                    title: "Required ",
                    text: "Email and Password are required!"
                    }, 4000)
                    
                }else{
                    let loader = this.$loading.show();
                    axios.post('login', {
                        email: this.email,
                        password:this.password
                    })
                    .then( (response) =>{
                        loader.hide()
                        console.log(response);
                    })
                    .catch(function (error) {
                        loader.hide()
                        if (error.response && error.response.status==401) {
                            this.$notify({
                            group: "error",
                            title: "Login Failed",
                            text: "Your email or password is incorrect"
                            }, 4000)
                        } else {
                            this.$notify({
                            group: "error",
                            title: "Error",
                            text: "Something happened in setting up the request that triggered a Error!",
                            }, 4000)
                        }
                    });
                }
            }
        }
    }
</script>
 
 <style>
 
 </style>