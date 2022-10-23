<script setup>
import TopicList from "./components/TopicList/TopicList.vue";
</script>

<template>
  <a-layout>
    <a-layout-header>
      <div class="logo" >
        <h1 class="white"> Pub/Sub UI </h1>
      </div>
    </a-layout-header>
    <a-layout-content class="content">
      <TopicList />
    </a-layout-content>
  </a-layout>

</template>

<style scoped>
  .content {
    padding: 20px 10%;
  }
  .white {
    color: white;
  }
</style>
