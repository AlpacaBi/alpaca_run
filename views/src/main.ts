import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import service from './utils/ajax';
import 'element-ui/lib/theme-chalk/index.css';
import {
  Button,
  Menu,
  MenuItem,
  MenuItemGroup,
} from 'element-ui';

Vue.component(Button.name, Button);
Vue.component(Menu.name, Menu);
Vue.component(MenuItem.name, MenuItem);
Vue.component(MenuItemGroup.name, MenuItemGroup);

// 封装好了axios，并接入Vue的原型链，从而在任何地方都能使用
Vue.prototype.$ajax = service;

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
