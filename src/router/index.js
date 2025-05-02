/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import { setupLayouts } from 'virtual:generated-layouts'
import { routes } from 'vue-router/auto-routes'
import { authenticated } from '@/api'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: setupLayouts(routes),
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (!localStorage.getItem('vuetify:dynamic-reload')) {
      console.debug('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    } else {
      console.error('Dynamic import error, reloading page did not fix it', err)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

router.beforeEach(async (to, from, next) => {
  const isAuthenticated = await authenticated(localStorage.getItem('token'))

  if (to.path === '/login' && !isAuthenticated) {
    console.debug('User is not authenticated, redirecting to login')
    next();
  } else if (!isAuthenticated) {
    console.debug('User is not authenticated, redirecting to login')
    next('/login')
  } else {
    if (to.path === '/login' && isAuthenticated) {
      console.debug('User is authenticated, redirecting to home')
      next('/')
      return
    }

    console.debug('User is authenticated')

    next()
  }
})

export default router
