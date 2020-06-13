<template>
  <div class="home">
    <img class="logo" alt="Curtain Call Logo" src="./assets/logo.png">
    <h1>Welcome to Curtain Call!</h1>
    <button v-on:click="viewShows" type="button">View Shows</button>
  </div>
</template>

<script>
  export default {
    name: 'Home',
	methods: {
		viewShows: function() {
			this.$router.push('shows')
		},
  }
  }
</script>

<style>

.home {
  text-align: center;
  color: #2c3e50;
  margin-top: 1em;
  padding-top: 1em;
}

.home h1 {
  font-size: 30px;
  padding-bottom: 1em;
}

.home button {
  line-height: 2.5rem;
  font-size: 1.5rem;
  padding: 10px 20px;

  background: white;
  border: 2px solid #ea7a2f;
  border-radius: 32px;
  color: #ea7a2f;
}
.home button:hover {
  border-color: white; 
}


</style>
