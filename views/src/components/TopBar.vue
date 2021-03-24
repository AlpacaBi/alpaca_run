<template>
  <div id="header">
    <div id="headerWrapper">
      <div id="headerContainer">
        <div class="logo pornhub_logo" @click="changeAiShow">
          <span class="text1">Alpaca</span><span class="text2">Run</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component,  Vue , Emit} from 'vue-property-decorator';
import { Notification } from 'element-ui';
import { State, Mutation } from 'vuex-class';

@Component
export default class TopBar extends Vue {
  @State private images!: any;
  @Mutation private changeAiShow!: () => void;
  @Mutation private openLoginDialogShow!: () => void;

  private handleOpen(key: any, keyPath: any) {
    console.log(key, keyPath);
  }
  private handleClose(key: any, keyPath: any) {
    console.log(key, keyPath);
  }

  private openWechatQRCode() {
    Notification({
      title: '扫描二维码加好友',
      dangerouslyUseHTMLString: true,
      message: '<img width=300px src="https://alpaca.cdn.bcebos.com/wechat.jpg" alt="">',
      position: 'bottom-left',
    });
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss" scoped>

#header {
  display: grid;
  width: 100%;
  grid-template-rows: 60px 40px;
  z-index: 2;
  position: fixed;
  #headerWrapper {
    width: 100%;
    background-color: #1b1b1b;
    #headerContainer {
      height: 100%;
      display: grid;
      grid-template-columns: 200px 1fr 250px 220px;
      grid-gap: 10px;
      justify-items: center;
      align-items: center;
      margin: auto;
      min-width: 991px;
      max-width: 1323px;
      width: 96%;
      .logo {
        display: inline-block;
        font-size: 30px;
        font-family: Arial, Helvetica, sans-serif;
        font-weight: 700;
        margin: 0;
        padding: 0;
        cursor: pointer;
        .text1 {
          color: white;
          padding: 5px 5px;
        }
        .text2 {
          color: rgb(0, 0, 0);
          background-color: rgb(255, 153, 0);
          padding: 5px 10px;
          border-radius: 7px;
        }
      }
    }
  } 
}
</style>
