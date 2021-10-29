import { shallowMount } from '@vue/test-utils'
import { createStore } from "vuex"
import FloatingButton from '@/components/FloatingButton.vue'

const store = createStore({
    state() {
      return {
        isLogged:false,
        isClient:false,
      }
    },
  
    getters: {

    }

})

describe('FloatingButton.vue', () => {
    it('renders when passed', () => {

      const wrapper = shallowMount(FloatingButton,{
        global: {
          plugins: [store]
        }
      })
      expect(wrapper.exists()).toBe(true)
      expect(wrapper.get('[data-test="active-btn"]').text()).toBe('SP')
    })


})