import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Lander from '@/components/Lander'
import Register from '@/components/register/Register'
import Verify from '@/components/verify/Verify'
import Resend from '@/components/Resend'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Lander',
      component: Lander
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/verify',
      name: 'Verify',
      component: Verify
    },
    {
      path: '/resend',
      name: 'Resend',
      component: Resend
    }
  ]
})
