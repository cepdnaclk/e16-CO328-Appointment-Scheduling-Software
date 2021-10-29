import { shallowMount } from '@vue/test-utils'
import Card from '@/components/Card.vue'

describe('Card.vue', () => {
    it('renders props  when passed', () => {
      const serviceName = 'test-serviceName'
      const serviceDiscription='test-serviceDiscription'
      const serviceType="test-serviceType"
      const serviceId="test-serviceId"
      const checkNow=()=>{}

      const wrapper = shallowMount(Card, {
        props: { serviceDiscription,serviceId,serviceName,serviceType,checkNow}
      })
      expect(wrapper.exists()).toBe(true)
    })
  })