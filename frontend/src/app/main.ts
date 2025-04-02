import { createApp, reactive } from 'vue';
import App from './App.vue';
import { router } from '@/router';
import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import PanelMenu from 'primevue/panelmenu';
import ScrollPanel from 'primevue/scrollpanel';
import '@mdi/font/css/materialdesignicons.css';

import './assets/style.css';

const app = createApp(App);
app.use(PrimeVue);
app.use(router);
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
