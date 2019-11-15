<template>
    <el-dialog title="登录" width="500px" :before-close="closeLoginDialogShow" :visible.sync="loginDialogShow">
        <div>
            <span @click="changeLoginState(0)">登录</span>
            <span @click="changeLoginState(1)">注册</span>
        </div>
        <div v-if="loginstate==0">
            <input type="text" v-model="account"/>
            <input type="password" v-model="password"/>
            <button @click="login">登录</button>
        </div>
        <div v-else>
            <input type="text" v-model="account" />
            <input type="password" v-model="password"/>
            <input type="password" v-model="repassword"/>
            <button @click="register">注册</button>
        </div>
    </el-dialog>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State, Mutation } from 'vuex-class';

@Component
export default class LoginDialog extends Vue {
  @State private loginDialogShow!: boolean;
  @Mutation private closeLoginDialogShow!: () => void;

  private loginstate: number = 0;
  private account: string = '';
  private password: string = '';
  private repassword: string = '';

  private changeLoginState(n: number): void {
      this.loginstate = n;
  }

  private login(): void  {
    console.log('login');
  }

  private register(): void  {
    console.log('register');
  }

}
</script>

<style lang="scss" scoped>
.el-dialog{
    width: 500px !important;
}
</style>
