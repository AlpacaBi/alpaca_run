<template>
  <div @mousewheel="mouseWheel">
    <div class="nav">
      <template v-for="(item, index) in this.$router.options.routes[1].children">
        <router-link :to="{name: item.name}" :key="index" exact><i @click="clickRouteChange" class="nav-item"></i></router-link>
      </template>
    </div>
    <div :class="'page'+this.$route.name">
      <transition
        mode="out-in"
        :duration="{ enter: animateOptions.enterTime, leave: animateOptions.leaveTime }"
        :leave-active-class="'animated ' + animateOptions.leave"
        :enter-active-class="'animated ' + animateOptions.enter">
        <router-view />
      </transition>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State, Action } from 'vuex-class';
import 'animate.css';


@Component
export default class About extends Vue {
  @State private animateOptions: any;

  @Action private next!: (x: string | undefined) => any;
  @Action private last!: (x: string | undefined) => any;
  @Action private clickRouteChange!: () => void;

  private lastScroll: number = 0;

  private mouseWheel(event: any) {
    // 防止用户短时间内滚动多次，设置滚动间隔大于一秒才能生效
    // 判断滚动间隔时间
    const scrollDuration = event.timeStamp - this.lastScroll;
    if (scrollDuration > 1000) {
    // 将这一次的滚动时间记录为上一次合法的滚动时间
    this.lastScroll = event.timeStamp;
    // console.log('合法的滚动')
    // 判断滚动方向进行操作
    if (event.deltaY > 0) {
      const presentName: string | undefined = this.$route.name;
      this.next(presentName).then((nextPageName: string)  => {
        this.$router.push({name: nextPageName});
      });
    } else {
      const presentName: string | undefined = this.$route.name;
      this.last(presentName).then((lastPageName: string) => {
        this.$router.push({name: lastPageName});
      });
    }
  }
}

}
</script>

<style lang="scss" scoped>
.nav {
  position: fixed;
  left: 310px;
  top: 45%;
  z-index: 100;
  color: grey;
  a {
    text-decoration-line: none;
  }

  .nav-item {
    display: block;
    font-size: 11px;
    margin-top: 5px;
    height: 12px;
    width: 12px;
    border-radius: 100%;
    background-color: grey;
    opacity: 0.5;
  }
  .router-link-active {
    .nav-item {
      background-color: white !important;
    }
  }
}
.pageinfo{
  background-color: black;
}
.pageskills{
  background-color: #00ADD8;
}
.pageprofile{
  background-color: #4fc08d;
}
.pagecontact{
  background-color: rgba(13, 16, 110, 0.575);
}



</style>