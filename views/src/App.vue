<template>
  <div id="app">
    <audio style="display:none" src="" controls id="sound"></audio>

    <!-- <LeftMenu/> -->
    <TopBar/>

    <div class="info">
      <div class="name">Alpaca Bi</div>
      <div class="intro">一个来自广州的95后全栈工程师</div>
      <div class="more">想了解更多？请在终端输入以下命令</div>
      <el-tooltip content="点击即可复制" placement="right" effect="dark">
        <div class="terminal" @click="copy">
          npx alpaca-bi
        </div>
      </el-tooltip>
    </div>

    <!-- <div class="right-content"><router-view/></div> -->

    <AlpacaAI/>
    
    <!-- <LoginDialog/> -->

  </div>
</template>

<script lang="ts">
  import { Component,  Vue } from 'vue-property-decorator';
  import { State, Mutation } from 'vuex-class';
  import { Notification } from 'element-ui';
  import LeftMenu from '@/components/LeftMenu.vue';
  import AlpacaAI from '@/components/AlpacaAI.vue';
  import LoginDialog from '@/components/LoginDialog.vue';
  import TopBar from '@/components/TopBar.vue';
  import sleep from '@/utils/sleep';

  @Component({
    components: {
      LeftMenu,
      AlpacaAI,
      LoginDialog,
      TopBar,
    },
  })
  export default class APP extends Vue {
    @State private aiShow!: boolean;
    @Mutation private openAiShow!: () => void;

    private dialogFormVisible: boolean = false;

    private async created() {
      const x = document.createElement('div');
      Object.defineProperty(x, 'id', {
          get: async () => {
            const frog = `
              ⣿⣿⣿⣿⣿⣿⢟⣡⣴⣶⣶⣦⣌⡛⠟⣋⣩⣬⣭⣭⡛⢿⣿⣿⣿⣿
              ⣿⣿⣿⣿⠋⢰⣿⣿⠿⣛⣛⣙⣛⠻⢆⢻⣿⠿⠿⠿⣿⡄⠻⣿⣿⣿
              ⣿⣿⣿⠃⢠⣿⣿⣶⣿⣿⡿⠿⢟⣛⣒⠐⠲⣶⡶⠿⠶⠶⠦⠄⠙⢿
              ⣿⠋⣠⠄⣿⣿⣿⠟⡛⢅⣠⡵⡐⠲⣶⣶⣥⡠⣤⣵⠆⠄⠰⣦⣤⡀
              ⠇⣰⣿⣼⣿⣿⣧⣤⡸ ⣿⡀⠂⠁⣸⣿⣿⣿⣿⣇⠄⠈⢀⣿⣿⠿
              ⣰⣿⣿⣿⣿⣿⣿⣿⣷⣤⣈⣙⠶⢾⠭⢉⣁⣴⢯⣭⣵⣶⠾⠓⢀⣴
              ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣉⣤⣴⣾⣿⣿⣦⣄⣤⣤⣄⠄⢿⣿
              ⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠈⢿
              ⣿⣿⣿⣿⣿⣿⡟⣰⣞⣛⡒⢒⠤⠦⢬⣉⣉⣉⣉⣉⣉⣉⡥⠴⠂⢸
              ⠻⣿⣿⣿⣿⣏⠻⢌⣉⣉⣩⣉⡛⣛⠒⠶⠶⠶⠶⠶⠶⠶⠶⠂⣸⣿
              ⣥⣈⠙⡻⠿⠿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⠿⠛⢉⣠⣶⣶⣿⣿
              ⣿⣿⣿⣶⣬⣅⣒⣒⡂⠈⠭⠭⠭⠭⠭⢉⣁⣄⡀⢾⣿⣿⣿⣿⣿⣿
            `;
            console.log('%c检测到你打开了控制台', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c5', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c4', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c3', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c2', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c1', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log('%c ', 'background: url(https://alpaca.cdn.bcebos.com/blackpeople.gif) no-repeat center;padding-left:640px;padding-bottom: 242px;');
            console.log('%c棺材一抬,世间白来', 'font-size:20px;color:white;background:black');
            console.log('%c脚一滑,人一叫,孟婆端汤阎王笑', 'font-size:20px;color:white;background:black');
            console.log('%c眼一闭,布一盖,亲戚朋友等上菜', 'font-size:20px;color:white;background:black');
            console.log('%c棺一抬,土一埋,一人一句古德拜', 'font-size:20px;color:white;background:black');
            await sleep(1.5);
            console.log(frog);
          },
      });
      console.log(x);
      const arr: string[] = [
        '帕秋莉~GO！',
        'ASS WE CAN！！',
        'boy next door',
        'f♂ck y♂u',
        'Oh！~YES,SIR♂',
        'OH~卖箫的~',
        '啊？乖♂乖♂站♂好',
        '啊？幻想乡~',
        '啊啊啊啊啊！！',
        '奥义很爽',
        '乖♂乖♂站♂好',
      ];

      document.onclick = () => {
          const soundID = Math.floor(Math.random() * arr.length);
          console.log(arr[soundID]);
          document.title = arr[soundID];
          setTimeout( () => {
            document.title = 'Free Technology Blog & Alpaca Bi Website';
          }, 3000);
          const url = `https://alpaca.cdn.bcebos.com/zhexue/${soundID}.mp3`;
          const sound: any = document.getElementById('sound');
          sound.src = url;
          sound.play();
      };

      await sleep(2);
      this.openAiShow();
    }

    private copy() {
      const input = document.createElement('input');
      document.body.appendChild(input);
      input.setAttribute('value', 'npx alpaca-bi');
      input.select();
      if (document.execCommand('copy')) {
          document.execCommand('copy');
      }
      document.body.removeChild(input);
      Notification({
        title: '成功',
          message: '复制成功',
          type: 'success',
          position: 'top-left',
      });
    }

  }
</script>

<style lang="scss">
body{
  margin: 0;
  color: white;
  background: black;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
.right-content {
  margin-left: 300px;
  background: #24292e;
  height: 100vh;
}
.info {
  position: absolute;
  display: flex;
  top: 0;
  height: 99.5vh;
  width: 99.5vw;
  color: white;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  @media screen and (min-width: 769px){
    .name {
      font-size: 80px;
    }
    .intro {
      font-size: 40px;
      margin-bottom: 200px;
    }
    .more {
      margin-bottom: 15px;
    }
    .terminal {
      border: 1px solid #ff9900;
      border-radius: 8px;
      padding: 13px;
      font-size: 20px;
      color: #ff9900;
      &:hover{
        color: white;
        background: #ff9900;
      }
    }
  }
  @media screen and (max-width: 767px){
    .name {
      font-size: 60px;
      margin-bottom: 20px;
    }
    .intro {
      font-size: 22px;
      margin-bottom: 180px;
    }
    .more {
      margin-bottom: 15px;
    }
    .terminal {
      border: 1px solid #ff9900;
      border-radius: 8px;
      padding: 13px;
      font-size: 20px;
      color: #ff9900;
      &:hover{
        color: white;
        background: #ff9900;
      }
    }
  }
}
</style>