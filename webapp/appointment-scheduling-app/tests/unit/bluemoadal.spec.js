import { shallowMount,mount } from '@vue/test-utils'
import BlueModal from '@/components/BlueModal.vue'

describe('BlueModal.vue', () => {
  it('renders props.eventName and props.customClases when passed', () => {
    const eventName = 'test-event'
    const customClases='m-auto'
    const wrapper = shallowMount(BlueModal, {
      props: { eventName,customClases }
    })
    expect(wrapper.exists()).toBe(true)
  })

  test('Emits an event when clicked background of BlueModal', () => {
    const wrapper = mount(BlueModal,{
        props: {
            eventName : 'test-event',
            customClases:'m-auto'
        }})
  
    wrapper.find('section').trigger('click')
  
    expect(wrapper.emitted()).toHaveProperty('test-event')
    })


    test('Test BlueModal slot contain', () => {
        const wrapper = mount(BlueModal, {
        slots: {
            default: '<div>Test Content</div>'
        }
        })
    
        expect(wrapper.html()).toContain('<div>Test Content</div>')
    })

})



  
  