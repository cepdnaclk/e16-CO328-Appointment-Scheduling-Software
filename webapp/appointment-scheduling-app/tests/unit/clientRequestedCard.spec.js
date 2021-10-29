import { shallowMount } from '@vue/test-utils'
import ClientRequestedCard from '@/components/ClientRequestedCard.vue'

describe('ClientRequestedCard.vue', () => {
    it('renders props approved=false when passed', () => {
      const serviceName = 'test-serviceName'
      const time='test-time'
      const servicOwnerEmail="test-servicOwnerEmail"
      const date="test-date"
      const approved=false
      const onPressCancel=()=>{}

      const wrapper = shallowMount(ClientRequestedCard, {
        props: { servicOwnerEmail,serviceName,time,date,approved,onPressCancel}
      })
      expect(wrapper.exists()).toBe(true)
      expect(wrapper.get('[data-test="status"]').text()).toBe('Pending')
    })

    it('renders props approved=true when passed', () => {
        const serviceName = 'test-serviceName'
        const time='test-time'
        const servicOwnerEmail="test-servicOwnerEmail"
        const date="test-date"
        const approved=true
        const onPressCancel=()=>{}
  
        const wrapper = shallowMount(ClientRequestedCard, {
          props: { servicOwnerEmail,serviceName,time,date,approved,onPressCancel}
        })
        expect(wrapper.exists()).toBe(true)
        expect(wrapper.get('[data-test="status"]').text()).toBe('Approved')
      })
  })