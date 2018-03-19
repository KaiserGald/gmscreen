<template>
  <div id="lander">
    <h3 class="main-body-header">Welcome</h3>
    <div class="main-body-text">
      <p>Please click <router-link class="no-decoration" to="/login">here</router-link> to login!</p>
      <p>If you don't have an account, click <router-link class="no-decoration" to="/register">here</router-link> to register a new account.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'lander'
}
</script>

<style>
</style>
