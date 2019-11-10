<template>
  <div class="new-articles">
    <el-carousel :interval="4000" type="card" height="250px" class="new-articles-carousel">
      <el-carousel-item v-for="item in newArticles" :key="item">
        <h3>{{ item }}</h3>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

@Component({
  components: {
    HelloWorld,
  },
})
export default class Home extends Vue {
  private newArticles: string[] = ['关于ES6的研究', 'TypeScript装饰器探究', '为什么我会使用Graphql', 'Golang日常总结'];
}
</script>

<style lang="scss" scoped>
  .new-articles{
    text-align: center;
    width: 100%;
    padding-top: 20px;
  }
  .new-articles-carousel{
    width: 90%;
    margin-left: 5%;
  }

  .el-carousel__item h3 {
    color: #475669;
    font-size: 35px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }

</style>


