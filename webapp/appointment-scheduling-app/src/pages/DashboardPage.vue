 <template>
  <div class="bg-gradient-to-r from-gray-400 to-gray-200 p-6 grid grid-cols-2 gap-4">
   <div v-for="(ele,index) in services" :key="index" class=" m-2" @dblclick="()=>deleteService(ele.serviceId)">
     <Card :serviceDiscription="ele.serviceDiscription" :serviceType="ele.serviceType" :serviceName="ele.serviceName" :serviceId="ele.serviceId" :checkNow="()=>getDetail(ele.serviceId)"/>
   </div>    
  </div>
  <Alert v-if="isDeleteAlertShowing" message="Do you want to delete this service ?" heading="Delete" redButtonLabel="Delete" whiteButtonLabel="Cancel" :redButtonFunc="()=>confirmDelete()" :whiteButtonFunc="()=>cancelDelete()"/>
  <Alert v-if="isExpiredAlertShowing" message="Your service expired,Do you want to Remove this service ?" heading="Remove" redButtonLabel="Remove" whiteButtonLabel="Cancel" :redButtonFunc="()=>confirmExpiredDelete()" :whiteButtonFunc="()=>cancelExpiedDelete()"/>
  <AppoimentModal  @close-appoimentmodel="closeAppoimentModel" />
 </template>
 
 <script>
   import {getDataFromApiPOST} from '../Functions/dataApi'
   import Card from '../components/Card.vue'
   import Alert from '../components/Alert.vue'
   import AppoimentModal from './AppoimentModal.vue'

   export default {
  computed:{
    isDeleteAlertShowing(){
      return this.$store.state.dashboardAlert
    },
    isExpiredAlertShowing(){
      return this.$store.state.dashboardAlertExpired
    }
  },
  components:{
    Card,
    Alert,
    AppoimentModal
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
      getDataFromApiPOST(this,'deleteService',{deleteID:this.deletingID},"Done ","Service Deleted!"
      ,'/dash-board',true,true)      
    },

    cancelDelete(){
      this.$store.commit('removeDashBoardAlert')
    },


    getDetail(id){
      let data=getDataFromApiPOST(this,'getDetailService',{serviceID:id},"","",'',false,false)
      if(data.length==0){
        this.deletingID=id
        this.$store.commit('setDashBoardExpiredAlert')
        return
      }else{

        this.$store.commit('setAppoimentModal')
      }
      
    },


    cancelExpiedDelete(){
      this.$store.commit('removeDashBoardExpiredAlert')
    },


    confirmExpiredDelete(){
      this.$store.commit('removeDashBoardExpiredAlert')
      getDataFromApiPOST(this,'deleteService',{deleteID:this.deletingID},"Done ","Service Removed!"
      ,'/dash-board',true,true)      
    },

    closeAppoimentModel(){
      this.$store.commit('removeAppoimentModal')
    }

  }
 }
 </script>
 
 <style>
 
 </style>