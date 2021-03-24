<template>
  <div class="page-item section">
    <div class="profile-title">GitHub Status</div>
    <div class="github-status">
      <div class="item">
        <img src="https://github-stat.alpaca.run/api?username=alpacabi&show_icons=true&title_color=f90&bg_color=000000&icon_color=ffffff&text_color=f90&hide_border=true#gitstat" alt="Github Stat">
      </div>
      <div class="item">
        <img src="https://github-stat.alpaca.run/api/top-langs?username=alpacabi&layout=compact&title_color=f90&bg_color=000000&icon_color=ffffff&text_color=f90&hide_border=true&card_width=445#gitstat" alt="Hot Lang">
      </div>
    </div>
    <div class="profile-title2">开源项目</div>
    <div class="github-repo">
      <div class="item" v-for="(item, i) in reposData" :key=i @click="skipTo(item.html_url)">
        <div class="title">{{item.name}}</div>
        <div class="info">{{item.description}}</div>
        <div class="data">
          <div class="star" title="star"><span class="icon">⭐</span>{{item.stargazers_count}}</div>
          <div class="fork" title="fork"><span class="icon">💾</span>{{item.forks}}</div>
          <div class="language" title="language"><span class="icon">🔧</span>{{item.language}}</div>
        </div>
      </div>
    </div>
    <div class="next" @click="nextPage">Next</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State, Action } from 'vuex-class';

@Component
export default class Profile extends Vue {
  @State private profile!: any;
  @State private reposData!: any;
  @Action private next!: (x: string | undefined | null ) => any;

  private skipTo(url: string) {
    window.open(url);
  }

  private nextPage() {
    const presentName: string | undefined | null = this.$route.name;
    this.next(presentName).then((nextPageName: string)  => {
      this.$router.push({name: nextPageName});
    });
  }
}
</script>

<style lang="scss" scoped>
@media screen and (min-width: 769px){
  .section {
    height: calc(100vh - 80px);
    width: 98vw;
    background-color: black !important;
    .profile-title{
      color: white;
      font-weight: bold;
      font-size: 32px;
      margin-top: 80px;
    }
    .github-status{
      display:flex;
      justify-content: center;
      align-items: center;
      .item{
        border: 1px dashed #f90;
        margin: 20px;
        img {
          width: 400px;
          height: 160px;
        }
      }
    }
    .profile-title2{
      color: white;
      font-weight: bold;
      font-size: 32px;
      margin-top: 15px;
      margin-bottom: 25px;
    }
    .github-repo{
      display: grid;
      grid-template-columns: 300px 300px 300px;
      grid-template-rows: 100px 100px 100px;
      justify-content: center;
      grid-gap: 20px 30px;
      .item{
        position: relative;
        border: 1px dashed #f90;
        .title{
          position: absolute;
          top: 8px;
          left: 10px;
          color: #f90;
          font-weight: bold;
          font-size: 18px;
        }
        .info{
          position: absolute;
          text-align: left;
          top: 35px;
          left: 10px;
          width: 280px;
          color: #f90;
          font-size: 10px;
        }
        .data{
          position: absolute;
          text-align: left;
          bottom: 5px;
          left: 10px;
          color: #f90;
          font-size: 14px;
          .star{ 
            display: inline-block;
            width: 60px;
            .icon {
              margin-right: 6px;
            }
          }
          .fork{
            display: inline-block;
            width: 60px;
            .icon {
              margin-right: 6px;
            }
          }
          .language{
            display: inline-block;
             .icon {
              margin-right: 6px;
            }
          }
        }
        &:hover{
          -webkit-transition:0.5s all;/* 适用Safari 和 Chrome兼容 */
          -moz-transition:0.5s all;/* 适用Firefox兼容 */
          -ms-transition:0.5s all;/* 适用IE9兼容 */
          -o-transition:0.5s all;/* Opera */
          transition: .5s;
          -webkit-transform:scale(1.1);/*在Safari 和 Chrome兼容下还原*/
          -moz-transform:scale(1.1);/*在Firefox兼容下还原*/
          -ms-transfrom:scale(1.1);/*在IE9兼容下还原*/
          -o-transform:scale(1.1);/*在opera兼容下还原*/
          transform:scale(1.1);/*适用于所有浏览兼容下还原*/
        }
      }
    }
    .next{
      display: none;
    }
  }
}

@media screen and (max-width: 767px){
  .section {
    height: 100vh;
    width: 100vw;
    background-color: black !important;
    .profile-title{
      color: white;
      font-weight: bold;
      font-size: 28px;
      margin-top: 80px;
    }
    .github-status{
      display:flex;
      flex-wrap: wrap;
      justify-content: center;
      align-items: center;
      .item{
        border: 1px dashed #f90;
        margin: 10px;
        img { 
          max-width: 90%;
          height: auto;
        }
      }
    }
    .profile-title2{
      color: white;
      font-weight: bold;
      font-size: 28px;
      margin-top: 10px;
    }
    .github-repo{
      display:flex;
      flex-wrap: wrap;
      justify-content: center;
      align-items: center;
      .item{
        border: 1px dashed #f90;
        margin: 10px;
        width: 100%;;
        height: 100px;
        position: relative;
        .title{
          position: absolute;
          top: 8px;
          left: 4%;
          color: #f90;
          font-weight: bold;
          font-size: 18px;
        }
        .info{
          position: absolute;
          text-align: left;
          top: 35px;
          left: 4%;
          width: 92%;
          color: #f90;
          font-size: 10px;
        }
        .data{
          position: absolute;
          text-align: left;
          bottom: 5px;
          left: 10px;
          color: #f90;
          font-size: 14px;
          .star{ 
            display: inline-block;
            width: 60px;
            .icon {
              margin-right: 6px;
            }
          }
          .fork{
            display: inline-block;
            width: 60px;
            .icon {
              margin-right: 6px;
            }
          }
          .language{
            display: inline-block;
             .icon {
              margin-right: 6px;
            }
          }
        }
      }
    }
  }
  .next {
    margin-top: 40px;
    margin-bottom: 40px;
    color: #ff9900;
    font-size:24px;
    text-decoration:underline;
  }
}
</style>
