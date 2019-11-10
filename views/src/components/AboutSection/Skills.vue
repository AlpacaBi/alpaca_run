<template>
  <div class="page-item section">
    <mu-container class="container">
      <transition-group name="flip-list" tag="div">
        <mu-chip v-for="(item, index) in skills" class="chip" :key="index" color="#b0bec5" text-color="rgb(56, 19, 59)">
          {{item}}
        </mu-chip>
      </transition-group>
    </mu-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class'

@Component
export default class Skills extends Vue {
  @State skills
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: rgb(56, 19, 59) !important;
  .container {
    width: 80%;
    .chip {
      margin: 10px;
    }
  }
}
</style>
