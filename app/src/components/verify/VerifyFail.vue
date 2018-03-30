<template>
  <div id="verify-fail" v-if="shouldRender">
    <h3 class="main-body=header">Verification Failed!</h3>
  </div>
</template>

<script>
export default {
  name: 'verifyFail',
  props: ['shouldRender']
}
</script>

<style>
</style>
