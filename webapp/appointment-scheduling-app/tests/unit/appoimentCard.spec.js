import { shallowMount } from '@vue/test-utils'
import AppoimentCard from '@/components/AppoimentCard.vue'


describe('AppoimentCard.vue', () => {
    it('renders props   when passed', () => {
      const date = 'test-date'
      const slotList=[
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
      const selecteServiceID="test-service-id"
      const wrapper = shallowMount(AppoimentCard, {
        props: { date,slotList,selecteServiceID}
      })
      expect(wrapper.exists()).toBe(true)
      expect(wrapper.findAll('[data-test="appoiment-card"]')).toHaveLength(slotList.length)
    })

    
  })