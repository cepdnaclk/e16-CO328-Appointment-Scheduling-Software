<template>
   <div class="container-bar">
        <div class="relative">
            <div class="input-tab"> <i class="fa fa-search text-gray-400 z-20 hover:text-gray-500"></i> </div> <input type="text" class="h-14 w-96 pl-10 pr-20 rounded-lg z-0 focus:shadow focus:outline-none" placeholder="Search anything...">
            <div class="searchbutton"> <button class="button-s" type='submit'>Search</button> </div>
        </div>
    </div>
    <div class="p-10 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-3 lg:grid-cols-3 xl:grid-cols-3 gap-5">
    <card serviceDiscription="bbbbbbbb" serviceId="000" serviceName="aseddddmm" serviceType="Monthly"/>
    <card />
    <card />
    <card />
    <card />
    <card />
    </div>
  
</template>

 <script>
   import card from '../components/CardClient.vue'
   import {getDataFromApiPOST} from '../Functions/dataApi'
   import AppoimentModal from './ClientAppointmentModal.vue'

   export default {
  computed:{
    isDeleteAlertShowing(){
      return this.$store.state.dashboardAlert
    },
    isExpiredAlertShowing(){
      return this.$store.state.dashboardAlertExpired
    },
    appoimentModelOpen(){
      return this.$store.state.appoimentModal
    }
  },
  components:{
    card,
    AppoimentModal
  },
  data(){
    return{
      services:[
        {
         "serviceName":"Topic Of service provider",
         "serviceDiscription":"asdddddddddddddddddddddddddddddddddd",
         "serviceType":"service category",
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
      deletingID:"",
      selecteServiceID:"",
      dayDataList:[
        {
          date:"2021-11-11",
          slotList:[
            {
              slotId:1,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:false,
              approved:false,
            },
            {
              slotId:2,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:true,
              approved:false,
            },
            {
              slotId:3,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:true,
              approved:true,
            }
          ]
        },
        {
          date:"2021-11-11",
          slotList:[
            {
              slotId:1,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:false,
              approved:false,
            },
            {
              slotId:2,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:true,
              approved:false,
            },
            {
              slotId:3,
              time:"11:00:00-12:00:00 am",
              clientName:"Supun",
              clientEmail:"Supun@email",
              clientRequested:true,
              approved:true,
            }
          ]
        }
      ]
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
      console.log(id)
      this.$store.commit('setAppoimentModal')
      this.selecteServiceID=id
      /*
      let data = getDataFromApiPOST(this,'getDetailService',{serviceID:id},"","",'',false,false)
      if(data.length==0){
        this.deletingID=id
        this.$store.commit('setDashBoardExpiredAlert')
        return
      }else{
        this.dayDataList=data
        this.$store.commit('setAppoimentModal')
      }
      */
      
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
@import url('../assets/findservice.css');
</style>