<template>
  <div id="verify">
    <VerifySuccess v-on:success="verifyOk" :should-render="showSuccess" />
    <VerifyFail v-on:fail="verifyFail" :should-render="showFail" />
    <div class="form-error" v-show="validToken == false">Submitted token is invalid.
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import VerifyFail from '@/components/verify/VerifyFail.vue'
import VerifySuccess from '@/components/verify/VerifySuccess.vue'

export default {
  name: 'verify',
  data () {
    return {
      showSuccess: false,
      showFail: false,
      validToken: true,
      expiredToken: false,
      emailVerified: false
    }
  },
  components: {
    VerifyFail,
    VerifySuccess
  },
  mounted: function () {
    var urlString = window.location.href
    var url = new URL(urlString)
    var self = this
    axios({
      url: 'gmscreen/verify',
      method: 'POST',
      params: url.searchParams
    }).then(function (response) {
      console.log(response)
      if (response.status === 200) {
        console.log('success')
        self.$emit('success')
        self.showSuccess = true
      }
    }).catch(function (error) {
      console.log(error)
      if (error.response.status === 400) {
        self.emailVerified = error.response.data.EmailVerified
        self.validToken = error.response.data.TokenValid
        if (self.validToken) {
          self.expiredToken = error.response.data.TokenExpired
        }
        if (self.emailVerified) {
          console.log('here')
          self.$router.push('/resend')
        }
        self.showFail = true
      }
    })
  },
  methods: {
    verifyOk () {
      this.showSuccess = true
      this.showFail = false
    },
    verifyFail () {
      this.showSuccess = false
      this.showFail = true
    }
  }
}
</script>

<style>
</style>
