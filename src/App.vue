<template>
  <div id="app">

 
    <el-menu  class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" router>
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
    <router-view/>
  </div>
</template>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 300px;
    height: 100vh;
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