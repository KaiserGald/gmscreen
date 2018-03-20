<template>
  <div id="register">
    <h3 class="main-body-header">Register</h3>
    <form class="main-body-text" @submit.prevent="postRegister">
      Username:
      <input type="text" v-model="username"><br>
      Email:
      <input type="text" v-model="email"><br>
      Password:
      <input type="text" v-model="password"><br>
      <button type="submit">Submit</button>
      </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'register',
  data () {
    return {
      username: '',
      email: '',
      password: ''
    }
  },
  methods: {
    postRegister () {
      var params = new URLSearchParams()
      params.append('username', this.username)
      params.append('email', this.email)
      params.append('password', this.password)

      axios({
        url: 'gmscreen/register',
        method: 'GET',
        params: params
      }).then(function (response) {
        if (response.status === 200) {
          console.log('test')
          axios({
            url: 'gmscreen/register',
            method: 'POST',
            params: params
          })
        }
      }).catch(function (error) {
        if (error.response) {
          console.log(error.response.data)
          console.log(error.response.status)
          console.log(error.response.headers)
        }
      })
    }
  }
}
</script>

<style>
</style>
