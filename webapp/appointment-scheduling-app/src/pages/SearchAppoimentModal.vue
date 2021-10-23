<template>
 <BlueModal eventName="close-search-appoimentmodel" customClases="">
     <SearchAppoimentCard v-for="(ele,index) in dayDataList" :key="index" :date="ele.date" :slotList="ele.slotList" :selecteServiceID="selecteServiceID" :ownerEmail="ownerEmail"
     @request-failed-alert="$emit('request-failed-alert')"
     @done="$emit('done')"
     />
</BlueModal>
</template>
 
 <script>
 import BlueModal from '../components/BlueModal.vue'
 import SearchAppoimentCard from '../components/SearchAppoimentCard.vue'
 
 export default {
 components:{
    BlueModal,
    SearchAppoimentCard
 },
  props:["dayDataList","selecteServiceID","ownerEmail"]
  
 }
 </script>
 
 <style>
 
 </style>