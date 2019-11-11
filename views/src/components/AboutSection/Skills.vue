<template>
  <div class="page-item section">
    skills
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
  /*background-color: rgb(56, 19, 59) !important;*/
  .container {
    width: 80%;
    .chip {
      margin: 10px;
    }
  }
}
</style>
