<template>
  <transition enter-active-class="animated slideInRight" leave-active-class="animated slideOutRight">
      <div class="ai" v-show="aiShow">
        <div class="title" @click="ping">
          Alpaca AI
        </div>
        <div class="close" @click="closeAiShow">
          <img src="@/assets/images/close.svg" alt="">
        </div>
        <div class="content">
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
          <div style="font-size:50px">content</div>
        </div>
        <div class="footer">
          <input placeholder="输入文字即可和Alapca AI聊天！！！"/>
          <img class="sendtext" src="@/assets/images/sendtext.svg" alt="">
          <img class="sendimg" src="@/assets/images/sendimg.svg" alt="">
        </div>
      </div>
    </transition>
</template>

<script lang="ts">
import { Component, Prop, Vue, PropSync } from 'vue-property-decorator';
import { State, Mutation } from 'vuex-class';
import 'animate.css';

@Component
export default class AlpacaAI extends Vue {
  @State private aiShow!: boolean;
  @Mutation private closeAiShow!: () => void;
  private async ping() {
    const res: any = await this.$ajax.get('/ping');
    console.log(res);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.right-content{
  display: inline-block;
  margin-left: 300px
}
.ai{
  position: fixed;
  z-index: 2;
  width: 450px;
  top: 0;
  right: 0;
  height:100vh;
  background:rgba(255,255,255,1);
  box-shadow:-10px 0px 10px rgba(0,0,0,0.05);
  opacity:1;
  .title{
    text-align: left;
    background: #545c64;
    color: white;
    height: 42px;
    font-size: 25px;
    padding-top: 10px;
    padding-left: 20px;
    position: absolute;
    width: 100%;
  }
  .close{
    width: 35px;
    height: 35px;
    position: absolute;
    right: 7px;
    top: 8px;
    cursor: pointer;
  }
  .content{
    width: 100%;
    overflow: auto;
    top: 52px;
    position: absolute;
    z-index: 10;
    bottom: 55px;
  }
  .content::-webkit-scrollbar {/*滚动条整体样式*/
    margin-right: 20px;
    width: 4px;     /*高宽分别对应横竖滚动条的尺寸*/
    height: 4px;
  }
  .content::-webkit-scrollbar-thumb {/*滚动条里面小方块*/
      border-radius: 2px;
      -webkit-box-shadow: inset 0 0 5px grey;
      background:rgba(239,239,239,1);
  }
  .content::-webkit-scrollbar-track {/*滚动条里面轨道*/
      -webkit-box-shadow: inset 0 0 5px transparent;
      border-radius: 0;
      background: transparent;
  }
  .footer{
    position: absolute;
    bottom: 0px;
    width: 100%;
    height: 55px;
    background: grey;
    text-align: left;
    input{
      display: inline-block;
      width: 270px;
      height: 32px;
      border-radius: 6px;
      margin-top: 11px;
      margin-left: 40px;
      outline: none;
      border: none;
      font-size: 18px;
      padding-right: 10px;
      padding-left: 10px;
    }
    .sendtext{
      display: inline-block;
      vertical-align: bottom;
      margin-left: 7px;
      height: 30px;
      width: 34px;
      cursor: pointer;
    }
    .sendimg{
      display: inline-block;
      vertical-align: bottom;
      margin-left: 7px;
      height: 34px;
      width: 34px;
      cursor: pointer;
    }
  }
}
</style>
