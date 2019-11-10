<template>
  <div class="page-item section">
    <mu-avatar size=80>
      <img :src="info.avatar">
    </mu-avatar>
    <h4>{{info.name}}</h4>
    <h3>{{info.introduce}}</h3>
    <mu-container class="papers">
      <template v-for="(item, index) in info.status">
        <mu-paper :key="index" class="paper" circle :z-depth="4">
          <i :class="`iconfont ${item.icon}`"></i>
          <div class="title">{{item.value}}</div>
        </mu-paper>
      </template>
    </mu-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class'


@Component
export default class Info extends Vue {
  @State info
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: rgb(20, 73, 73) !important;
  .papers {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    .paper {
      margin: 15px;
      width: 50px;
      height: 50px;
      color: rgb(192, 200, 207);
      background-color: rgba(97, 136, 207, 0.363);
      .iconfont {
        height: 50px;
        line-height: 50px;
        font-size: 30px;
      }
      .title {
        font-size: .9rem;
        margin: 5px 0;
      }
    }
  }
}
</style>


