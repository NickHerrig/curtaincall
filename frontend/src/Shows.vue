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
      v-bind:name="show.name">
      v-bind:company="show.company">
      v-bind:description="show.description">
    </ShowTile>
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
          name: 'name1',
          company: 'company1',
          description: 'description1',
        },
        {
          id: 2,
          name: 'name2',
          company: 'company2',
          description: 'description2',
        },
        {
          id: 3,
          name: 'name3',
          company: 'company3',
          description: 'description3',
        },
      ],
    }
  },
}

</script>

<style>
</style>
