<template>
  <div class="page-item section">

    <div class="info-title">Info</div>
  
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
  background-color: #101010;
  color: white;
  /*background-color: #4fc08d !important;*/
  .info-title{
    color: white;
    font-weight: bold;
    font-size: 60px;
    padding-top: 8vh;
    margin-bottom: 4vh;
  }
  img{
    border-radius: 100%;
    height: 200px;
    width: 200px;
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
    margin-top: 64px;
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


