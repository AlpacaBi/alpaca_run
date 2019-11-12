<template>
<div class="page-item section">
  <div class="contact-title">Contact</div>
  <transition mode="out-in" enter-active-class="animated bounceIn" :duration="{ enter: 1000, leave: 1000 }">
    <div class="welcome" v-show="show0">欢迎各位大佬前来交流学♂习心得，共同进步！！</div>
  </transition>

  <div class="cards">
    <transition mode="out-in" enter-active-class="animated rotateInDownLeft" :duration="{ enter: 1000, leave: 1000 }">
      <div class="card" v-show="show1">
        <img src="@/assets/images/cphone.svg" alt="">
        <div class="text">17666503029</div>
      </div>
    </transition>
    <transition mode="out-in" enter-active-class="animated rotateInDownRight" :duration="{ enter: 1000, leave: 1000 }">
      <div class="card" v-show="show2" @click="openWechatQRCode">
        <img src="@/assets/images/cwechat.svg" alt="">
        <div class="text">workbiguokang</div>
      </div>
    </transition>
  </div>
  <div class="cards">
    <transition mode="out-in" enter-active-class="animated rotateInUpLeft" :duration="{ enter: 1000, leave: 1000 }">
      <div class="card" v-show="show3">
        <img src="@/assets/images/cgithub.svg" alt="">
        <div class="text"><a class="text" href="https://github.com/AlpacaBi" target="_blank">AlpacaBi</a></div>
      </div>
    </transition>
    <transition mode="out-in" enter-active-class="animated rotateInUpRight" :duration="{ enter: 1000, leave: 1000 }">
      <div class="card" v-show="show4">
        <img src="@/assets/images/cmail.svg" alt="">
        <div class="text"><a class="text" href="mailto:let@alpaca.run">let@alpaca.run</a></div>
      </div>
    </transition>
  </div>

  <transition mode="out-in" enter-active-class="animated fadeInUpBig" :duration="{ enter: 1000, leave: 1000 }">
    <div class="mp" v-show="show5">
      <img src="@/assets/images/wpQRCode.jpg" alt="">
      <div class="text">可以的话，请扫码关注我的公众号</div>
    </div>
  </transition>

  
</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Notification } from 'element-ui';
import { State } from 'vuex-class';

@Component
export default class Contact extends Vue {
  @State private contact!: any;
  private show0: boolean = false;
  private show1: boolean = false;
  private show2: boolean = false;
  private show3: boolean = false;
  private show4: boolean = false;
  private show5: boolean = false;

  private async showCard() {
    await new Promise( (resolve) => {
     setTimeout(() => { this.show0 = true; resolve(); }, 1500);
    });
    await new Promise( (resolve) => {
     setTimeout(() => { this.show1 = true; this.show2 = true; resolve(); }, 1000);
    });
    await new Promise( (resolve) => {
     setTimeout(() => { this.show3 = true; this.show4 = true; resolve(); }, 1000);
    });
    await new Promise( (resolve) => {
     setTimeout(() => { this.show5 = true; resolve(); }, 1000);
    });
  }

  private openWechatQRCode() {
    Notification({
      title: '扫描二维码加好友',
      dangerouslyUseHTMLString: true,
      message: '<img width=300px src="https://alpaca.cdn.bcebos.com/wechat.jpg" alt="">',
      position: 'bottom-left',
    });
  }

  private created() {
    this.showCard();
  }
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: #83B8EF !important;
  .contact-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
  .welcome{
    color: white;
    font-size: 26px;
    margin-top:80px;
  }
  .cards{
    margin-top: 25px;
    .card{
      display: inline-block;
      width: 250px;
      height: 55px;
      margin: 20px;
      cursor: pointer;
      background-color: white;
      border-radius: 6px;
      text-align: left;
      box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
      -webkit-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
      -o-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
      img{
        width: 40px;
        margin-top: 7px;
        margin-left: 16px;
      }
      .text{
        color: black;
        display: inline-block;
        vertical-align: top;
        margin-top: 16px;
        margin-left: 10px;
        font-size: 20px;
        a{
          display: inline-block;
          margin-top: 0px;
          text-decoration:none;
          color: black;
        }
      }
      
    }
    .card:hover{
      -webkit-transition:0.5s all;/* 适用Safari 和 Chrome兼容 */
      -moz-transition:0.5s all;/* 适用Firefox兼容 */
      -ms-transition:0.5s all;/* 适用IE9兼容 */
      -o-transition:0.5s all;/* Opera */
      transition: .5s;
      -webkit-transform:scale(1.1);/*在Safari 和 Chrome兼容下还原*/
      -moz-transform:scale(1.1);/*在Firefox兼容下还原*/
      -ms-transfrom:scale(1.1);/*在IE9兼容下还原*/
      -o-transform:scale(1.1);/*在opera兼容下还原*/
      transform:scale(1.1);/*适用于所有浏览兼容下还原*/
    }
  }
  .mp{
    margin-top: 60px;
    img{
      width: 150px;
      height: 150px;
    }
    .text{
      color: white;
      margin-top: 10px;
      font-size: 20px;
    }
  }
}
</style>
