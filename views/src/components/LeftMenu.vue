<template>
  <div class="left-nav">
      <div class="avator">
        <img src="@/assets/images/avatar.jpeg" alt="" title="点击唤醒人工智能" @click="changeAiShow">
        <div class="name">Alpaca Bi</div>
        <div class="info">
          <div class="item">
            <img :src="images.GitHub" alt="">
            <a class="text" href="https://github.com/AlpacaBi" target="_blank">AlpacaBi</a>
          </div>
          <div class="item">
            <img :src="images.mail" alt="">
            <a class="text" href="mailto:let@alpaca.run">let@alpaca.run</a>
          </div>
          <div class="item">
            <img :src="images.wechat" alt="">
            <a class="text">workbiguokang</a>
          </div>
        </div>
      </div>

      <el-menu  
        class="el-menu-vertical-demo" 
        @open="handleOpen" 
        @close="handleClose" 
        router 
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b">
        <el-menu-item index="/">
          <i class="el-icon-menu"></i>
          <span slot="title">首页</span>
        </el-menu-item>
        <el-menu-item index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span slot="title">博客</span>
          </template>
        </el-menu-item>
        <el-menu-item index="3">
          <i class="el-icon-document"></i>
          <span slot="title">微勃</span>
        </el-menu-item>
        <el-menu-item index="4">
          <i class="el-icon-setting"></i>
          <span slot="title">逼乎</span>
        </el-menu-item>
        <el-menu-item index="/about">
          <i class="el-icon-setting"></i>
          <span slot="title">关于</span>
        </el-menu-item>
      </el-menu>

      <div class="beian">
        <div>
          <a href="http://www.beian.miit.gov.cn" target="_blank">粤ICP备17017545号</a>
        </div>
        <div style="margin-top: 0.5vh;">
          Create by 
          <a href="https://vuejs.org" target="_blank" style="color:#4fc08d">Vue.js</a>
          +
          <a href="https://golang.org" target="_blank" style="color:#00ADD8">Golang</a>
        </div>
      </div>
    </div>
</template>

<script lang="ts">
import { Component,  Vue , Emit} from 'vue-property-decorator';
import { State, Mutation } from 'vuex-class';

@Component
export default class LeftMenu extends Vue {
  @State private images!: any;
  @Mutation private changeAiShow!: () => void;

  private handleOpen(key: any, keyPath: any) {
    console.log(key, keyPath);
  }
  private handleClose(key: any, keyPath: any) {
    console.log(key, keyPath);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 300px;
}
.left-nav{
  position: fixed;
  background-color: #545c64;
  height: 100%;
  width: 300px;
  color: white;
  z-index: 6;
  .el-menu {
    border-right: none;
  }
  .avator{
    text-align: center;
    img{
      margin-top: 1vh;
      width: 20vh;
      border-radius:100%;
      cursor: pointer;
    }
    .name{
      font-size: 4vh;
    }
    .info{
      text-align: left;
      margin-left: 80px;
      margin-top: 3vh;
      .item{
        margin-bottom: 1vh;
        .text{
          margin-left: 1.5vh;
          text-decoration:none;
          color: white;
          cursor: pointer;
        }
        .text:hover{
          text-decoration:underline;
          color: white
        }
        img{
          width: 2vh;
          margin-top: 0vh;
        }
      }
    }
  }
  .beian{
    position: absolute;
    bottom: 3vh;
    text-align: center;
    width: 300px;
    a{
      text-decoration:none;
      color: white;
      cursor: pointer;
    }
    a:hover{
      text-decoration:underline;
      color: white
    }
  }
}

</style>
