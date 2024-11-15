import { createRouter, createWebHistory } from 'vue-router'
import LoginLayout from '@/views/layouts/LoginLayout.vue'
import HomeLayout from '@/views/layouts/HomeLayout.vue'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import { useUserStore } from '@/stores/user.js'
import AccountLayout from '@/views/layouts/AccountLayout.vue'
import AccountView from '@/views/AccountView.vue'
import ChangePasswordLayout from '@/views/layouts/ChangePasswordLayout.vue'
import ChangePasswordView from '@/views/ChangePasswordView.vue'
import AccountFormView from '@/views/AccountFormView.vue'
import FormLayout from '@/views/layouts/FormLayout.vue'
import AuthorizedLayout from '@/views/layouts/AuthorizedLayout.vue'
import TabungakuLayout from '@/views/layouts/TabungakuLayout.vue'
import TabungankuView from '@/views/TabungankuView.vue'
import BillerLayout from '@/views/layouts/BillerLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      meta: { layout: LoginLayout },
      component: LoginView,
    },
    {
      path: '/',
      name: 'home',
      meta: { layout: HomeLayout },
      component: HomeView,
    },
    {
      path: '/account',
      name: 'account',
      meta: { layout: AccountLayout },
      component: AccountView,
    },
    {
      path: '/changepassword',
      name: 'changepassword',
      meta: { layout: ChangePasswordLayout },
      component: ChangePasswordView,
    },
    {
      path: '/account/create',
      name: 'CreateAccount',
      meta: { layout: FormLayout },
      component: AccountFormView,
    },
    {
      path: '/account/edit/:accountId',
      name: 'EditAccount',
      meta: { layout: FormLayout },
      component: AccountFormView,
    },
    {
      path: '/tabunganku',
      name: 'tabunganku',
      meta: { layout: TabungakuLayout },
      component: TabungankuView,
    },
    {
      path: '/authorized',
      name: 'AuthorizedView',
      meta: { layout: AuthorizedLayout },
      component: () => import('@/views/AuthorizedView.vue'),
    },
    {
      path: '/biller',
      name: 'BillerView',
      meta: { layout: BillerLayout },
      component: () => import('@/views/BillingView.vue'),
    },
  ],
})

// Global route guard to check for token
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()

  if (to.name !== 'login' && !userStore.token) {
    next({ name: 'login' })
  } else if (to.name == 'login' && userStore.token) {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router
