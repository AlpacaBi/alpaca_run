import Vue from 'vue';
import Vuex from 'vuex';
import router from './router';
import { info, skills, profiles,  contact } from './data/data'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    aiShow: false,
    num: 0,
    routes: router.options.routes[1].children,
    animateOptions: {},
    showArrow: true,
    info,
    skills,
    profiles,
    contact
  },
  mutations: {
    add(state) {
      state.num++;
    },
    del(state) {
      state.num--;
    },
    changeAiShow(state) {
      state.aiShow = !state.aiShow
    },
    closeAiShow(state) {
      state.aiShow = false
    },
    openAiShow(state) {
      state.aiShow = true
    },
    // 改变切换路由页面的过度动画
    changeAnimateDirection (state, data) {
      state.animateOptions = data
    },
    // 控制是否显示下一页的箭头
    changeShowArrow (state) {
      state.showArrow = !state.showArrow
    }
  },
  actions: {
    // 下一页  =====> 鼠标滚轮向下滚动 或 向下滑动（移动端）
    next ({ commit, state }, presentName) {
      return new Promise(resolve => {
        // console.log(presentPath)
        // 获取当前页面路径在整体路由的位置
        const index = state.routes.findIndex(route => {
          return route.name === presentName
        })
        // 获取当前路由的下一个路由路径
        let nextPageName
        console.log(index)
        if (index + 1 === state.routes.length) {
          nextPageName = state.routes[0].name
        } else {
          nextPageName = state.routes[index + 1].name
        }
        console.log("nextPageName:"+nextPageName)
        const animateDirection = {
          leave: 'fadeOutDown',
          enter: 'fadeInDown',
          leaveTime: 1500,
          enterTime: 1500
        }
        commit('changeAnimateDirection', animateDirection)
        // commit('nextPage', nextPagePath)
        commit('changeShowArrow')
        resolve(nextPageName)
        setTimeout(() => {
          commit('changeShowArrow')
        }, state.animateOptions.leaveTime + state.animateOptions.enterTime)
      })
    },
    // 上一页 鼠标滚轮向上滚动 或 向上滑动（移动端）
    last ({ commit, state }, presentName) {
      return new Promise(resolve => {
        // console.log(presentPath)
        // 获取当前页面路径在整体路由的位置
        const index = state.routes.findIndex(route => route.name === presentName)
        // 获取当前路由的上一个路由路径
        let lastPageName
        if (index === 0) {
          const routesLength = state.routes.length
          lastPageName = state.routes[routesLength - 1].name
        } else {
          lastPageName = state.routes[index - 1].name
        }
        // console.log(lastPagePath)
        const animateDirection = {
          leave: 'fadeOutUp',
          enter: 'fadeInUp',
          leaveTime: 1500,
          enterTime: 1500
        }
        commit('changeAnimateDirection', animateDirection)
        // commit('lastPage', lastPagePath)
        commit('changeShowArrow')
        resolve(lastPageName)
        setTimeout(() => {
          commit('changeShowArrow')
        }, state.animateOptions.leaveTime + state.animateOptions.enterTime)
      })
    },
    // 点击小圆点切换页面
    clickRouteChange ({ commit, state }) {
      const animateDirection = {
        leave: 'zoomOut',
        enter: 'zoomIn',
        leaveTime: 1500,
        enterTime: 1500
      }
      commit('changeAnimateDirection', animateDirection)
      commit('changeShowArrow')
      setTimeout(() => {
        commit('changeShowArrow')
      }, state.animateOptions.leaveTime + state.animateOptions.enterTime)
    }
  },
});
