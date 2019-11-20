import Vue from 'vue';
import Vuex from 'vuex';
import router from './router';
import { images, info, skills, profiles } from './data/data';

Vue.use(Vuex);

interface IAnimateOptions {
  leave: string;
  enter: string;
  leaveTime: number;
  enterTime: number;
}

interface IState {
  aiShow: boolean;
  loginDialogShow: boolean;
  num: number;
  routes: any;
  animateOptions: IAnimateOptions;
  showArrow: boolean;
  images: any;
  info: any;
  skills: any;
  profiles: any;
}

const State: IState = {
  aiShow: false,
  loginDialogShow: false,
  num: 0,
  routes: router,
  animateOptions: {
    leave: 'fadeOutDown',
    enter: 'fadeInDown',
    leaveTime: 1500,
    enterTime: 1500,
  },
  showArrow: true,
  images,
  info,
  skills,
  profiles,
};


export default new Vuex.Store({
  state: State,
  mutations: {
    add(state) {
      state.num++;
    },
    del(state) {
      state.num--;
    },
    changeAiShow(state) {
      state.aiShow = !state.aiShow;
    },
    openAiShow(state) {
      state.aiShow = true;
    },
    closeAiShow(state) {
      state.aiShow = false;
    },
    openLoginDialogShow(state) {
      state.loginDialogShow = true;
    },
    closeLoginDialogShow(state) {
      state.loginDialogShow = false;
    },
    // 改变切换路由页面的过度动画
    changeAnimateDirection(state, data) {
      state.animateOptions = data;
    },
    // 控制是否显示下一页的箭头
    changeShowArrow(state) {
      state.showArrow = !state.showArrow;
    },
  },
  actions: {
    // 下一页  =====> 鼠标滚轮向下滚动 或 向下滑动（移动端）
    next({ commit, state }, presentName) {
      return new Promise((resolve) => {
        // console.log(presentPath)
        // 获取当前页面路径在整体路由的位置
        const index = state.routes.options.routes[1].children.findIndex((route: any) => {
          return route.name === presentName;
        });
        // 获取当前路由的下一个路由路径
        let nextPageName;
        if (index + 1 === state.routes.options.routes[1].children.length) {
          nextPageName = state.routes.options.routes[1].children[0].name;
        } else {
          nextPageName = state.routes.options.routes[1].children[index + 1].name;
        }
        const animateDirection = {
          leave: 'zoomOut',
          enter: 'zoomIn',
          leaveTime: 1000,
          enterTime: 1000,
        };
        commit('changeAnimateDirection', animateDirection);
        // commit('nextPage', nextPagePath)
        commit('changeShowArrow');
        resolve(nextPageName);
        setTimeout(() => {
          commit('changeShowArrow');
        }, state.animateOptions.leaveTime + state.animateOptions.enterTime);
      });
    },
    // 上一页 鼠标滚轮向上滚动 或 向上滑动（移动端）
    last({ commit, state }, presentName) {
      return new Promise((resolve) => {
        // console.log(presentPath)
        // 获取当前页面路径在整体路由的位置
        const index = state.routes.options.routes[1].children.findIndex((route: any) => route.name === presentName);
        // 获取当前路由的上一个路由路径
        let lastPageName;
        if (index === 0) {
          const routesLength = state.routes.options.routes[1].children.length;
          lastPageName = state.routes.options.routes[1].children[routesLength - 1].name;
        } else {
          lastPageName = state.routes.options.routes[1].children[index - 1].name;
        }
        // console.log(lastPagePath)
        const animateDirection = {
          leave: 'lightSpeedOut',
          enter: 'lightSpeedIn',
          leaveTime: 1000,
          enterTime: 1000,
        };
        commit('changeAnimateDirection', animateDirection);
        // commit('lastPage', lastPagePath)
        commit('changeShowArrow');
        resolve(lastPageName);
        setTimeout(() => {
          commit('changeShowArrow');
        }, state.animateOptions.leaveTime + state.animateOptions.enterTime);
      });
    },
    // 点击小圆点切换页面
    clickRouteChange({ commit, state }) {
      const animateDirection = {
        leave: 'rollOut',
        enter: 'rollIn',
        leaveTime: 1000,
        enterTime: 1000,
      };
      commit('changeAnimateDirection', animateDirection);
      commit('changeShowArrow');
      setTimeout(() => {
        commit('changeShowArrow');
      }, state.animateOptions.leaveTime + state.animateOptions.enterTime);
    },
  },
});
