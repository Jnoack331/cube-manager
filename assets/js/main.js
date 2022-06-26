import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';
import { Login } from "./components/login.js";
import { Dashboard } from "./components/dashboard.js";
import { FileEntry } from "./components/file-entry.js";
import { DashboardHeader } from "./components/dashboard-header.js";

let app = new Vue({
    el: '#app',
    components: {
        Login,
        Dashboard,
        FileEntry,
        DashboardHeader,
    },
    data: {
        authenticated: null,
    },
    created: function () {
        axios.get('/authenticated')
            .then(response => {
                this.authenticated = response.data.authenticated;
            })
            .catch(error => {
                this.authenticated = false;
            });
    },
})