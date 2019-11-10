<template>
  <div id="app">

    <LeftMenu/>

    <div class="right-content"><router-view/></div>

    <AlpacaAI/>

  </div>
</template>

<style lang="scss">
body{
  margin: 0
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.right-content {
  margin-left: 300px;
}
</style>
<script lang="ts">
  import { Component,  Vue } from 'vue-property-decorator';
  import { State, Mutation } from 'vuex-class'
  import LeftMenu from '@/components/LeftMenu.vue';
  import AlpacaAI from '@/components/AlpacaAI.vue';

  @Component({
    components: {
      LeftMenu,
      AlpacaAI,
    },
  })
  export default class APP extends Vue {
    @State aiShow!: boolean
    @Mutation openAiShow!: Function

    private created() {
      setTimeout(() => {
        this.openAiShow();
      }, 2000);
    }
  }
</script>