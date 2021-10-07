<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label>
        <input
          type="file"
          id="file"
          ref="file"
          v-on:change="handleFileUpload()"
          required
        />
      </label>
      <label for="pass">Password (8 characters minimum):</label>
      <input type="password" id="pass" ref="pass" minlength="8" required />
      <button @click="submitFile">Submit</button>
    </div>
    <div class="large-12 medium-12 small-12 cell">
      <div v-if="resp" class="response-listing">{{ resp }}</div>
    </div>
  </div>
</template>

<script>
import { upload } from "../ifps-connector";
function b64EncodeUnicode(str) {
  return btoa(str);
}
export default {
  name: "file-setter",
  data() {
    return {
      resp: null,
    };
  },
  methods: {
    submitFile() {
      let formData = new FormData();
      var byteArrayFile = b64EncodeUnicode(this.file);

      formData.append("file", byteArrayFile);
      formData.append("password", this.$refs.pass.value);
      this.resp = upload(formData)
        .then((response) => {
          this.resp = response["data"];
        })
        .catch((error) => {
          this.resp = error;
        });
    },
    handleFileUpload() {
      this.file = this.$refs.file.files[0];
      const reader = new FileReader();
      if (this.file.name.includes(".txt")) {
        reader.onload = (res) => {
          this.file = res.target.result;
        };
        reader.onerror = (err) => console.log(err);
        reader.readAsText(this.file);
      } else {
        reader.onload = (res) => {
          this.file = res.target.result;
        };
        reader.onerror = (err) => console.log(err);
        reader.readAsText(this.file);
      }
    },
  },
};
</script>
