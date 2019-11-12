<template>
  <div class="page-item section">
    <div class="skills-title">Skills</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';

@Component
export default class Skills extends Vue {
  @State private skills!: any;
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: #409EFF !important;
  .skills-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
  .container {
    width: 80%;
    .chip {
      margin: 10px;
    }
  }
}
</style>
