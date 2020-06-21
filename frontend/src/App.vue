<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
  export default {
    name: 'App',
  }
</script>

<style>

@import url("https://cdn.rawgit.com/jgthms/minireset.css/master/minireset.min.css");
@import url('https://fonts.googleapis.com/css2?family=Varela+Round&display=swap');

html {
  font-family: 'Varela Round', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

h1,
h2,
h3,
h4,
h5 {
  line-height: 1.2;
}

</style>
