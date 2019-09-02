<template>
  <div id="app">

    <div class="left-nav">
      <div class="avator">
        <img src="@/assets/images/avatar.jpeg" alt="">
        <h1 class="name">Alpaca Bi</h1>
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
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span slot="title">博客</span>
          </template>
          <el-menu-item-group>
            <span slot="title">前端</span>
            <el-menu-item index="1-1">JavaScript</el-menu-item>
            <el-menu-item index="1-2">html</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="后端">
            <el-menu-item index="1-3"></el-menu-item>
          </el-menu-item-group>
          <el-submenu index="1-4">
            <span slot="title">选项4</span>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-submenu>
        </el-submenu>
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
    </div>
    <router-view/>
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
  background-color: #545c64;
  height: 100%;
  width: 300px;
  .avator{
  text-align: center;
  height: 40vh;
  img{
    margin-top: 2vh;
    width: 230px;
    border-radius:100%;
  }
  .name{
    color: white
  }
}

}

</style>
<script>
  export default {
    data() {
      return {
        isCollapse: true
      };
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>