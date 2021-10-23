 <template>
  <div class="bg-gradient-to-r from-gray-400 to-gray-200 p-6 grid grid-cols-2 gap-4">
   <div v-for="(ele,index) in services" :key="index" class=" m-2" @dblclick="()=>deleteService(ele.serviceId)">
     <Card :serviceDiscription="ele.serviceDiscription" :serviceType="ele.serviceType" :serviceName="ele.serviceName" :serviceId="ele.serviceId" :checkNow="()=>getDetail(ele.serviceId)"/>
   </div>    
  </div>
  <Alert v-if="isDeleteAlertShowing" message="Do you want to delete this service ?" heading="Delete" redButtonLabel="Delete" whiteButtonLabel="Cancel" :redButtonFunc="()=>confirmDelete()" :whiteButtonFunc="()=>cancelDelete()"/>
  <Alert v-if="isExpiredAlertShowing" message="Your service expired,Do you want to Remove this service ?" heading="Remove" redButtonLabel="Remove" whiteButtonLabel="Cancel" :redButtonFunc="()=>confirmExpiredDelete()" :whiteButtonFunc="()=>cancelExpiedDelete()"/>
  <AppoimentModal v-if="appoimentModelOpen" :dayDataList="dayDataList" @close-appoimentmodel="closeAppoimentModel" :selecteServiceID="selecteServiceID"
  @refresh-appoiment-model="refreshModal"
  />
</template>
 
 <script>
   import {getDataFromApiPOST,getDataFromApiGET} from '../Functions/dataApi'
   import Card from '../components/Card.vue'
   import Alert from '../components/Alert.vue'
   import AppoimentModal from './AppoimentModal.vue'
   import {GET_ALL_CREATED_SERVICES_URL,GET_DETAILS_OF_SERVICE,DELETE_SERVICE} from '../constatns'

   export default {
  beforeMount(){
    getDataFromApiGET(this,GET_ALL_CREATED_SERVICES_URL,"","","",false,false).then(data=>{
      this.services=data
    }).catch(_err=>{
      this.services=[]
    })
  },
  computed:{
    isDeleteAlertShowing(){
      return this.$store.state.dashboardAlert
    },
    isExpiredAlertShowing(){
      return this.$store.state.dashboardAlertExpired
    },
    appoimentModelOpen(){
      return this.$store.state.appoimentModal
    },
  
  },
  components:{
    Card,
    Alert,
    AppoimentModal
  },
  data(){
    return{
      services:[],
      deletingID:"",
      selecteServiceID:"",
      dayDataList:[]
    }
  },
  methods:{
    deleteService(id){
      this.deletingID=id
      this.$store.commit('setDashBoardAlert')
    },
    confirmDelete(){
      this.$store.commit('removeDashBoardAlert')
      getDataFromApiPOST(this,DELETE_SERVICE,{serviceId:this.deletingID},"Done ","Service Deleted!"
      ,'/dash-board',true,true).then(data=>{
        this.$route.go()
      })      
    },

    cancelDelete(){
      this.$store.commit('removeDashBoardAlert')
    },


    getDetail(id){
     console.log(id)
     this.selecteServiceID=id
     getDataFromApiPOST(this,GET_DETAILS_OF_SERVICE,{serviceID:id},"","",'',false,false).then(data=>{
          if(data.length==0){
            this.deletingID=id
            this.$store.commit('setDashBoardExpiredAlert')
          }else{
            this.dayDataList=data
            this.$store.commit('setAppoimentModal')
          }
      }).catch(error=>{
        console.log(error)
      })
    },


    cancelExpiedDelete(){
      this.$store.commit('removeDashBoardExpiredAlert')
    },


    confirmExpiredDelete(){
      this.$store.commit('removeDashBoardExpiredAlert')
      getDataFromApiPOST(this,DELETE_SERVICE,{serviceId:this.deletingID},"Done ","Service Removed!"
      ,'/dash-board',true,true)      
    },

    closeAppoimentModel(){
      this.$store.commit('removeAppoimentModal')
    },
    refreshModal(){
      this.$store.commit('removeAppoimentModal')
      this.getDetail(this.selecteServiceID)
    }

  }
 }
 </script>
 
 <style>
 
 </style>