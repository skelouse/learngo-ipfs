<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label>File
        <input
          type="file"
          id="file"
          ref="file"
          v-on:change="handleFileUpload()"
          required
        />
      </label>
      <label>
        <label for="pass">Password (8 characters minimum):</label>
        <input type="password" id="pass" ref="pass" minlength="8" required />
      </label>
      <button @click="submitFile">Submit</button>
    </div>
    <div class="large-12 medium-12 small-12 cell">
  <div v-if="resp" class="response-listing">{{ resp }} </div>
</div>
  </div>
</template>

<script>
import { upload } from "./file-upload.service";
export default {
  name: "App",
  data() {
    return {
      resp: null,
    };
  },
  methods: {
    submitFile() {
      let formData = new FormData();
      formData.append("data", this.file);
      formData.append("password", this.$refs.pass.pass);
      upload(formData)
    },
    handleFileUpload() {
      this.file = this.$refs.file.files[0];
    },
  },
};
</script>

<style lang="scss">
.dropbox {
  outline: 2px dashed grey; /* the dash box */
  outline-offset: -10px;
  background: lightcyan;
  color: dimgray;
  padding: 10px 10px;
  min-height: 200px; /* minimum height */
  position: relative;
  cursor: pointer;
}

.input-file {
  opacity: 0; /* invisible but it's there! */
  width: 100%;
  height: 200px;
  position: absolute;
  cursor: pointer;
}

.dropbox:hover {
  background: lightblue; /* when mouse over to the drop zone, change color */
}

.dropbox p {
  font-size: 1.2em;
  text-align: center;
  padding: 50px 0;
}
</style>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
