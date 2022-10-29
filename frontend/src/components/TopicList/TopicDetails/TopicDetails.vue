<template>
  <a-modal v-model:visible="isShowCreateSubscription" title="Create a new subscription">
    <SubscriptionCreate :topicName="topic.Name" @created="onSubscriptionCreated"/>
    <template #footer></template>
  </a-modal>

  <a-modal v-model:visible="isShowPublishMessageForm" title="Publish message">
    <PublishMessageForm :topicName="topic.Name" @published="isShowPublishMessageForm=false"/>
    <template #footer></template>
  </a-modal>

  <a-row justify="end" :gutter="16">
    <a-col>
      <a-button type="primary" @click="showCreateSubscription">Create subscription</a-button>
    </a-col>
    <a-col>
      <a-button type="default" @click="isShowPublishMessageForm=true">Publish message</a-button>
    </a-col>
    <a-col>
      <a-button type="danger" @click="onTopicDelete">Delete topic</a-button>
    </a-col>
  </a-row>

  <a-row style="margin-top: 10px">
    <SubscriptionList :subscriptions="topic.Subscriptions" @subscriptionInfoUpdated="onSubscriptionInfoUpdated"/>
  </a-row>
</template>

<script>
import {ref} from "vue";
import SubscriptionCreate from "./SubscriptionCreate.vue";
import SubscriptionList from "./SubscriptionList/SubscriptionList.vue";
import PublishMessageForm from "@/components/TopicList/TopicDetails/PublishMessageForm.vue";
import {deleteApiData} from "../../../composables/fetchData";

export default {
  name: "TopicDetails",
  components: {PublishMessageForm, SubscriptionList, SubscriptionCreate},
  emits: ['topicInfoUpdated'],
  props: {
    topic: {
      type: Object,
      required: true
    }
  },
  setup(props, {emit}) {
    const isShowCreateSubscription = ref(false);
    const isShowPublishMessageForm = ref(false);

    const showCreateSubscription = () => {
      isShowCreateSubscription.value = true;
    };

    const onSubscriptionCreated = () => {
      isShowCreateSubscription.value = false;
      emit('topicInfoUpdated');
    };

    const onSubscriptionInfoUpdated = () => {
      emit('topicInfoUpdated');
    };

    const onTopicDelete = () => {
      deleteApiData(
          `/api/topic/${props.topic.Name}`,
          (data) => {
            console.log('Success:', data);
            emit('topicInfoUpdated');
          }
      );
    };

    return {
      onTopicDelete,
      showCreateSubscription,
      isShowCreateSubscription,
      isShowPublishMessageForm,
      onSubscriptionCreated,
      onSubscriptionInfoUpdated
    }
  }
}
</script>

<style scoped>

</style>