import { createApp } from 'vue';
import { App } from '@/app';
import router from '@/router';

// UI 라이브러리
import PrimeVue from 'primevue/config';
import PanelMenu from 'primevue/panelmenu';
import ScrollPanel from 'primevue/scrollpanel';
import Tooltip from 'primevue/tooltip';
import Password from 'primevue/password';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import ToastService from 'primevue/toastservice';
import Toast from 'primevue/toast';

// 스타일
import '@mdi/font/css/materialdesignicons.css';
import 'es6-promise/auto';
import './assets/style.css';

// 스토어
import pinia from '@/store';

// 애플리케이션 생성 및 설정
const app = createApp(App);

// 플러그인 등록
app.use(pinia);
app.use(PrimeVue);
app.use(ToastService);
app.use(router);

app.directive('tooltip', Tooltip);

// 전역 컴포넌트 등록
app.component('Button', Button);
app.component('InputText', InputText);
app.component('PanelMenu', PanelMenu);
app.component('ScrollPanel', ScrollPanel);
app.component('Password', Password);
app.component('Toast', Toast);

// 애플리케이션 마운트
app.mount('#app');
