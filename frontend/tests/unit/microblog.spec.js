import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import microblog from '@/components/microblog.vue'

describe('microblog.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(microblog, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(microblog, {})
    expect(wrapper.text()).to.include('Ник')
    expect(wrapper.text()).to.include('Имя')
    expect(wrapper.text()).to.include('Фамилия')
    expect(wrapper.text()).to.include('E-mail')
    expect(wrapper.text()).to.include('GitHub')
  })
})
