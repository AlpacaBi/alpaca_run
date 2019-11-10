<template>
  <div class="page-item section">
    <mu-carousel class="carousel" :hide-controls="controls" :hide-indicators="indicators" :interval="interval"  transition="slide" :cycle="cycle">
      <template v-for="(profile, index) in profiles">
        <mu-carousel-item :key="index">
          <mu-container class="box">
            <mu-card class="card" style="width: 100%; margin: 0 auto;">
              <mu-card-header :title="profile.title" :sub-title="profile.subTitle" />
              <!-- <mu-card-media class="img">
                <img src="../assets/design.jpg">
              </mu-card-media> -->
              <mu-card-title :sub-title="'技术栈：' + profile.skills"></mu-card-title>
              <mu-card-text>
                {{profile.text}}
              </mu-card-text>
              <mu-card-actions><a :href="profile.github" target="_blank">
                  <mu-icon class="icon" value=":iconfont icon-GitHub"></mu-icon></a>
              </mu-card-actions>
            </mu-card>
          </mu-container>
        </mu-carousel-item>
      </template>
    </mu-carousel>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class'

@Component
export default class Profile extends Vue {
  @State profile

  private cycle: boolean = true
  private indicators: boolean = true
  private controls: boolean = false
  private interval: number = 10000
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: rgb(73, 80, 10) !important;
  .carousel {
    // height: auto;
    min-height: 50vh;
    .mu-carousel-item {
      display: flex;
      justify-content: center;
      align-items: center;
      .box {
        width: 85%;
        .card {
          background: rgba(126, 123, 94, 0.664);
          border-radius: 10px;
          .mu-card-header {
            padding-bottom: 0;
          }
          .mu-card-header-title {
            padding: 0;
            width: 100%;
          }
          .mu-card-title {
            width: 100%;
            color: rgb(0, 0, 0);
            font-size: 1.3rem;
            text-align: left;
          }
          .mu-card-sub-title {
            color: rgb(34, 34, 34);
            text-align: left;
          }
          .mu-card-text {
            padding: 0 16px;
            // letter-spacing: .07rem;
            letter-spacing: 1px;
            color: rgb(0, 0, 0);
            text-align: left;
          }
          .mu-card-actions {
            padding: 0;
            .icon {
              font-size: 30px;
              color: #000;
            }
          }
          .img {
            @media screen and (min-width: 960px) {
              display: none;
            }
          }
        }
      }
    }
  }
}
</style>
