 <template>
    <BlueModal eventName="close-signup" customClases="bg-blue-100 m-auto">
   <div class="p-2">
    <h1 class="font-black text-center antialiased ">SIGN UP</h1>
    <form @submit.prevent="signup">
        <div class="my-4">
          <label for="" class=" pb-20">Frist Name</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Frist Name" v-model.trim="fristName">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Last Name</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Last Name" v-model.trim="lastName">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Email</label>
          <input type="text" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Email" v-model.trim="email">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">WhatsAPP Number</label>
          <input type="tel" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your WhatsAPP No" v-model.trim="phone">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Password</label>
          <input type="password" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Enter Your Password" v-model.trim="password">
        </div>
        <div class="my-4">
          <label for="" class=" pb-20">Confirm Password</label>
          <input type="password" class=" rounded shadow-md p-2 w-full border border-transparent focus:outline-none focus:ring-2 focus:ring-gray-600 focus:border-transparent"
          placeholder="Re Enter Your Password" v-model.trim="repassword">
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
    import axios from 'axios'
    export default {
        components:{
            BlueModal,
        },
        data(){
          return{
            fristName:"",
            lastName:"",
            password:"",
            repassword:"",
            email:"",
            phone:""
          }
        },
        methods:{
          signup(){
            
            if(this.password!==this.repassword){
              this.$notify({
              group: "generic",
              title: "Password ",
              text: "Password are not matching!"
              }, 4000);

              return
            }
            if(this.email==="" || this.password==="" || this.lastName==="" || this.fristName===""){
              this.$notify({
              group: "generic",
              title: "Required ",
              text: "All fields are required!"
              }, 4000)
                    
            }else{
              let loader = this.$loading.show();
              axios.post('signup', {
                  email: this.email,
                  password:this.password,
                  fristName:this.fristName,
                  lastName:this.lastName,
                  phoneNo:this.phone

              })
              .then( (response) =>{
                  loader.hide()
                  console.log(response);
                  this.$emit('close-signup')
                  this.$notify({
                  group: "done",
                  title: "Done ",
                  text: "Signup completed!"
                  }, 4000);
              })
              .catch(function (error) {
                  loader.hide()
                  if (error.response && error.response.status==409) {
                      this.$notify({
                      group: "error",
                      title: "SignUp Failed",
                      text: "There is a account in this email."
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