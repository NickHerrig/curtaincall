<template>
<div>
  <div v-if="shows.length <= 0">
    <p>Sorry there are no shows yet!</p>
  </div>
  <div v-else>
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
  data() {
    return {
      shows: [
        {
          id: 1,
          name: 'Hamilton',
          company: 'Lin Manuel',
          description: 'The musical centres on the life of Alexander Hamilton, who was orphaned and moved to New York in hope of a better life. While there, the smart up-start impressed with his hunger for revolution and reform, to take the United States away from British forces. The story sees Hamilton become George Washington’s right-hand-man, fall in love, and go on to become the first Secretary of the Treasury of the United States.',
          logo: 'hamilton.jpg',
        },
        {
          id: 2,
          name: 'Dear Evan Hansen',
          company: 'Creative Co',
          description: 'Dear Evan Hansen tells the story of a young man with social anxiety disorder who so yearns to make a connection with his peers that he fabricates a relationship with a deceased student to become closer to the boy’s family.',
          logo: 'dear-evan-hansen.jpg',
        },
      ],
    }
  },
}

</script>

<style>
</style>
