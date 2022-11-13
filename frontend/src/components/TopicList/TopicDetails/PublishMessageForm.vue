<template>
  <a-form
    :model="formState"
    :label-col="{ span: 5 }"
    :wrapper-col="{ span: 16 }"
    @finish="onSubmit"
  >
    <a-form-item
      name="message"
      label="Message"
      :rules="[{ required: true, message: 'Please input message!' }]"
    >
      <a-textarea v-model:value="formState.message" />
    </a-form-item>
    <a-form-item
      name="attributes"
      label="Attributes"
      :rules="[{ required: true, message: 'Please input attributes!' }]"
    >
      <a-textarea v-model:value="formState.attributes" />
    </a-form-item>

    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" html-type="submit">Publish</a-button>
    </a-form-item>
  </a-form>
</template>

<script>
import { ref } from "vue";
import { sendDataToApi } from "@/composables/fetchData";

export default {
  name: "PublishMessageForm",
  props: {
    topicName: {
      type: String,
      required: true,
    },
  },
  emits: ["published"],
  setup(props, { emit }) {
    const formState = ref({
      message: "",
      attributes: "{}",
    });

    const onSubmit = () => {
      sendDataToApi(
        `/api/topic/${props.topicName}/message/publish`,
        formState.value,
        (data) => {
          console.log("Success:", data);
          emit("published");
        }
      );
    };

    return {
      formState,
      onSubmit,
    };
  },
};
</script>

<style scoped></style>
