import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';
import { Login } from "./components/login.js";
import { Dashboard } from "./components/dashboard.js";
import { FileEntry } from "./components/file-entry.js";
import { DashboardHeader } from "./components/dashboard-header.js";
import { Notification } from "./components/notification.js";
import { UploadStatus } from "./components/upload-status.js";

let app = new Vue({
    el: '#app',
    components: {
        Login,
        Dashboard,
        FileEntry,
        DashboardHeader,
        Notification,
        UploadStatus,
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