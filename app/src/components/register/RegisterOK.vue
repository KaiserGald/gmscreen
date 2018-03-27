<template>
  <div id="register-ok" v-if="shouldRender">
    <h3 class="main-body-header">Thank you for registering!</h3>
    <p>Click <router-link class="no-decoration" to="/">here</router-link> to return to the main page.</p>
  </div>
</template>

<script>
export default {
  name: 'registerOK',
  props: ['shouldRender']
}
</script>

<style>
</style>
