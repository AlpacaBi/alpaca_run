<template>
  <div id="app">

    <LeftMenu @emitAiShow="aiShowChange"/>

    <div class="right-content"><router-view/></div>

    <AlpacaAI :aiShow="aiShow"/>

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
</style>
<script lang="ts">
  import animate from 'animate.css'
  import { Component,  Vue } from 'vue-property-decorator';
  import LeftMenu from '@/components/LeftMenu.vue'
  import AlpacaAI from '@/components/AlpacaAI.vue'

  @Component({
    components: {
      LeftMenu,
      AlpacaAI
    },
  })
  export default class APP extends Vue {
    private aiShow:boolean = false;

    private aiShowChange() {
      this.aiShow = !this.aiShow;
    }

    created() {
      setTimeout(() => {
        this.aiShow = true;
      }, 2000);
    }
  }
</script>