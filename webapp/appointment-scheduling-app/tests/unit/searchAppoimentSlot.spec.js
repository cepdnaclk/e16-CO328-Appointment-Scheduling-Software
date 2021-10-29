import { shallowMount } from '@vue/test-utils'
import SearchAppoimentSlot from '@/components/SearchAppoimentSlot.vue'


describe('SearchAppoimentSlot.vue', () => {
    it('renders props clientRequested true  when passed', () => {
      const time = 'test-time'
      const clientRequested=true
      const requestFunc=()=>{}
      const wrapper = shallowMount(SearchAppoimentSlot, {
        props: { time,clientRequested,requestFunc}
      })
      expect(wrapper.exists()).toBe(true)
      expect(wrapper.get('[data-test="status"]').text()).toBe('Already Requested..')
      expect(wrapper.get('[data-test="btn-text"]').text()).toBe('Not Available')
    })

    it('renders props clientRequested false  when passed', () => {
        const time = 'test-time'
        const clientRequested=false
        const requestFunc=()=>{}
        const wrapper = shallowMount(SearchAppoimentSlot, {
          props: { time,clientRequested,requestFunc}
        })
        expect(wrapper.exists()).toBe(true)
        expect(wrapper.get('[data-test="status"]').text()).toBe('Not Requested Yet')
        expect(wrapper.get('[data-test="btn-text"]').text()).toBe('Request')
      })
  })