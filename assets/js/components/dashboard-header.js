import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export const DashboardHeader = Vue.component('dashboard-header', {
    mounted: function () {
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
                    <form class="d-flex" action="/restart" method="post" data-set-current-path>
                      <button class="btn btn-outline-primary mx-auto" type="submit">
                        <i class="bi bi-arrow-clockwise"></i>
                        Restart
                      </button>
                    </form>
                  </li>
                </ul>
                <form class="d-flex" action="/logout" method="post">
                  <button class="btn btn-outline-secondary mx-auto" type="submit">
                    <i class="bi bi-box-arrow-right"></i>
                    Logout
                  </button>
                </form>
              </div>
            </div>
        </nav>
    `
})
