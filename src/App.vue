<template>
  <div id="app">

    <div class="left-nav">
      <div class="avator">
        <img src="@/assets/images/avatar.jpeg" alt="" title="点击唤醒人工智能" @click="aiShowChange">
        <div class="name">Alpaca Bi</div>
        <div class="info">
          <div class="item">
            <img src="@/assets/GitHub.png" alt="">
            <a class="text" href="https://github.com/AlpacaBi" target="_blank">AlpacaBi</a>
          </div>
          <div class="item">
            <img src="@/assets/wechat.png" alt="">
            <a class="text" href="mailto:let@alpaca.run">let@alpaca.run</a>
          </div>
          <div class="item">
            <img src="@/assets/mail.png" alt="">
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
          Create by <a href="https://vuejs.org" target="_blank" style="color:#4fc08d">Vue.js</a>
        </div>
      </div>
    </div>


    <div class="right-content"><router-view/></div>

    <div class="ai" v-show="aiShow">
      <div class="title">
        Alpaca AI
      </div>

    </div>
    
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
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 300px;
}
.left-nav{
  position: fixed;
  background-color: #545c64;
  height: 100%;
  width: 300px;
  color: white;
  .avator{
    text-align: center;
    img{
      margin-top: 1vh;
      width: 20vh;
      border-radius:100%;
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
.right-content{
  display: inline-block;
  margin-left: 300px
}

.ai{
  position: absolute;
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
    height: 45px;
    font-size: 25px;
    padding-top: 10px;
    padding-left: 20px;
  }
}

</style>
<script>
  /*import { Component, Prop, Vue } from 'vue-property-decorator';
  @Component
  export default class APP extends Vue {
    isCollapse: boolean = true;
    aiShow: boolean = false;

  }*/
  export default {
  
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      aiShowChange(){
        console.log(this.data)
        this.data.aiShow=!this.data.aiShow
      }
    }
  }
</script>