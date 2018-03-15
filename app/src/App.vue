<template>
  <div id="app">
    <TitleBar/>
    <router-view/>
  </div>
</template>

<script>
import TitleBar from '@/components/TitleBar'
export default {
  name: 'App',
  components: {
    TitleBar
  }
}
</script>

<style>
</style>
