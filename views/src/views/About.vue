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