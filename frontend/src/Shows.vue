<template>
<div>
  <div>
    <h1>Popular Shows Tonight!</h1>
  </div>

  <div v-if="shows.length <= 0">
    <p>Sorry there are no shows yet!</p>
  </div>
  <div v-else>
    <ShowTile 
      v-for="show in shows" 
      v-bind:key="show.id" 
      v-bind:name="show.name"
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
  data() {
    return {
      shows: [
        {
          id: 1,
          name: 'Hamilton',
          company: 'Lin Manuel',
          description: 'A fantastic show',
          logo: 'hamilton.jpg',
        },
        {
          id: 2,
          name: 'Dear Evan Hansen',
          company: 'Creative Co',
          description: 'Another fantastic show',
          logo: 'hamilton.jpg',
        },
      ],
    }
  },
}

</script>

<style>
</style>
