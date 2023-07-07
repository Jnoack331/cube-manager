<script setup>
</script>

<template>

  <nav class="navbar navbar-expand-lg bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">
        <img src="src/assets/creeper.png" alt="" width="57" height="57">
      </a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <RouterLink class="nav-link" activeClass="active" to="/">
              Chat & Commands
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" activeClass="active" to="/files">
              Files
            </RouterLink>
          </li>
          <li class="nav-item">
            <button class="btn btn-outline-primary mx-auto">
              <i class="bi bi-arrow-clockwise"></i>
              Restart
            </button>
          </li>
        </ul>
        <button class="btn btn-outline-secondary">
          <i class="bi bi-box-arrow-right"></i>
          Logout
        </button>
      </div>
    </div>
  </nav>
</template>

<style scoped>

</style>
