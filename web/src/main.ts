import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import searchInput from '@/components/SearchInput/index.vue'


import '@arco-design/web-vue/dist/arco.css';

const app = createApp(App)

app.use(createPinia())
app.use(router)
// app.use(ArcoVue);
app.component('search-input', searchInput)
app.mount('#app')
