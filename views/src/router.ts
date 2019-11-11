import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/Home.vue';

Vue.use(Router);

export default new Router({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
      children: [
        {
          path: '/',
          name: 'info',
          component: () => import( './components/AboutSection/Info.vue'),
        },
        {
          path: 'skills',
          name: 'skills',
          component: () => import('./components/AboutSection/Skills.vue'),
        },
        {
          path: 'profile',
          name: 'profile',
          component: () => import( './components/AboutSection/Profile.vue'),
        },
        {
          path: 'contact',
          name: 'contact',
          component: () => import('./components/AboutSection/Contact.vue'),
        },
      ],
    },
  ],
});
