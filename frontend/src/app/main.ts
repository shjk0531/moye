import { createApp } from 'vue';
import App from './App.vue';
import { router } from '@/router';
import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import PanelMenu from 'primevue/panelmenu';
import ScrollPanel from 'primevue/scrollpanel';
import 'primeicons/primeicons.css';

import './assets/style.css';

createApp(App)
    .use(PrimeVue)
    .use(router)
    .use(PrimeVue)
    .component('Button', Button)
    .component('InputText', InputText)
    .component('PanelMenu', PanelMenu)
    .component('ScrollPanel', ScrollPanel)
    .mount('#app');
