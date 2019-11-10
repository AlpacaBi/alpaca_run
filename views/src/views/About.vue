<template>
  <div @mousewheel="mouseWheel">
    <div class="nav">
      <template v-for="(item, index) in this.$router.options.routes[1].children">
        <router-link :to="{name: item.name}" :key="index" exact><i class="iconfont icon-dian nav-item"></i></router-link>
      </template>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State, Action } from 'vuex-class'
import 'animate.css';


@Component
export default class About extends Vue {

  @Action next
  @Action last

  private lastScroll: number = 0

  private mouseWheel (event) {
    // 防止用户短时间内滚动多次，设置滚动间隔大于一秒才能生效
    // 判断滚动间隔时间
    let scrollDuration = event.timeStamp - this.lastScroll
    if (scrollDuration > 1000) {
    // 将这一次的滚动时间记录为上一次合法的滚动时间
    this.lastScroll = event.timeStamp
    // console.log('合法的滚动')
    // 判断滚动方向进行操作
    if (event.deltaY > 0) {
      const presentName = this.$route.name
      this.next(presentName).then(nextPageName => {
        this.$router.push({name:nextPageName})
      })
    } else {
      const presentName = this.$route.name
      this.last(presentName).then(lastPageName => {
        this.$router.push({name:lastPageName})
      })
    }
  }
}

}
</script>

<style lang="scss" scoped>
@import url('//at.alicdn.com/t/font_886340_sspbd7xw47.css');
.router-link-active {
  color: rgb(122, 7, 28);
  font-size: 1.1rem;
}
.nav {
  position: fixed;
  right: 10px;
  top: 45%;
  z-index: 100;
  color: grey;
  a {
    text-decoration-line: none;
  }
  .nav-item {
    display: block;
    font-size: 11px;
    padding: 3px 0;
  }
}

</style>