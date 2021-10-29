import { shallowMount } from '@vue/test-utils'
import { createStore } from "vuex"
import AppHeader from '@/components/AppHeader.vue'

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


describe('AppHeader.vue', () => {
    it('renders when passed', () => {

      const wrapper = shallowMount(AppHeader,{
        global: {
          plugins: [store]
        }
      })
      expect(wrapper.exists()).toBe(true)
    })

    test('Click login', async () => {
        const wrapper = shallowMount(AppHeader, {
            global: {
              plugins: [store]
            }
          })
      
        await wrapper.get('[data-test="login"]').trigger('click')
        
      
        expect(wrapper.emitted()).toHaveProperty('open-login')
    })


    test('Click Signup', async () => {
        const wrapper = shallowMount(AppHeader, {
            global: {
              plugins: [store]
            }
          })
      
        await wrapper.get('[data-test="signup"]').trigger('click')
        
      
        expect(wrapper.emitted()).toHaveProperty('open-signup')
    })


    test('Check Logged Header', async () => {
        const wrapper = shallowMount(AppHeader, {
            global: {
              plugins: [store]
            }
          })
      
        
      
        expect(wrapper.find('[data-test="sp-logged-header"]').exists()).toBe(false)
        expect(wrapper.find('[data-test="client-logged-header"]').exists()).toBe(false)
        expect(wrapper.find('[data-test="not-logged-header"]').exists()).toBe(true)
    })


})