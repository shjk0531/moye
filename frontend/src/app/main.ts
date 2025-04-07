import { createApp, reactive } from 'vue';
import { App } from '@/app';
import router from '@/router';
import store from '@/store';

// UI 라이브러리
import PrimeVue from 'primevue/config';
import PanelMenu from 'primevue/panelmenu';
import ScrollPanel from 'primevue/scrollpanel';

// 공유 UI 컴포넌트
import { Button, InputText } from '@/shared/ui';

// 스타일
import '@mdi/font/css/materialdesignicons.css';
import 'es6-promise/auto';
import './assets/style.css';

// 애플리케이션 생성 및 설정
const app = createApp(App);

// 플러그인 등록
app.use(PrimeVue);
app.use(router);
app.use(store);

// 전역 컴포넌트 등록
app.component('Button', Button);
app.component('InputText', InputText);
app.component('PanelMenu', PanelMenu);
app.component('ScrollPanel', ScrollPanel);

// 전역 상태
app.config.globalProperties.$globalState = reactive({
    studyName: '',
    studyIcon: '',
    activeChannelMap: {},
});

// 애플리케이션 마운트
app.mount('#app');
