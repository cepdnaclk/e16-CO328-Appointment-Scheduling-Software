 <template>
 <BlueModal eventName="close-appoimentmodel" customClases="">
     <AppoimentCard v-for="(ele,index) in dayDataList" :key="index" :date="ele.date" :slotList="ele.slotList" :selecteServiceID="selecteServiceID"/>
</BlueModal>
 </template>
 
 <script>
 import BlueModal from '../components/BlueModal.vue'
 import AppoimentCard from '../components/AppoimentCard.vue'
 
 export default {
 components:{
    BlueModal,
    AppoimentCard
 },
  props:["dayDataList","selecteServiceID"]
  
 }
 </script>
 
 <style>
 
 </style>