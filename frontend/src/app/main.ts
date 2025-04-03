import { createApp, reactive } from 'vue';
import App from './App.vue';
import router from '@/router';
import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import PanelMenu from 'primevue/panelmenu';
import ScrollPanel from 'primevue/scrollpanel';
import '@mdi/font/css/materialdesignicons.css';
import 'es6-promise/auto';
import './assets/style.css';
import store from '@/store';

const app = createApp(App);
app.use(PrimeVue);
app.use(router);
app.use(store);
app.component('Button', Button);
app.component('InputText', InputText);
app.component('PanelMenu', PanelMenu);
app.component('ScrollPanel', ScrollPanel);
app.config.globalProperties.$globalState = reactive({
    studyName: '',
    studyIcon: '',
    activeChannelMap: {},
});
app.mount('#app');
