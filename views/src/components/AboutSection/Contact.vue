<template>
  <div class="page-item section">
    <!-- <div class="contact-title">Contact</div> -->
    <div class="welcome" v-show="show0">欢迎各位大佬前来交流学♂习心得，共同进步！！</div>

    <div class="cards">
      <transition mode="out-in" enter-active-class="animated slideInUp" :duration="{ enter: 1000, leave: 1000 }">
        <div class="card" v-show="show1" @mouseenter="phone = '打电话'" @mouseleave="phone = '17666503029'">
          <img :src="images.cphone" alt="">
          <div class="text"><a class="text" href="tel:17666503029">{{phone}}</a></div>
        </div>
      </transition>
      <transition mode="out-in" enter-active-class="animated slideInUp" :duration="{ enter: 1000, leave: 1000 }">
        <div class="card" v-show="show2" @click="openWechatQRCode" @mouseenter="wechat = '弹出微信二维码'" @mouseleave="wechat = 'workbiguokang'">
          <img :src="images.cwechat" alt="">
          <div class="text">{{wechat}}</div>
        </div>
      </transition>
      <transition mode="out-in" enter-active-class="animated slideInUp" :duration="{ enter: 1000, leave: 1000 }">
        <div class="card" v-show="show3" @mouseenter="github = '进入AlpacaBi的GitHub'" @mouseleave="github = 'AlpacaBi'">
          <img :src="images.cgithub" alt="">
          <div class="text"><a class="text" href="https://github.com/AlpacaBi" target="_blank">{{github}}</a></div>
        </div>
      </transition>
      <transition mode="out-in" enter-active-class="animated slideInUp" :duration="{ enter: 1000, leave: 1000 }">
        <div class="card" v-show="show4" @mouseenter="email = '发邮件'" @mouseleave="email = 'biguokang@outlook.com'">
          <img :src="images.cmail" alt="">
          <div class="text"><a class="text" href="Mailto:biguokang@outlook.com">{{email}}</a></div>
        </div>
      </transition>
    </div>

    <div class="next" @click="nextPage">Back To Home</div>

    <!-- <transition mode="out-in" enter-active-class="animated fadeInUpBig" :duration="{ enter: 1000, leave: 1000 }">
      <div class="mp" v-show="show5">
        <img :src="images.wpQRCode" alt="">
        <div class="text">可以的话，请扫码关注我的公众号</div>
      </div>
    </transition> -->

    
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Notification } from 'element-ui';
import { State, Action } from 'vuex-class';
import sleep from '@/utils/sleep';

@Component
export default class Contact extends Vue {
  @State private contact!: any;
  @State private images!: any;
  @Action private next!: (x: string | undefined | null ) => any;

  private show0: boolean = false;
  private show1: boolean = false;
  private show2: boolean = false;
  private show3: boolean = false;
  private show4: boolean = false;
  private show5: boolean = false;

  private phone: string = '17666503029';
  private wechat: string = 'workbiguokang';
  private github: string = 'AlpacaBi';
  private email: string = 'biguokang';

  private created() {
    this.showCard();
  }

  private async showCard() {
    await sleep(0.2);
    this.show0 = true;
    await sleep(1);
    this.show1 = true;
    await sleep(1);
    this.show2 = true;
    await sleep(1);
    this.show3 = true;
    await sleep(1);
    this.show4 = true;
  }

  private openWechatQRCode() {
    Notification({
      title: '扫描二维码加好友',
      dangerouslyUseHTMLString: true,
      message: '<img width=300px src="https://alpaca.cdn.bcebos.com/wechat.jpg" alt="">',
      position: 'bottom-left',
    });
  }

  private nextPage() {
    const presentName: string | undefined | null = this.$route.name;
    this.next(presentName).then((nextPageName: string)  => {
      this.$router.push({name: nextPageName});
    });
  }

}
</script>

<style lang="scss" scoped>
.section {
  height: calc(100vh - 80px);
  width: 98vw;
  background-color: black !important;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .contact-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
  @media screen and (min-width: 769px){
    .welcome{
      color: white;
      font-size: 50px;
      margin-bottom: 80px;
    }
    .next{
      display: none;
    }
  }
  @media screen and (max-width: 767px){
    .welcome{
      color: white;
      font-size: 25px;
      margin-bottom: 80px;
    }
    .next {
      position: absolute;
      color: #ff9900;
      bottom: 20px;
      font-size:24px;
      text-decoration: underline;
    }
  }
 
  .cards{
    // margin-top: 25px;
    .card{
      // display: inline-block;
      width: 330px;
      height: 55px;
      margin: 20px;
      cursor: pointer;
      background-color: #ff9900;
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
