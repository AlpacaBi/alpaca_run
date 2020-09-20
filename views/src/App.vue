<template>
  <div id="app">

    <!-- <LeftMenu/> -->
    <TopBar/>

    <!-- <div class="right-content"><router-view/></div> -->

    <AlpacaAI/>
    
    <!-- <LoginDialog/> -->

  </div>
</template>

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
<script lang="ts">
  import { Component,  Vue } from 'vue-property-decorator';
  import { State, Mutation } from 'vuex-class';
  import LeftMenu from '@/components/LeftMenu.vue';
  import AlpacaAI from '@/components/AlpacaAI.vue';
  import LoginDialog from '@/components/LoginDialog.vue';
  import TopBar from '@/components/TopBar.vue';

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

    private created() {
      const x = document.createElement('div');
      Object.defineProperty(x, 'id', {
          get: () => {
            console.log('%c检测到你打开了控制台', 'font-size:20px;color:white;background:black');
            setTimeout(() => {
              console.log('%c5', 'font-size:20px;color:white;background:black');
              setTimeout(() => {
                console.log('%c4', 'font-size:20px;color:white;background:black');
                setTimeout(() => {
                  console.log('%c3', 'font-size:20px;color:white;background:black');
                  setTimeout(() => {
                    console.log('%c2', 'font-size:20px;color:white;background:black');
                    setTimeout(() => {
                      console.log('%c1', 'font-size:20px;color:white;background:black');
                      setTimeout(() => {
                        console.log('%c ', 'background: url(https://alpaca.cdn.bcebos.com/blackpeople.gif) no-repeat center;padding-left:640px;padding-bottom: 242px;');
                        console.log('%c棺材一抬,世间白来', 'font-size:20px;color:white;background:black');
                        console.log('%c脚一滑,人一叫,孟婆端汤阎王笑', 'font-size:20px;color:white;background:black');
                        console.log('%c眼一闭,布一盖,亲戚朋友等上菜', 'font-size:20px;color:white;background:black');
                        console.log('%c棺一抬,土一埋,一人一句古德拜', 'font-size:20px;color:white;background:black');
                        setTimeout(() => {
                          const a = `
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
                          console.log(a);
                        }, 1500);
                      }, 1500);
                    }, 1500);
                  }, 1500);
                }, 1500);
              }, 1500);
            }, 1500);
          },
      });
      console.log(x);
      setTimeout(() => {
        this.openAiShow();
      }, 2000);
    }
  }
</script>