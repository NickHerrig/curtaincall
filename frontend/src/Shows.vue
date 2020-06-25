<template>

<div class="container">
  <div class="flex-container">

    <ShowTile 
      v-for="show in shows" 
      v-bind:key="show.id" 
      v-bind:name="show.name"
      v-bind:company="show.company"
      v-bind:description="show.description"
      v-bind:logo="show.logo"
    ></ShowTile>

  </div>
</div>

</template>

<script>

import ShowTile from '@/components/ShowTile.vue'

export default {
  name: 'Shows',
  components: {
    ShowTile,
  },
  data: function() {
    return {
      shows: [],
    }
  },
  methods: {
    fetchAllShows: async function() {

      let response = await fetch("http://localhost:8888/shows", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })

      let data = await response.json()
      if (data.length > 0) {
        this.shows = data
      } else {
          this.errorMessage = data.error
      }
      this.dataReady = true;
    },
  },
  beforeMount(){
    this.fetchAllShows()
  },
}

</script>

<style>
* {
  &::before,
  &::after {
    box-sizing: border-box;
  }
}

.container {
  max-width: 850px;
  margin: 0 auto;
  padding: 0 15px;
}

@media screen and (min-width: 800px) {
  .flex-container {
    display: flex;
  }
}
</style>
