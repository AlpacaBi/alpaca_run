import avatar from '@/assets/images/avatar.jpeg'
import location from '@/assets/images/location.svg'
import qualifications from '@/assets/images/qualifications.svg'
import marriage from '@/assets/images/marriage.svg'
import workstatus from '@/assets/images/workstatus.svg'

export const info = {
  name: '毕国康',
  avatar,
  introduce: 'Talk and code is cheap, show me your money',
  intros: [
    "一个来自广州的95后web全栈工程师",
    "热爱IT技术，从前端到deep learning都有所涉猎",
    "目前从事web开发工作，目前工作技术栈：Regularjs+Golang",
    "工作时间之外也鼓捣项目，业余技术栈：Vue.js+TypeScript+Golang",
    "喜欢唱、跳、rap、写代码"
  ],
  status: [
    {
      img: location,
      value: '广州'
    },
    {
      img: qualifications,
      value: '本科'
    },
    {
      img: marriage,
      value: '未婚'
    },
    {
      img: workstatus,
      value: '在职'
    }
  ]
}
export const skills = ['Vue', 'Vue-Router', 'Vuex', 'React', 'React-Native', 'NodeJs', 'JavaScript', 'MongoDB', 'Html5', 'Css3', 'Koa2', 'MpVue', 'Ubuntu', 'Webpack', 'electron', 'express', 'git']
export const profiles = [
  {
    title: 'vue商城',
    subTitle: '一个简单的PC-WEB电商代理商城',
    skills: 'vue vue-router vuex axios qiniu echart element-ui mavon-editor',
    text: '这个代理商城项目有完整的前后端页面，基本实现了商城的所有功能，用户登录注册，状态管理，商品分类展示搜索，用户购买、发货，后端页面也完全实现了网站整体设置，用户统计，购买统计，商品管理、用户管理、订单管理等的功能！',
    github: 'https://github.com/lyttonlee/learn'
  },
  {
    title: 'node-express-server',
    subTitle: '电商代理商城服务端',
    skills: 'node express mongoose mongodb express-promise-router ',
    text: '这是代理商城项目的服务端，express框架,数据库为mongodb3.4,使用async/await语法，使用express-promise-router统一捕获error，七牛云作为图片存储空间，开发中用到RoBo 3T,postman等工具，功能也完全响应代理商城客户端!',
    github: 'https://github.com/lyttonlee/learn'
  },
  {
    title: 'wx-idioms',
    subTitle: '微信小程序-成语接龙',
    skills: 'mpvue vuex wxapi',
    text: '一个基于mpvue的微信小程序,可以和AI玩成语接龙的游戏,采用面向对象的编程方式，虽然这个机器人Ai对象非常非常简单，但还是基本理解了面向对象和面向过程编程方式的一些区别',
    github: 'https://github.com/lyttonlee/wx-idioms'
  },
  {
    title: 'electron-mdEditor',
    subTitle: 'MarkDown文章编辑器',
    skills: 'electron nedb vue 七牛云SDK mavon-editor',
    text: '这是一个以mavon-editor为编辑器，nedb为本地数据库，七牛云作为图片存储空间，以Electron包装可以运行在各个桌面平台的应用程序，方便自己保存md语法的文章，也可以当成个人专属图床工具',
    github: 'https://github.com/lyttonlee/md-editor'
  },
  {
    title: 'fe-if',
    subTitle: '火焰纹章if升级模拟器',
    skills: 'vue element-ui',
    text: '这是一个用来模拟火焰纹章IF人物升级的小应用，功能实现也很简单就是遍历生成随机数与人物属性成长率相比较，最终得出每一个人物成长期望值。',
    github: 'https://github.com/lyttonlee/if'
  },
  {
    title: 'vue-todos',
    subTitle: '一个Todo List小应用',
    skills: 'vue TypeScript vue-router vuex axios apiCloud vue-chartJs',
    text: '诚实的讲这个todo和网上遍地的todo demo没什么区别，功能还是一样的功能，不过这算是自己尝试TypeScript的开始，也是不用UI框架的测试，也是自己学着造轮子的起点',
    github: 'https://github.com/lyttonlee/vue-todo'
  },
  {
    title: 'react-demo',
    subTitle: '一个简单的react使用demo',
    skills: 'react react-router redux antd-mobile',
    text: '这是一个看了react文档，尝试如何使用react的小demo，包含简单的react使用方法，功能只有一个小的评论列表包含添加和删除，一个数字加减的简单运算',
    github: 'https://github.com/lyttonlee/react-demo'
  },
  {
    title: 'react-native-demo',
    subTitle: '一个简单的react-native使用demo',
    skills: 'react-native redux React-Navgation ',
    text: '主要学会了使用react-native的基本组件，react-navigation的使用，对react-native的基本环境搭建及运行有了基础理解',
    github: 'https://github.com/lyttonlee/react-demo'
  },
  {
    title: '原生JS实现五子棋',
    subTitle: '简单的canvas实现五子棋',
    skills: 'JavaScript',
    text: '了解了canvas的基本使用,对DOM对象的操作有了进一步的理解',
    github: 'https://github.com/lyttonlee/wuziqi'
  }
]
export const career = [
  {
    time: '2010-2013',
    text: '从事工业产品设计，主要对产品的外形、拼装等的合理设计，对多种软件如PS、CAD、ProE等使用较多！'
  },
  {
    time: '2013-2016',
    text: '从事于矿业公司业务部门，主要是对矿山机械的采购、运行、维护等工作.自2014年起开始自学编程知识Html、JavaScript、CSS,并且使用WordPress框架为公司搭建了自己的站点'
  },
  {
    time: '2017-2018',
    text: '从事前端开发,主要使用Vue技术栈开发了电子商城网站,包括注册,商品列表,新闻,代理人员记录以及后台权限管理等模块,使用NodeJs开发了对应的服务端!'
  }
]
export const contact = {
  experiences: ['熟悉JavaScript,Html5,Css,Vue', '理解微信小程序,MpVue,NodeJs,Webpack,Git', '基本了解React,Electron,Ubuntu,TypeScript'],
  evaluation: '高效的自学能力,具备独立分析解决问题能力强烈的自我驱动力,爱好知识社区',
  // hobby: [],
  email: 'lzr3278@163.com',
  tel: '15181589155',
  socials: [
    {
      url: 'https://github.com/lyttonlee',
      icon: 'icon-GitHub'
    },
    {
      url: 'https://segmentfault.com/u/siwuweihesi',
      icon: 'icon-sf'
    },
    {
      url: 'https://www.zhihu.com/people/li-zai-rong-27/activities',
      icon: 'icon-zhihu'
    },
    {
      url: 'https://blog.csdn.net/weixin_43760383',
      icon: 'icon-csdn'
    }
  ]
}
