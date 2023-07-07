import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import './assets/main.css'

const axiosInstance = axios.create({
    baseURL: 'http://localhost:8080',
    withCredentials: true
})

const app = createApp(App)

app.config.globalProperties.$axios = axiosInstance

app.use(router)

app.mount('#app')
