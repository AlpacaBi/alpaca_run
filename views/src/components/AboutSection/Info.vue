<template>
    <div class="info">
      <div class="name">Alpaca Bi</div>
      <div class="intro">一个来自广州的95后全栈工程师</div>
      <div class="more">想了解更多？请在终端输入以下命令</div>
      <el-tooltip content="点击即可复制" placement="right" effect="dark">
        <div class="terminal" @click="copy">
          npx alpaca-bi
        </div>
      </el-tooltip>
      <div class="next" @click="nextPage">Next</div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Action } from 'vuex-class';
import { Notification } from 'element-ui';

@Component
export default class Info extends Vue {

  @Action private next!: (x: string | undefined | null ) => any;

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

  private nextPage() {
    const presentName: string | undefined | null = this.$route.name;
    this.next(presentName).then((nextPageName: string)  => {
      this.$router.push({name: nextPageName});
    });
  }
}
</script>

<style lang="scss" scoped>
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
    .next{
      display: none;
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
    .next {
      position: absolute;
      color: #ff9900;
      bottom: 20px;
      font-size:24px;
      text-decoration: underline;
    }
  }
}
</style>


