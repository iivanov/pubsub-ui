<template>
  <a-form
      :model="formState"
      :label-col="{ span: 10 }"
      :wrapper-col="{ span: 16 }"
      @finish="onSubmit"
  >
    <a-form-item name="name" label="Subscription name" :rules="[{ required: true, message: 'Please input subscription name!' }]">
      <a-input v-model:value="formState.name" />
    </a-form-item>
    <a-form-item name="publishEndpoint" label="Publish endpoint">
      <a-input v-model:value="formState.publishEndpoint" />
    </a-form-item>
    <a-form-item name="ackDeadline" label="Ack deadline seconds" :rules="[{ required: true, message: 'Please input ack deadline!' }]">
      <a-input-number v-model:value="formState.ackDeadline" />
    </a-form-item>

    <a-form-item name="enableMessageOrdering" label="Enable exactly once delivery">
      <a-switch v-model:checked="formState.enableMessageOrdering" />
    </a-form-item>
    <a-form-item name="enableExactlyOnceDelivery" label="Enable message ordering">
      <a-switch v-model:checked="formState.enableExactlyOnceDelivery" />
    </a-form-item>

    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" html-type="submit">Create</a-button>
    </a-form-item>
  </a-form>
</template>

<script>
import {ref} from "vue";

export default {
  name: "SubscriptionCreate",
  emits: ['created'],
  props: {
    topicName: {
      type: String,
      required: true
    }
  },
  setup(props, {emit}) {
    const formState = ref({
      "name": "",
      "publishEndpoint": "",
      "ackDeadline": 10,
      "enableMessageOrdering": false,
      "enableExactlyOnceDelivery": false
    });

    const onSubmit = () => {
      fetch(`/api/topic/${props.topicName}/subscription`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formState.value),
      })
        .then((response) => response.json())
        .then((data) => {
          console.log('Success:', data);
          emit('created');
        })
        .catch((err) => {
          console.error('Error:', err);
          message.error(err.message);
        });
    };

    return {
      formState,
      onSubmit
    }
  }
}
</script>

<style scoped>

</style>