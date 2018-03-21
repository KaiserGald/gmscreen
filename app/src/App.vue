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
.no-decoration:link {
  text-decoration: none;
  color: #008080;
}

.no-decoration:visited {
  text-decoration: none;
  color: #008075;
}

.no-decoration:hover {
  text-decoration: none;
  color: #008075;
}

.no-decoration:active {
  text-decoration: none;
  color: #008080;
}

.main-body-header {
  margin: 2% 6% 2% 6%;
}

.main-body-text {
  margin: 0% 6% 0% 6%;
}

.form-error {
  color: red;
}

html {
  background-color: #202222;
  color: white;
}

body {
  margin: 0px;
}

</style>
