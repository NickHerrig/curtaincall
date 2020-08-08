<template>

<div>
  <ShowTile
    v-for="show in shows" 
    v-bind:key="show.id" 
    v-bind:name="show.name"
    v-bind:company="show.company"
    v-bind:description="show.description"
    v-bind:logo="show.logo"
  ></ShowTile>
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

      let response = await fetch(process.env.VUE_APP_API_FQDN + "/shows", {
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
</style>
