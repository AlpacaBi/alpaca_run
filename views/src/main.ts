import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import services from './utils/ajax';
import 'element-ui/lib/theme-chalk/index.css';
import {
  Button,
  Menu,
  MenuItem,
  MenuItemGroup,
  Carousel,
  CarouselItem,
  Tooltip,
  Notification,
  Dialog,
} from 'element-ui';

Vue.component(Button.name, Button);
Vue.component(Menu.name, Menu);
Vue.component(MenuItem.name, MenuItem);
Vue.component(MenuItemGroup.name, MenuItemGroup);
Vue.component(Carousel.name, Carousel);
Vue.component(CarouselItem.name, CarouselItem);
Vue.component(Tooltip.name, Tooltip);
Vue.component(Notification.name, Notification);
Vue.component(Dialog.name, Dialog);

// 封装好了axios，并接入Vue的原型链，从而在任何地方都能使用
Vue.prototype.$ajax = services.service;
Vue.prototype.$github_ajax = services.githubService;

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
