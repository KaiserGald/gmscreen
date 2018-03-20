c<template>
  <div id="login">
    <h3 class="main-body-header">Login</h3>
    <form class="main-body-text" @submit.prevent="postLogin">
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
  name: 'login',
  data () {
    return {
      username: '',
      password: ''
    }
  },

  methods: {
    postLogin () {
      var params = new URLSearchParams()
      params.append('username', this.username)
      params.append('password', this.password)
      axios.post('gmscreen/login', params)
    }
  }
}
</script>

<style>
</style>
