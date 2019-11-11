<template>
<div class="page-item section">
    contact
</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';

@Component
export default class Contact extends Vue {
  @State private contact!: any;
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  /*background-color: rgba(13, 16, 110, 0.575) !important;*/
  .contact {
    width: 85%;
    text-align: left;
    h3 {
      // text-align: center;
      margin-top: 0;
    }
    .experience {
      .p-item {
        margin: 5px 0;
      }
    }
    .socials {
      .box {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        .socials-item {
          display: block;
          height: 40px;
          width: 40px;
          background: rgba(11, 5, 48, 0.3);
          border-radius: 20px;
          text-align: center;
          .icon {
            font-size: 1.5rem;
            line-height: 40px;
            color: rgb(162, 165, 168);
          }
        }
      }
    }
    .contact-me {
      .btn {
        margin: 0 auto;
        display: block;
        height: 30px;
        width: 85%;
        background: rgba(11, 5, 48, 0.3);
        border-radius: 15px;
        font-size: 1.1rem;
        line-height: 30px;
        color: rgb(148, 154, 161);
        margin-bottom: 5px;
        text-align: left;
        text-indent: 1rem;
        .icon {
          font-size: 1.1rem;
          color: rgb(210, 213, 216);
          padding: 0 10px;
        }
      }
    }
  }
}
</style>
