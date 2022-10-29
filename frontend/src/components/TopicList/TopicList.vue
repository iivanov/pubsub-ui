<template>
  <a-modal v-model:visible="isShowCreateTopic" title="Create a new subscription">
    <TopicCreate @created="onTopicCreated"/>
    <template #footer></template>
  </a-modal>

  <div v-if="topics.length === 0">
    <a-empty>
      <a-button @click="showCreateTopic" type="primary">Create topic</a-button>
    </a-empty>
  </div>
  <div v-else>
    <a-row type="flex" justify="space-between">
      <a-col flex="auto">
        <a-input v-model:value="filter" placeholder="Filter..." />
      </a-col>
      <a-col flex="100px">
        <a-button type="primary" @click="showCreateTopic">Create topic</a-button>
      </a-col>
    </a-row>
    <a-collapse accordion>
      <a-collapse-panel v-for="topic in filteredTopics">
        <template #header>
          <a-col style="padding-right: 10px">
            <a-badge
                :count="topic.Subscriptions.length"
                :number-style="{
              backgroundColor: '#fff',
              color: '#999',
              boxShadow: '0 0 0 1px #d9d9d9 inset',
            }"
            />
          </a-col>
          <a-col>{{ topic.Name }}</a-col>
        </template>
        <TopicDetails :topic="topic" @topicInfoUpdated="loadTopics"/>
      </a-collapse-panel>
    </a-collapse>
  </div>
</template>

<script>
import {defineComponent, onMounted, ref, computed} from 'vue';
import TopicCreate from "./TopicCreate.vue";
import TopicDetails from "./TopicDetails/TopicDetails.vue";
import {getDataFromApi} from "../../composables/fetchData";


export default defineComponent({
  components: {TopicDetails, TopicCreate},
  setup() {
    const topics = ref([]);
    const filter = ref('');
    const filteredTopics = computed(() => {
      return topics.value.filter((topic) => topic.Name.includes(filter.value));
    });

    const loadTopics = () => {
      getDataFromApi('/api/topic', (data) => topics.value = data);
    };


    onMounted(() => {
      loadTopics();
    })

    const isShowCreateTopic = ref(false);

    const showCreateTopic = () => {
      isShowCreateTopic.value = true;
    };

    const onTopicCreated = () => {
      loadTopics();
      isShowCreateTopic.value = false;
    };

    return {
      topics,
      filteredTopics,
      filter,
      isShowCreateTopic,
      showCreateTopic,
      onTopicCreated,
      loadTopics,
    };
  },
});
</script>

<style scoped>

</style>