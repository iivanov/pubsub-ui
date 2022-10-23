<template>
    <a-form @finish="onSubmit" :model="formState">
      <a-form-item label="Topic name" name="name" :rules="[{ required: true, message: 'Please input topic name!' }]">
        <a-input v-model:value="formState.name" />
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" html-type="submit">Create</a-button>
      </a-form-item>
    </a-form>

</template>

<script>
import {ref} from "vue";
import {message} from "ant-design-vue";

export default {
  name: "TopicCreate",
  emits: ['created'],
  setup(props, {emit}) {
    const formState = ref({"name": ""});

    const onSubmit = () => {
      fetch('/api/topic', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formState.value),
      })
        .then((response) => response.json())
        .then((data) => {
          console.log('Success:', data);
          formState.value.name = "";
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