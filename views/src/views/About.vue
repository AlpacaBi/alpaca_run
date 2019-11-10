<template>
  <div>
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


@Component
export default class About extends Vue {
  private created() {
    console.log(this.$router.options.routes)
  }

}
</script>

<style lang="scss" scoped>
.nav {
  position: fixed;
  right: 10px;
  top: 45%;
  z-index: 100;
  .iconfont {
    font-family: "iconfont" !important;
    font-size: 16px;
    font-style: normal;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
  .icon-dian:before {
    content: "\e621";
    box-sizing: border-box;
  }
  .nav-item {
    display: block;
    font-size: 11px;
    padding: 3px 0;
  }
}

</style>