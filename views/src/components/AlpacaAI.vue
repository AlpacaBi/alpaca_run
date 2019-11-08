<template>
  <transition enter-active-class="animated slideInRight" leave-active-class="animated slideOutRight">
      <div class="ai" v-show="aiShow">
        <div class="title">
          Alpaca AI
        </div>
        <div class="close" @click="closeAlpacaAI">
          <img src="@/assets/close.svg" alt="">
        </div>
      </div>
    </transition>
</template>

<script lang="ts">
import { Component, Prop, Vue, PropSync } from 'vue-property-decorator';
import 'animate.css';

@Component
export default class LeftMenu extends Vue {
    @PropSync('syncedAiShow', { type: Boolean, default: false }) public aiShow!: boolean;
    private closeAlpacaAI() {
      this.aiShow = false;
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
  }
  .close{
    width: 35px;
    height: 35px;
    position: absolute;
    right: 7px;
    top: 8px;
    cursor: pointer;
  }
}
</style>
