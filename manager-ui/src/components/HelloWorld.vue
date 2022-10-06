<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";
interface Inquiry {
  ID?: string;
  CreatedAt?: Date;
  UpdateAt?: Date;
  DeletedAt?: Date;
  assignee?: string;
  message?: string;
  deadline?: string;
}

defineProps<{ msg: string }>();

const inquiries = ref<Inquiry[]>();
const message = ref<string>("");
const assignee = ref<string>("");

const submitForm = () => {
  axios
    .post("http://localhost:8080/inquiries/", {
      shop_id: 1,
      message: message.value,
      assignee: assignee.value,
    })
    .then(() => {
      message.value = "";
      assignee.value = "";
      axios.get("http://localhost:8080/inquiries/").then((res) => {
        inquiries.value = res.data;
      });
    });
};

const deleteInquiry = (id: string) => {
  axios.delete(`http://localhost:8080/inquiries/${id}`).then(() => {
    axios.get("http://localhost:8080/inquiries/").then((res) => {
      inquiries.value = res.data;
    });
  });
};

onMounted(() => {
  axios.get("http://localhost:8080/inquiries/").then((res) => {
    inquiries.value = res.data;
  });
});
</script>

<template>
  <h1>{{ msg }}</h1>
  <div
    style="display: flex; flex-direction: column; width: 300px; margin: auto"
  >
    <p>Form</p>
    <label for="message">Message</label>
    <input name="message" v-model="message" />

    <label for="assignee">assignee</label>
    <input name="assignee" v-model="assignee" />

    <button @click="submitForm">Submit</button>
  </div>
  <div class="inquiry-container">
    <div class="inquiry-block" v-for="inquiry in inquiries" :key="inquiry.ID">
      <p>{{ inquiry.message }}</p>
      <p>{{ inquiry.CreatedAt }}</p>
      <p>{{ inquiry.assignee || "No assignee" }}</p>
      <button
        @click="deleteInquiry(inquiry.ID as string)"
        class="delete-button"
      >
        Delete
      </button>
    </div>
  </div>
</template>

<style scoped>
.delete-button {
  background-color: #ffcdd2;
  color: white;
}

.read-the-docs {
  color: #888;
}

.inquiry-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.inquiry-block {
  margin: 20px;
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  color: #242424;
  transition: all ease-in-out 0.3s;
  cursor: pointer;
  user-select: none;
}

.inquiry-block:hover {
  transform: scale(1.1);
}
</style>
