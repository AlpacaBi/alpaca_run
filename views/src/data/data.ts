import avatar from '@/assets/images/avatar.jpeg';

const cdnImageURL: string = 'https://cdn.alpaca.run/alpaca_blog%2Fimages%2F';

export const images = {
  // LeftMenu
  GitHub: cdnImageURL + 'GitHub.png',
  mail: cdnImageURL + 'mail.png',
  wechat: cdnImageURL + 'wechat.png',
  // AlpacaAI
  close: cdnImageURL + 'close.svg',
  sendtext: cdnImageURL + 'sendtext.svg',
  sendimg: cdnImageURL + 'sendimg.svg',
  // Info
  location: cdnImageURL + 'location.svg',
  qualifications: cdnImageURL + 'qualifications.svg',
  marriage: cdnImageURL + 'marriage.svg',
  workstatus: cdnImageURL + 'workstatus.svg',
  // Contant
  cphone: cdnImageURL + 'cphone.svg',
  cwechat: cdnImageURL + 'cwechat.svg',
  cgithub: cdnImageURL + 'cgithub.svg',
  cmail: cdnImageURL + 'cmail.svg',
  wpQRCode: cdnImageURL + 'wpQRCode.jpg',
};

export const repoData = [
  {
      name: 'live2d_ai',
      description: '基于live2d.js实现的动画小人ai，拥有聊天功能，还有图片识别功能，可以嵌入到网页里',
      stargazers_count: 36,
      html_url: 'https://github.com/AlpacaBi/live2d_ai',
      forks: 13,
      language: 'JavaScript',
  },
  {
      name: 'bgk_weixin_bot',
      description: '基于python3的微信聊天机器人，增加了图片识别功能',
      stargazers_count: 26,
      html_url: 'https://github.com/AlpacaBi/bgk_weixin_bot',
      forks: 13,
      language: 'Python',
  },
  {
      name: 'alpaca_run',
      description: '一个基于TypeScript+Golang+Vue+Graphql的个人网站',
      stargazers_count: 24,
      html_url: 'https://github.com/AlpacaBi/alpaca_run',
      forks: 0,
      language: 'TypeScript',
  },
  {
      name: '30-days-of-react-native-expo',
      description: '30 days of React Native Expo demos',
      stargazers_count: 15,
      html_url: 'https://github.com/AlpacaBi/30-days-of-react-native-expo',
      forks: 2,
      language: 'JavaScript',
  },
  {
      name: 'qrcode_terminal',
      description: 'QRCode Terminal For Deno',
      stargazers_count: 7,
      html_url: 'https://github.com/AlpacaBi/qrcode_terminal',
      forks: 0,
      language: 'JavaScript',
  },
  {
      name: 'bgk-wechat-cleaner',
      description: '可以检测到微信好友、微信群的违规发言和违规图片，还微信交流良好环境',
      stargazers_count: 7,
      html_url: 'https://github.com/AlpacaBi/bgk-wechat-cleaner',
      forks: 3,
      language: 'Python',
  },
];
