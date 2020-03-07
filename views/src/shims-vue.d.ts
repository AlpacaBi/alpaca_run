import Vue from 'vue';
import qs from 'qs';
import VueRouter, { Route } from "vue-router";
declare module '*.vue' {
  export default Vue;
}

declare module 'qs'

declare module "vue/types/vue" {
  interface Vue {
    $router: VueRouter; // 这表示this下有这个东西
    $route: Route;
    $ajax: any;
  }
}

