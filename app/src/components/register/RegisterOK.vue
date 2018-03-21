<template>
  <div id="register-ok" v-if="shouldRender">
    <h3 class="main-body-header">Thank you for registering!</h3>
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
