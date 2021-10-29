import { shallowMount } from '@vue/test-utils'
import Alert from '@/components/Alert.vue'

describe('Alert.vue', () => {
    it('renders props  when passed', () => {
      const message = 'test-msg'
      const heading='test-heading'
      const redButtonLabel="test-redButtonLabel"
      const whiteButtonLabel="test-whiteButtonLabel"
      const redButtonFunc=()=>{}
      const whiteButtonFunc=()=>{}
      const wrapper = shallowMount(Alert, {
        props: { message,heading,redButtonFunc,redButtonLabel,whiteButtonFunc,whiteButtonLabel}
      })
      expect(wrapper.exists()).toBe(true)
    })
  })