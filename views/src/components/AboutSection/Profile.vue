<template>
  <div class="page-item section">
    <div class="profile-title">Profile</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';

@Component
export default class Profile extends Vue {
  @State private profile!: any;
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: #24292e !important;
  .profile-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
}
</style>
