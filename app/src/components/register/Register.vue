<template>
  <div id="register">
    <RegisterForm v-on:regOk="regOk" :should-render="showRegForm"/>
    <RegisterOK :should-render="showThanksForm"/>
  </div>
</template>

<script>
import RegisterForm from '@/components/register/RegisterForm.vue'
import RegisterOK from '@/components/register/RegisterOK.vue'

export default {
  name: 'register',
  data () {
    return {
      showRegForm: true,
      showThanksForm: false
    }
  },
  components: {
    RegisterForm,
    RegisterOK
  },
  methods: {
    regOk () {
      this.showRegForm = !this.showRegForm
      this.showThanksForm = !this.showThanksForm
    }
  }
}
</script>

<style>
</style>
