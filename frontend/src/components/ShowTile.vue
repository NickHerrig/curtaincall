<template>

<div class="show-tile">
  <div class="show-content">
    <h1>{{ name }}</h1>
    <div class="show-company">{{ company }}</div>
    <p>{{ description }}</p>
  </div>
</div>

</template>


<script>

export default {
  name: 'ShowTile',
  props: {
    name: String,
    company: String,
    description: String,
    logo: String
  },
}

</script>


<style scoped>

.show-tile {
  flex: 1 1 50%;
  position: relative;
  border: 1px solid #efefef;
  border-radius: 32px;
  box-shadow: 2px 6px 25px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  margin: 15px 0;
}

.show-tile:hover {
  box-shadow: 2px 6px 30px rgba(0, 0, 0, 0.4);
}

.show-content {
  padding: 30px 30px 30px;
}

.show-tile h1 {
  color: #444444;
  font-weight: 700;
  font-size: 2rem;
  margin: 0;
}

.show-tile p {
  color: #666666;
  margin: 0 0 30px;
}

.show-company {
  color: #ea7a2f;
  font-weight: 700;
  font-size: 1.25rem;
  margin: 0 0 15px;
}

</style>
