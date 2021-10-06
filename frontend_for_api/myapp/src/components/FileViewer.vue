

<script>
import { upload } from "../file-upload.service";
export default {
  name: "FileViewer",
  data() {
    return {
      resp: "",
    };
  
  },
  methods: {
    submitFile() {
      let formData = new FormData();
      formData.append("Data", this.file);
      formData.append("Password");
      upload(formData)
        .then((x) => {
          this.resp = x.data;
          this.currentStatus = x.currentStatus;
        })
        .catch((err) => {
          this.uploadError = err.response;
          this.currentStatus = err.currentStatus;
        });
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