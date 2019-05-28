<template>
  <div id="app">
    <b-navbar toggleable="lg" type="dark" variant="primary" fixed="top">
      <b-navbar-brand href="#">
        Truora
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-item :to="'/'">
            Consultar
          </b-nav-item>
          <b-nav-item :to="'data'">
            Reporte de consultas
          </b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <b-container class="container-app">
      <router-view></router-view>
    </b-container>
  </div>
</template>

<script>
import SearchInfo from './components/SearchInfo.vue'

export default {
  name: 'app',
  components: {
    SearchInfo
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
  .container-app{
    padding-top:100px;
  }
</style>
