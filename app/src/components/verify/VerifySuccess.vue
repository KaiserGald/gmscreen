<template>
  <div id="verify-success" v-if="shouldRender">
    <h3 class="main-body-header">Email Verified!</h3>
  </div>
</template>

<script>
export default {
  name: 'verifySuccess',
  props: ['shouldRender']
}
</script>

<style>
</style>
