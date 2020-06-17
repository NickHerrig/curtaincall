<template>
<div>
  <h1>{{ name }}</h1>
  <h1>{{ company }}</h1>
  <h1>{{ description }}</h1>
</div>
</template>

<script>
export default {
  name: 'ShowTile',
  props: {
    name: String,
    company: String,
    description: String,
  },
}
</script>


<style scoped>

</style>
