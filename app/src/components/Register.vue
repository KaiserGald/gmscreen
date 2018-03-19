<template>
  <div id="register">
    <h3 class="main-body-header">Register</h3>
    <form class="main-body-text" @submit.prevent="postRegister">
      Username:
      <input type="text" v-model="username"><br>
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
      password: ''
    }
  },
  methods: {
    postRegister () {
      var params = new URLSearchParams()
      params.append('username', this.username)
      params.append('password', this.password)
      axios({
        url: 'gmscreen/register',
        method: 'POST',
        params: params
      })
    }
  }
}
</script>

<style>
</style>
