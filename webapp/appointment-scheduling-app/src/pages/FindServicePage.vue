<template>
   <div class="container-bar">
        <div class="relative">
            <div class="input-tab"> <i class="fa fa-search text-gray-400 z-20 hover:text-gray-500"></i> </div> <input type="text" class="h-14 w-96 pl-10 pr-20 rounded-lg z-0 focus:shadow focus:outline-none" placeholder="Search anything..." v-bind="searchKey">
            <div class="searchbutton"> <button class="button-s" @click="search" >Search</button> </div>
        </div>
    </div>
    <div class="p-10 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-3 lg:grid-cols-3 xl:grid-cols-3 gap-5">
    <card v-for="ele,index in services" :key="index" :serviceDiscription="ele.serviceDiscription" :serviceName="ele.serviceName" :serviceId="ele.serviceId" :serviceType="ele.serviceType"
    :fname="ele.ownerFristName" :lname="ele.ownerLastName" :ownerEmail="ele.serviceOwnerEmail"
    :checkNow="()=>checkNow(ele.serviceOwnerEmail,ele.serviceId)"
    />
    </div>
  <Alert v-if="isFailedRequestAlert" message="Someone has requested slot earlier" heading="Already Requested" redButtonLabel="Refresh" whiteButtonLabel="Cancel" :redButtonFunc="()=>refresh()" :whiteButtonFunc="()=>cancelAlert()"/>
  <SearchAppoimentModal v-if="isOpenedSearchAppoimentModal"
   @done="requestDone"
   @request-failed-alert="alertOpen"
   @close-search-appoimentmodel="closeSearchAppoimentModal" 
  :dayDataList="getDayDataList" 
  :selecteServiceID="selecteServiceID"
  :ownerEmail="ownerEmail"
  />
</template>

 <script>
   import card from '../components/CardClient.vue'
   import Alert from '../components/Alert.vue'
   import {getDataFromApiPOST,getDataFromApiGET} from '../Functions/dataApi'
   import SearchAppoimentModal from './SearchAppoimentModal.vue'
   import {SEARCH_URL,GET_NEWEST_SERVICES,GET_CLIENT_SLOTS_OF_A_SERVICE} from '../constatns'
  export default {
    components:{
      card,
      SearchAppoimentModal,
      Alert
    },
    computed:{
      isOpenedSearchAppoimentModal(){
        return this.openappoimentmodal
      },
      isFailedRequestAlert(){
        return this.isRequestfailed
      },
      getDayDataList(){
        return this.dayDataList
      }
    },
    data(){
      return{
        services:[],
        searchKey:"",
        selecteServiceID:"",
        ownerEmail:"",
        dayDataList:[],
        openappoimentmodal:false,
        isRequestfailed:false
      }
    },
    beforeMount(){
      getDataFromApiGET(this,GET_NEWEST_SERVICES,"","","",false,false).then(data=>{
      this.services=data
      }).catch(_err=>{
        this.services=[] 
      })
    },
    methods:{

      checkNow(ownerEmail,serviceId){
        this.ownerEmail=ownerEmail
        this.selecteServiceID=serviceId
        this.openappoimentmodal=true
        getDataFromApiPOST(this,GET_CLIENT_SLOTS_OF_A_SERVICE,{serviceId,ownerEmail},"","",'',false,false).then(data=>{
          this.dayDataList=data
        }).catch(err=>{
          console.log(err)
          this.dayDataList=[]
        })
        
      },

      search(){
        getDataFromApiPOST(this,SEARCH_URL,{serviceName:this.searchKey},"","",'',false,false).then(data=>{
          this.services=data
        }).catch(_err=>{
          this.services=[]
        })
      },
      closeSearchAppoimentModal(){
        this.openappoimentmodal=false
      },
      cancelAlert(){
        this.isRequestfailed=false
      },alertOpen(){
        this.openappoimentmodal=false
        this.isRequestfailed=true
      },
      refresh(){
        this.isRequestfailed=false
        this.checkNow(this.ownerEmail,this.selecteServiceID)
      },
      requestDone(){
        this.checkNow(this.ownerEmail,this.selecteServiceID)
      }
    }
  }
 </script>

<style>
@import url('../assets/findservice.css');
</style>