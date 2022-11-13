import { message } from "ant-design-vue";

const getDataFromApi = (url, onSuccess, onError) => {
  fetch(url)
    .then((response) => response.json())
    .then((data) => {
      if (!data.success) {
        throw new Error(data.errorMessage);
      }
      onSuccess(data.data);
    })
    .catch((err) => {
      console.error("Error:", err);
      message.error(err.message);
      if (onError) {
        onError(err);
      }
    });
};

const sendDataToApi = (url, formData, onSuccess, onError) => {
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then((response) => response.json())
    .then((data) => {
      if (!data.success) {
        throw new Error(data.errorMessage);
      }
      if (onSuccess) {
        onSuccess(data.data);
      }
    })
    .catch((err) => {
      console.error("Error:", err);
      message.error(err.message);
      if (onError) {
        onError(err);
      }
    });
};

const deleteApiData = (url, onSuccess, onError) => {
  fetch(url, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      if (!data.success) {
        throw new Error(data.errorMessage);
      }
      if (onSuccess) {
        onSuccess(data.data);
      }
    })
    .catch((err) => {
      console.error("Error:", err);
      message.error(err.message);
      if (onError) {
        onError(err);
      }
    });
};

export { getDataFromApi, sendDataToApi, deleteApiData };
