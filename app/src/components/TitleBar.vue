<template>
  <div id="titlebar">
    <div id="title">
      <h1><router-link class="no-decoration" to="/">Unlicht GM Screen</router-link></h1>
    </div>
    <div id="nav">
      <router-link class="no-decoration" to="/login">Login</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'titlebar'
}
</script>

<style>
#titlebar {
  background-color: black;
  padding: 1px;
  margin: 0px;
  padding-left: 5%;
  display: grid;
  grid-template-columns: auto auto;
}

#nav {
  align-self: end;
  padding-bottom: 10px;
}

</style>
