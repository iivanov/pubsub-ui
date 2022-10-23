<template>
  <a-row justify="center">
    <a-spin v-if="!isLoaded" size="large" />
  </a-row>

  <a-card v-for="message in messages" style="width: 100%; margin-top: 5px">
    <p>Message: {{message.Message}}</p>
    <p>Attributes: {{message.Attribute}}</p>
  </a-card>
</template>

<script>
import {onMounted, ref, watch} from "vue";

export default {
  name: "MessageList",
  props: {
    subscriptionName: {
      type: String,
      required: true
    },
    isVisible: {
      type: Boolean,
      required: true
    }
  },
  setup(props) {
    const messages = ref([]);
    const isLoaded = ref(false);

    const getMessages = async () => {
      isLoaded.value = false;
      messages.value = [];
      fetch(`/api/subscription/${props.subscriptionName}/messages`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
          .then(response => response.json())
          .then(data => {
            messages.value = data.messages;
            isLoaded.value = true;
          })
          .catch(error => {
            console.error('Error:', error);
            isLoaded.value = true;
          });
    }

    watch(() => props.subscriptionName, () => {
      getMessages();
    });
    watch(() => props.isVisible, () => {
      getMessages();
    });

    onMounted(() => {
      getMessages();
    })

    return {
      messages,
      isLoaded
    }
  }
}
</script>

<style scoped>

</style>