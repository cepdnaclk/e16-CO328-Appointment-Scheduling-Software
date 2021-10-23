<template>
<ClientRequestedCard v-for="(ele,index) in services" :key="index" :approved="ele.approved" :discription="ele.serviceDiscription" :serviceName="ele.serviceName" :time="ele.time" :onPressCancel="()=>cancelPress(ele.serviceId,ele.serviceOwnerEmail,ele.slotId,ele.date)" :servicOwnerEmail="ele.servicOwnerEmail" :date="ele.date" />
</template>

<script>
import ClientRequestedCard from '../components/ClientRequestedCard.vue'
import {GET_ALL_CLIENT_REQUESTED_SERVICES,CANSEL_REQUESTED_SERVICE} from '../constatns'
 import {getDataFromApiGET,getDataFromApiPOST} from '../Functions/dataApi'
export default {
beforeMount(){
    getDataFromApiGET(this,GET_ALL_CLIENT_REQUESTED_SERVICES,"","","",false,false).then(data=>{
      this.services=data
      console.log(data)
    }).catch(_err=>{
      this.services=[]
    })
},
components:{
 ClientRequestedCard
},
data(){
    return{
        services:[]
    }
},
methods:{
    cancelPress(serviceId,serviceOwnerEmail,slotId,date){
        getDataFromApiPOST(this,CANSEL_REQUESTED_SERVICE,{serviceId,serviceOwnerEmail,slotId,date},"Done","Canceled request",'',false,true).then(data=>{
            this.$router.go()
        }).catch(err=>{
            console.log(err)
        })
    }
}
}
</script>

<style>

</style>