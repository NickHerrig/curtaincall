<template>

<div class="show-tile">
  <img :src="require('@/assets/'+logo)" alt="No logo" style="width:100%">
  <div class="container">
    <h3><b>{{ name }}</b></h3>
    <h4><b>{{ company }}</b></h4>
    <h4><b>{{ logo }}</b></h4>
  </div>
</div>

</template>

<script>
export default {
  name: 'ShowTile',
  props: {
    name: String,
    company: String,
    logo: String
  },
}
</script>

<style scoped>
.show-tile {
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
  transition: 0.3s;
}

.show-tile:hover {
  box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
}

.container {
  padding: 2px 16px;
}
</style>
