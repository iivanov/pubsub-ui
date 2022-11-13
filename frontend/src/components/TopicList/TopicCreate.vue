<template>
  <a-form @finish="onSubmit" :model="formState">
    <a-form-item
      label="Topic name"
      name="name"
      :rules="[{ required: true, message: 'Please input topic name!' }]"
    >
      <a-input v-model:value="formState.name" />
    </a-form-item>

    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" html-type="submit">Create</a-button>
    </a-form-item>
  </a-form>
</template>

<script>
import { ref } from "vue";
import { sendDataToApi } from "@/composables/fetchData";

export default {
  name: "TopicCreate",
  emits: ["created"],
  setup(props, { emit }) {
    const formState = ref({ name: "" });

    const onSubmit = () => {
      sendDataToApi("/api/topic", formState.value, () => {
        formState.value.name = "";
        emit("created");
      });
    };

    return {
      formState,
      onSubmit,
    };
  },
};
</script>

<style scoped></style>
