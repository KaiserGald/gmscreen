<template>
  <div id="register">
    <h3 class="main-body-header">Register</h3>
    <form class="main-body-text" @submit.prevent="postRegister">
      <div class="form-error" v-show="usernameExists == true">Username already exists!</div>
      Username:
      <input type="text" v-model="username"><br>
      <div class="form-error" v-show="emailExists == true">Email already exists!</div>
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
      password: '',
      usernameExists: false,
      emailExists: false
    }
  },
  methods: {
    postRegister () {
      var params = new URLSearchParams()
      params.append('username', this.username)
      params.append('email', this.email)
      params.append('password', this.password)

      var self = this
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
        if (error.response.status === 400) {
          console.log(error.response.data.Username)
          self.usernameExists = error.response.data.Username
          self.emailExists = error.response.data.Email
        }
      })
    }
  }
}
</script>

<style>
</style>
