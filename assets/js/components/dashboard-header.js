import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export const DashboardHeader = Vue.component('dashboard-header', {
    methods: {
        onLogout: function () {
            axios.post('/logout').then(response => {
                window.location = '/login';
            }).catch(_ => {});
        },
        onRestart: function () {
            axios.post('/restart').catch(_ => {});
        }
    },
    template: `
        <nav class="navbar navbar-expand-lg bg-light">
           <div class="container-fluid">
              <a class="navbar-brand" href="/">
                <img src="/public/assets/img/creeper.png" alt="" width="57" height="57">
              </a>
              <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
              </button>
              <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                  <li class="nav-item">
                      <button class="btn btn-outline-primary mx-auto" v-on:click="onRestart()">
                        <i class="bi bi-arrow-clockwise"></i>
                        Restart
                      </button>
                  </li>
                </ul>
                  <button class="btn btn-outline-secondary" v-on:click="onLogout()">
                    <i class="bi bi-box-arrow-right"></i>
                    Logout
                  </button>
              </div>
            </div>
        </nav>
    `
})
