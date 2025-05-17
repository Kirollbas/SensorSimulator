import { createApp } from 'vue'
import './style.css'
import { BASE_TO_NAME} from './constants/base';
import App from './App.vue'

const app = createApp(App)

app.config.globalProperties.$dictionaries = {
    BASE_TO_NAME
};

app.mount('#app')
