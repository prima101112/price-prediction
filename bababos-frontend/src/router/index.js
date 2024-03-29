import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import HistopricalpoView from '../views/HistoricalpoView.vue'
import CustomersView from '../views/CustomersView.vue'
import SuggestionView from '../views/SuggestionView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/customers',
      name: 'customers',
      component: CustomersView
    },
    {
      path: '/historicalpo',
      name: 'historicalpo',
      component: HistopricalpoView
    },
    {
      path: '/suggestion',
      name: 'suggestion',
      component: SuggestionView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
