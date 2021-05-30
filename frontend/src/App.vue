<template>
  <v-app>

    <b-navbar type="dark" variant="dark">
      <b-navbar-brand tag="h1" class="mb-0">Training Exercise</b-navbar-brand>
      <b-navbar-nav align='center'>
        <b-nav-item to="/">Sincronizar fecha</b-nav-item>
        <b-nav-item to="/buyers">Ver compradores</b-nav-item>
        <b-nav-item to="/buyerinfo">Información de compradores</b-nav-item>
      </b-navbar-nav>
    </b-navbar>
      <!-- -->

    <!-- Sizes your content based upon application components -->
    <v-main>

      <!-- Provides the application the proper gutter -->
      <v-container fluid>

        <!-- If using vue-router -->
        <router-view></router-view>
      </v-container>
    </v-main>

    <v-footer app>
      <!-- -->
    </v-footer>
  </v-app>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
