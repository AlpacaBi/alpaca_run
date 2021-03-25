<template>
  <div id="app">
    <audio style="display:none" src="" controls id="sound"></audio>

    <!-- <LeftMenu/> -->
    <Home />
    <TopBar/>
    

    <!-- <div class="right-content"><router-view/></div> -->

    <AlpacaAI/>
    
    <!-- <LoginDialog/> -->

  </div>
</template>

<script lang="ts">
  import { Component,  Vue } from 'vue-property-decorator';
  import { State, Mutation, Action } from 'vuex-class';
  import LeftMenu from '@/components/LeftMenu.vue';
  import Home from '@/views/Home.vue';
  import AlpacaAI from '@/components/AlpacaAI.vue';
  import LoginDialog from '@/components/LoginDialog.vue';
  import TopBar from '@/components/TopBar.vue';
  import sleep from '@/utils/sleep';
  import {repoData} from '@/data/data';

  @Component({
    components: {
      LeftMenu,
      AlpacaAI,
      LoginDialog,
      TopBar,
      Home,
    },
  })
  export default class APP extends Vue {
    @State private aiShow!: boolean;
    @Mutation private openAiShow!: () => void;
    @Action private setRepos!: (repos: any ) => void;

    private dialogFormVisible: boolean = false;

    private async created() {
      this.getGithubRepo();
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

      await sleep(1);
      this.openAiShow();
    }

    // 加载github api
    private async getGithubRepo() {
      // const url: string = '/github';
      // const res: any = await this.$ajax.get(url);
      // const repos = res.data.sort((a: any, b: any) => b.stargazers_count - a.stargazers_count).slice(0, 6);
      // console.log(repos);
      this.setRepos(repoData);
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
</style>