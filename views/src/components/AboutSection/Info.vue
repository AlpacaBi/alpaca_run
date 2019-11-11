<template>
  <div class="page-item section">
  
    <img :src="info.avatar">
    
    <h4>{{info.name}}</h4>
    <h3>{{info.introduce}}</h3>
    <div class="papers">
      <template v-for="(item, index) in info.intros">
        <div :key="index" class="paper">
          <div class="title">{{item}}</div>
        </div>
      </template>
    </div>

    <div class="infos">
      <template v-for="(item, index) in info.status">
        <div :key="index" class="info">
          <div><img :src="item.img"></div>
          <div class="title">{{item.value}}</div>
        </div>
      </template>
    </div>

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';


@Component
export default class Info extends Vue {
  @State private info!: any;
}
</script>

<style lang="scss" scoped>
.section {
  height: 100vh;
  background-color: black;
  color: white;
  /*background-color: rgb(20, 73, 73) !important;*/
  img{
    border-radius: 100%;
    height: 200px;
    width: 200px;
    margin-top: 12vh;
    border: 1px solid white;
  }
  h4{
    font-size: 40px;
    margin-top: 2px;
    margin-bottom: 0px
  }
  h3{
    font-size: 22px;
    margin-top: 6px;
  }
  .papers{
    margin-top: 30px;
  }
  .paper{
    margin-top: 5px;
  }

  .infos{
    margin-top: 0
  }
  .info{
    display: inline-block;
    margin-left: 12px;
    margin-right: 12px;
    img{
      height: 50px;
      width: 50px;
      border: none;
      border-radius: 0;
    }
  }

}
</style>


