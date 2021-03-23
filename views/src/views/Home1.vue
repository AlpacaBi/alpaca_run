<template>
  <div class="new-articles">
    <el-carousel :interval="4000" type="card" height="250px" class="new-articles-carousel">
      <el-carousel-item v-for="item in newArticles" :key="item">
        <h3>{{ item }}</h3>
      </el-carousel-item>
    </el-carousel>
    <div>{{num}}</div>
    <button @click="del">-</button>
    <button @click="add">+</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import { State, Mutation } from 'vuex-class';

@Component({
  components: {
    HelloWorld,
  },
})
export default class Home extends Vue {
  private newArticles: string[] = ['关于ES6的研究', 'TypeScript装饰器探究', '为什么我会使用Graphql', 'Golang日常总结'];
  @State private num!: number;
  @Mutation private del!: () => void;
  @Mutation private add!: () => void;
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

  .el-carousel__item--card{
    background-color: white;
    border-radius: 6px;
    
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
    -webkit-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
    -o-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 40px rgba(0, 0, 0, 0.1) inset;
  }

  .el-carousel__item--card.is-active:hover{
    /* 明亮的边框 */
    border: 1px solid grey;
    /* 闪烁的动画 */
    animation:flash 2s linear infinite;         
  }
  /* 添加不同透明度的投影 */
  @keyframes flash{
    0%{box-shadow:0 0 10px grey; }
    50%{box-shadow:0 0 10px white; }
    100%{box-shadow:0 0 10px grey; }
  }

</style>


