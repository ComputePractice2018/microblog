<template>
  <div class="microblog">
    <h1>{{ title }}</h1>

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
      <button v-if = "edit_index > -1" v-on:click="make_new_blog">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'microblog',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      blog_list: [
        {
          'nikname': 'hrilev',
          'name': 'Максим',
          'surname': 'Хрылев',
          'email': 'hrilev@dyandex.ru',
          'github': 'hrilev'
        },
        {
          'nikname': 'Ник',
          'name': 'Имя',
          'surname': 'Фамилия',
          'email': 'user@domain.ru',
          'github': 'userr'
        }
      ],
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
  methods: {
    add_new_blog: function () {
      this.blog_list.push(this.new_blog)
    },
    remove_blog: function (item) {
      this.blog_list.splice(this.blog_list.indexOf(item), 1)
    },
    edit_blog: function (item) {
      this.edit_index = this.blog_list.indexOf(item)
      this.new_blog = this.blog_list[this.edit_index]
    },
    make_new_blog: function () {
      this.edit_index = -1
      this.new_blog = {
        'nikname': '',
        'name': '',
        'surname': '',
        'email': '',
        'github': ''
      }
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
</style>
