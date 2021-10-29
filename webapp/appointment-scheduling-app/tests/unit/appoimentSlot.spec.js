import { shallowMount } from '@vue/test-utils'
import AppoimentSlot from '@/components/AppoimentSlot.vue'


describe('AppoimentSlot.vue', () => {
    it('renders props clientRequested true approved=true when passed', () => {
      const clientName = 'test-clientName'
      const clientEmail="test-clientEmail"
      const clientRequested=true
      const approved=true

      const wrapper = shallowMount(AppoimentSlot, {
        props: { clientName,clientEmail,clientRequested,approved}
      })
      expect(wrapper.exists()).toBe(true)
      expect(wrapper.get('[data-test="clientName"]').text()).toBe(clientName)
      expect(wrapper.get('[data-test="clientEmail"]').text()).toBe(clientEmail)
      expect(wrapper.get('[data-test="status"]').text()).toBe('Approved')
      expect(wrapper.get('[data-test="approveBtn"]').text()).toBe('Approved')

      let classes;
      classes=["p-2", "hover:shadow-lg" ,"text-xs", "font-thin", "m-1" , "outline-none","bg-gray-500" ,"text-black"]
      classes.forEach(ele=>{
        expect(wrapper.get('[data-test="approveBtn"]').classes()).toContain(ele)
      })
      
      classes=["p-2" , "hover:shadow-lg" ,"text-xs" ,"font-thin" ,"m-1" , "outline-none", "bg-red-500", "text-white", "hover:bg-white" ,"hover:text-red-500"]

      classes.forEach(ele=>{
        expect(wrapper.get('[data-test="removeBtn"]').classes()).toContain(ele)
      })
     
    })

    it('renders props clientRequested true approved=false  when passed', () => {
        const clientName = 'test-clientName'
        const clientEmail="test-clientEmail"
        const clientRequested=true
        const approved=false
  
        const wrapper = shallowMount(AppoimentSlot, {
            props: { clientName,clientEmail,clientRequested,approved}
        })
        expect(wrapper.exists()).toBe(true)
        expect(wrapper.get('[data-test="clientName"]').text()).toBe(clientName)
        expect(wrapper.get('[data-test="clientEmail"]').text()).toBe(clientEmail)
        expect(wrapper.get('[data-test="status"]').text()).toBe('Pending..')
        expect(wrapper.get('[data-test="approveBtn"]').text()).toBe('Approve')
        let classes;
        classes=["p-2", "hover:shadow-lg", "text-xs", "font-thin", "m-1" , "bg-blue-500" , "text-white", "hover:bg-white" ,"hover:text-blue-500"]
        classes.forEach(ele=>{
            expect(wrapper.get('[data-test="approveBtn"]').classes()).toContain(ele)
        })
        classes=["p-2",  "hover:shadow-lg", "text-xs", "font-thin", "m-1" , "outline-none", "bg-red-500", "text-white" ,"hover:bg-white" ,"hover:text-red-500"]
        classes.forEach(ele=>{
            expect(wrapper.get('[data-test="removeBtn"]').classes()).toContain(ele)
        })
        
    })

    it('renders props clientRequested false approved=false when passed', () => {
        const clientName = 'test-clientName'
        const clientEmail="test-clientEmail"
        const clientRequested=false
        const approved=false
  
        const wrapper = shallowMount(AppoimentSlot, {
            props: { clientName,clientEmail,clientRequested,approved}
        })
        expect(wrapper.exists()).toBe(true)
        expect(wrapper.get('[data-test="clientName"]').text()).toBe('')
        expect(wrapper.get('[data-test="clientEmail"]').text()).toBe('')
        expect(wrapper.get('[data-test="status"]').text()).toBe('Not Requested Yet')
        expect(wrapper.get('[data-test="approveBtn"]').text()).toBe('Approve')
        let classes;
        classes=["p-2", "hover:shadow-lg", "text-xs", "font-thin" ,"m-1" , "outline-none", "bg-gray-500" ,"text-black"]
        classes.forEach(ele=>{
            expect(wrapper.get('[data-test="approveBtn"]').classes()).toContain(ele)
        })
        classes=["p-2" , "hover:shadow-lg" ,"text-xs", "font-thin", "m-1" , "outline-none" ,"bg-gray-500" ,"text-black"]
        classes.forEach(ele=>{
            expect(wrapper.get('[data-test="removeBtn"]').classes()).toContain(ele)
        })
        
    })

    
  })