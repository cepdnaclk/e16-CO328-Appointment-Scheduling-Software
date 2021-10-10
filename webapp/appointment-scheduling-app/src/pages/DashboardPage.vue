 <template>
  <div class="bg-gradient-to-r from-gray-400 to-gray-200 p-6 grid grid-cols-2 gap-4">
   <div v-for="(ele,index) in services" :key="index" class=" m-2" @dblclick="deleteService(ele.serviceId)">
     <Card :serviceDiscription="ele.serviceDiscription" :serviceType="ele.serviceType" :serviceName="ele.serviceName" :serviceId="ele.serviceId"/>
   </div>    
  </div>
  <Alert v-if="isAlertShowing" message="Do you want to delete this service ?" heading="Delete" redButtonLabel="Delete" whiteButtonLabel="Cancel" :redButtonFunc="confirmDelete" :whiteButtonFunc="cancelDelete"/>
 </template>
 
 <script>
   //import {getDataFromApiGET} from '../Functions/dataApi'
   import axios from 'axios'
   import Card from '../components/Card.vue'
   import Alert from '../components/Alert.vue'
 export default {
  computed:{
    isAlertShowing(){
      return this.$store.state.dashboardAlert
    }
  },
  components:{
    Card,
    Alert
  },
  data(){
    return{
      services:[
        {
         "serviceName":"Topic Of Name",
         "serviceDiscription":"asdddddddddddddddddddddddddddddddddd",
         "serviceType":"Monthly",
         "serviceId":'0000000',
        },
        {
         "serviceName":"Topic Of Name",
         "serviceDiscription":"asdddddddddddddddddddddddddddddddddd",
         "serviceType":"Monthly",
         "serviceId":'0000000',
        },
        {
         "serviceName":"Topic Of Name",
         "serviceDiscription":"asdddddddddddddddddddddddddddddddddd",
         "serviceType":"Monthly",
         "serviceId":'0000000',
        }
      ],
      deletingID:""
    }
  },
  methods:{
    deleteService(id){
      this.deletingID=id
      this.$store.commit('setDashBoardAlert')
    },
    confirmDelete(){
      this.$store.commit('removeDashBoardAlert')
      
      let loader = this.$loading.show();
      axios.post('deleteService', {
         deleteID:this.deletingID
          },{
          
          headers: {
            'Authorization': this.$cookies.get("user_token")
          }
      }).then( (response) =>{
          loader.hide()
          this.$notify({
          group: "done",
          title: "Done ",
          text: "Service Deleted!"
          }, 4000);
        this.$router.push('/dash-board')
        console.log(response);
      }).catch(function (error) {
         loader.hide()
          
          if (error.response && error.response.status==403) {
              this.$cookies.remove("user_token")
              this.$cookies.remove("user_type")
              this.$notify({
                group: "generic",
                title: "Session Expires ",
                text: "Session Expired pleace login"
                }, 4000)
                this.$store.commit('setLogedOut')
                this.$router.push('/')
              } else {
                this.$notify({
                  group: "error",
                  title: "Error",
                  text: "Something happened in setting up the request that triggered a Error!",
                }, 4000)
            }
               
         });
          
    },
    cancelDelete(){
      this.$store.commit('removeDashBoardAlert')
    }
  }
 }
 </script>
 
 <style>
 
 </style>