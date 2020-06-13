<template>
  <div class="theater-tile">
    <div class="theater-info">
      <h2 class="theater-name">{{ name }}</h2>
      <h3 class="theater-address">{{ address }}</h3>
      <p class="theater-description">{{ description }}</p>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'theater-tile',
    props: {
      name: String,
      address: String,
      description:String,
    },
  }
</script>

<style scoped>

</style>
