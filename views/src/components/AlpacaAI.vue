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
          <template v-for="(item, index) in chatlist">
          <div :key="index">
            <div v-if="item.type == 0">
              <div class="left-bubble">{{item.content}}</div>
            </div>
            <div v-else>
              <div class="right-bubble">{{item.content}}</div>
            </div>
          </div>
        </template>
        </div>
        <div class="footer">
          <input v-model.trim="chatContent" placeholder="输入文字即可和Alapca AI聊天！！！" @keyup.enter="sendText"/>
          <el-tooltip content="发送文字" placement="top">
            <img class="sendtext" @click="sendText" src="@/assets/images/sendtext.svg" alt="" >
          </el-tooltip>
          <el-tooltip content="发送图片" placement="top">
            <img class="sendimg" src="@/assets/images/sendimg.svg" alt="" @click="sendImage">
          </el-tooltip>
        </div>
      </div>
    </transition>
</template>

<script lang="ts">
import { Component, Prop, Vue, PropSync } from 'vue-property-decorator';
import { State, Mutation } from 'vuex-class';
import { Notification } from 'element-ui';
import 'animate.css';

interface IChatArrList {
  type: number;
  content: string;
}

@Component
export default class AlpacaAI extends Vue {
  @State private aiShow!: boolean;
  @Mutation private closeAiShow!: () => void;

  private chatlist: IChatArrList[] = [];
  private chatContent: string = '';

  private async ping() {
    const res: any = await this.$ajax.get('/ping');
    console.log(res);
  }
  private AlpacaAISaid(content: string): void {
    this.chatlist.push({type: 0, content});
    const container: any = this.$el.querySelector('.content');
    container.scrollTop = container.scrollHeight;
  }
  private VisitorSaid(content: string): void {
    this.chatlist.push({type: 1, content});
    const container: any = this.$el.querySelector('.content');
    container.scrollTop = container.scrollHeight;
  }
  private async sendText() {
    const container: any = this.$el.querySelector('.content');
    if (this.chatContent === '') {
      this.AlpacaAISaid('抱歉，我没听到');
      return;
    } else if (this.chatContent.indexOf('微信') !== -1 || this.chatContent.indexOf('wechat') !== -1) {
      this.AlpacaAISaid('为您显示Alpaca Bi的微信二维码');
      Notification({
        title: '扫描二维码加好友',
        dangerouslyUseHTMLString: true,
        message: '<img width=300px src="https://alpaca.cdn.bcebos.com/wechat.jpg" alt="">',
        position: 'bottom-left',
      });
    } else if (this.chatContent.indexOf('毕国康') !== -1 ||
              this.chatContent.indexOf('Alpaca') !== -1 ||
              this.chatContent.indexOf('alpaca') !== -1) {
        this.AlpacaAISaid('正在为您跳转到Alpaca Bi个人介绍');
        this.$router.push({name: 'info'});
    } else {
      this.VisitorSaid(this.chatContent);
      const res: any = await this.$ajax.post('/ai/text', {
        text: this.chatContent,
      });
      this.chatContent = '';
      if (res.data.status === 'ok') {
        this.AlpacaAISaid(res.data.message);
      } else {
        this.AlpacaAISaid('抱歉，AI服务器网络错误');
      }
    }
  }
  private sendImage() {
    this.AlpacaAISaid('图片识别功能开发中...');
  }
  private created() {
    setTimeout(() => {
      this.AlpacaAISaid('你好啊，我是Alpaca AI,一个人工智能，你可以打字和我聊天！！！！');
    }, 3500);
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
    padding-bottom: 80px;
    .left-bubble{
      float: left;
      width: 300px;
      text-align: left;
      padding: 10px;
      margin: 10px;
      color: white;
      word-break:break-all;
      background-color: grey;
      border-radius: 8px;
      -webkit-box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      -moz-box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      -o-box-shadow:0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
    }
    .right-bubble{
      float: right;
      width: 300px;
      text-align: left;
      padding: 10px;
      margin: 10px;
      color: black;
      word-break:break-all;
      border-radius: 8px;
      -webkit-box-shadow: -10px 0px 10px rgba(0, 0, 0, 0.05);
      box-shadow: -10px 0px 10px rgba(0, 0, 0, 0.05);
      -webkit-box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      -moz-box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      box-shadow: 0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
      -o-box-shadow:0 1px 4px rgba(0,0,0,0.3),0 0 40px rgba(0,0,0,0.1) inset;
    }
  }
  .content::-webkit-scrollbar {/*滚动条整体样式*/
    margin-right: 20px;
    width: 4px;     /*高宽分别对应横竖滚动条的尺寸*/
    height: 4px;
  }
  .content::-webkit-scrollbar-thumb {/*滚动条里面小方块*/
      border-radius: 2px;
      --webkit-box-shadow: inset 0 0 5px grey;
      background:rgba(239,239,239,1);
  }
  .content::-webkit-scrollbar-track {/*滚动条里面轨道*/
      --webkit-box-shadow: inset 0 0 5px transparent;
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
      font-size: 12px;
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
    .sendtext:active{
      transform: scale(0.9)
    }
    .sendimg{
      display: inline-block;
      vertical-align: bottom;
      margin-left: 7px;
      height: 34px;
      width: 34px;
      cursor: pointer;
    }
    .sendimg:active{
      transform: scale(0.9)
    }
  }
}
</style>
