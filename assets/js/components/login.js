import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export const Login = Vue.component('login', {
    data: function () {
        return {
            username: '',
            password: '',
        };
    },
    methods: {
      onSubmit: function (event) {
          event.preventDefault();
          event.stopPropagation();

          axios.post('/login', {
              username: this.username,
              password: this.password,
          }).then(response => {
              this.$root.authenticated = response.data.authenticated;
          }).catch(_ => {});
      }
    },
    template: `
        <form method="post" v-on:submit="onSubmit($event)">
            <img class="mb-4" src="/public/assets/img/creeper.png" alt="" width="57" height="57">
            <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

            <div class="alert alert-danger" role="alert">
            </div>

            <div class="form-floating">
                <input type="text" name="user" class="form-control" id="floatingInput" placeholder="Username" v-model="username">
                <label for="floatingInput">Username</label>
            </div>
            <div class="form-floating">
                <input type="password" name="password" class="form-control" id="floatingPassword" placeholder="Password" v-model="password">
                <label for="floatingPassword">Password</label>
            </div>
            <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    `
})
