import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';
import { LoginComponent } from "./components/login.js";

let app = new Vue({
    el: '#app',
    components: {
        LoginComponent
    },
    data: {
    },
    created: function () {
    }
})