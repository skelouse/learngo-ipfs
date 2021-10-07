<template>
  <h1>Send file:</h1>
  <div class="container">
    <FileViewer />
  </div>

  <h1>Get file:</h1>
  <div class="container">
    <FileGetter />
  </div>
</template>

<script>
import FileViewer from "./components/FileSetter.vue";
import FileGetter from "./components/FileGetter.vue";

export default {
  name: "App",
  components: {
    FileViewer,
    FileGetter,
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