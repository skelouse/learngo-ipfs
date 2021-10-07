<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label for="cid">CID</label>
      <input type="input" id="cid" ref="cid" minlength="46" required />
      <label for="pass">Password (8 characters minimum):</label>
      <input type="password" id="pass" ref="pass" minlength="8" required />
      <button @click="getFile">Download</button>
    </div>
    <div class="large-12 medium-12 small-12 cell">
      <div v-if="resp" class="response-listing">{{ resp }}</div>
    </div>
  </div>
</template>

<script>
import { download } from "../ifps-connector";

function handleFileDownload(File) {
  var binaryData = [];
  binaryData.push(File); //My blob
  const url = URL.createObjectURL(
    new Blob([binaryData], {
      type: "text/plain",
    })
  );
  const link = document.createElement("a");
  link.href = url;
  link.setAttribute("download", "file.txt");
  document.body.appendChild(link);
  link.click();
  document.body.removeChild();
}

function decodeBase64(base64) {
  return atob(base64);
}

export default {
  name: "file-getter",
  data() {
    return {
    };
  },
  methods: {
    getFile() {
      var cid = this.$refs.cid.value;
      download(cid, this.$refs.pass.value).then((response) => {
        var File = decodeBase64(response.data);
        handleFileDownload(File);
      });
    },
  },
};
</script>
