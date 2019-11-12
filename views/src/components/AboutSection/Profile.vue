<template>
  <div class="page-item section">
    <div class="profile-title">Profile</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';

@Component
export default class Profile extends Vue {
  @State private profile!: any;
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: #4fc08d !important;
  .profile-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
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
