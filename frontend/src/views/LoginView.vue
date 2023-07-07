<script>
import { notification } from "../../../assets/js/components/notification.js";

export default {
  name: "Login",
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

      this.$axios.post('/login', {
        username: this.username,
        password: this.password,
      }).then(response => {
        if (!response.data.authenticated) {
          this.$root.$emit('notification', notification('Invalid Credentials.', 'danger'));
        }

        if (response.data.authenticated) {
          this.$emit('login-succeeded')
          this.$router.push('/');
        }
      }).catch(_ => {
        this.$root.$emit('notification', notification('Error during Login.', 'danger'));
      });
    }
  }
}
</script>

<template>
  <div class="form-signin w-100 m-auto" style="height: 100%; display: flex; justify-content: center; align-items: center;">
    <form method="post" v-on:submit="onSubmit($event)" style="margin-bottom: 50%;">
      <img class="mb-4" src="src/assets/creeper.png" alt="" width="57" height="57">
      <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

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
  </div>
</template>

<style scoped>

</style>