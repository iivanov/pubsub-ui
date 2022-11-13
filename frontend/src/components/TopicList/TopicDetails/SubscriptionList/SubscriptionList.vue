<template>
  <a-drawer
    v-model:visible="isShowGetMessageForm"
    width="40%"
    title="Messages"
    placement="right"
    @after-visible-change="afterVisibleChange"
  >
    <MessageList
      :subscriptionName="activeSubscriptionName"
      :isVisible="isGetMessageVisible"
    />
  </a-drawer>

  <a-card
    v-for="subscription in subscriptions"
    v-bind:key="subscription.Name"
    style="width: 100%; margin-top: 5px"
  >
    <h3>{{ subscription.Name }}</h3>
    <p>Ack deadline seconds: {{ subscription.AckDeadlineSeconds }}</p>
    <p v-if="subscription.PublishEndpoint">
      Publish endpoint: {{ subscription.PublishEndpoint }}
    </p>
    <p v-else>Consume mode</p>
    <p>Exactly once delivery: : {{ subscription.ExactlyOnceDelivery }}</p>
    <p>Enable message ordering: : {{ subscription.EnableMessageOrdering }}</p>

    <a-row justify="end" :gutter="16">
      <a-col>
        <a-button type="default" @click="onGetMessages(subscription)"
          >Get messages</a-button
        >
      </a-col>
      <a-col>
        <a-button type="danger" @click="onSubscriptionDelete(subscription)"
          >Delete subscription</a-button
        >
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import { ref } from "vue";
import MessageList from "@/components/TopicList/TopicDetails/SubscriptionList/MessageList.vue";
import { deleteApiData } from "@/composables/fetchData";

export default {
  name: "SubscriptionList",
  components: { MessageList },
  emits: ["subscriptionInfoUpdated"],
  props: {
    subscriptions: {
      type: Object,
      required: true,
    },
  },
  setup(props, { emit }) {
    const isShowGetMessageForm = ref(false);
    const activeSubscriptionName = ref("");
    const isGetMessageVisible = ref(false);

    const afterVisibleChange = (visible) => {
      isGetMessageVisible.value = visible;
    };

    const onGetMessages = (subscription) => {
      activeSubscriptionName.value = subscription.Name;
      isShowGetMessageForm.value = true;
    };

    const onSubscriptionDelete = (subscription) => {
      deleteApiData(
        `/api/topic/subscription/${subscription.Name}`,
        (response) => {
          console.log("Success:", response);
          emit("subscriptionInfoUpdated");
        }
      );
    };

    return {
      onSubscriptionDelete,
      afterVisibleChange,
      isShowGetMessageForm,
      activeSubscriptionName,
      onGetMessages,
      isGetMessageVisible,
    };
  },
};
</script>

<style scoped></style>
