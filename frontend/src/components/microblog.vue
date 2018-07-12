<template>
  <div class="microblog">
    <h1>{{ title }}</h1>

    <h3 v-if="error">ОШИБКА: {{error}}</h3>

    <table>
     <tr>
        <th>Ник</th>
        <th>Имя</th>
        <th>Фамилия</th>
        <th>E-mail</th>
        <th>GitHub</th>
      </tr>
      <tr v-for="contact in blog_list" v-bind:key="contact.email">
        <td>{{contact.nikname}}</td>
        <td>{{contact.name}}</td>
        <td>{{contact.surname}}</td>
        <td>{{contact.email}}</td>
        <td>{{contact.github}}</td>
        <td><button v-on:click="edit_blog(contact)">Редактировать блог</button> </td>
        <td><button v-on:click="remove_blog(contact)">Удалить блог</button> </td>
      </tr>
    </table>

    <h3 v-if = "edit_index == -1">Новый блог</h3>
    <form>
      <p>Ник:<input type="text" v-model="new_blog.nikname"></p>
      <p>Имя:<input type="text" v-model="new_blog.name"></p>
      <p>Фамилия:<input type="text" v-model="new_blog.surname"></p>
      <p>E-mail:<input type="text" v-model="new_blog.email"></p>
      <p>GitHub:<input type="text" v-model="new_blog.github"></p>
      <button v-if = "edit_index == -1" v-on:click="add_new_blog">Добавить блог</button>
      <button v-if = "edit_index > -1" v-on:click="end_of_edition">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
const axios = require('axios')
export default {
  name: 'microblog',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      error: '',
      blog_list: [],
      new_blog:
        {
          'nikname': '',
          'name': '',
          'surname': '',
          'email': '',
          'github': ''
        }
    }
  },
  mouned: function () {
    this.get_blog()
  },
  methods: {
    get_blog: function () {
      this.error = ''
      const url = '/api/microblog/profiles'
      axios.get(url).then(response => {
        this.blog_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_blog: function () {
      const url = '/api/microblog/profiles'
      axios.post(url, this.new_blog).then(response => {
        console.log(response)
        this.blog_list.push(this.new_blog)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_blog: function (item) {
      const url = '/api/microblog/profiles/' + this.blog_list.indexOf(item)
      axios.delete(url).then(response => {
        this.blog_list.splice(this.blog_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_blog: function (item) {
      this.edit_index = this.blog_list.indexOf(item)
      this.new_blog = this.blog_list[this.edit_index]
    },
    end_of_edition: function () {
      const url = '/api/microblog/profiles/' + this.edit_index
      axios.put(url, this.new_blog).then(response => {
        console.log(response)
        this.edit_index = -1
        this.new_blog = {
          'nikname': '',
          'name': '',
          'surname': '',
          'email': '',
          'github': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
</style>
