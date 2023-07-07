<script>
import { status } from "../../../assets/js/components/upload-status.js";
import { notification } from "../../../assets/js/components/notification.js";
import ListEntry from "../components/ListEntry.vue";

export default {
  name: "FileList",
  components: {ListEntry},
  data: function () {
    return {
      currentPath: "",
      fileList: [],
      shiftPressed: false,
      ctrlPressed: false,
      startElement: null,
    };
  },
  methods: {
    changeDirectory: function (directory = this.currentPath) {
      this.$axios.get('/filelist', {
        params: {
          path: directory,
        },
      }).then(response => {
        this.currentPath = response.data.currentPath;
        this.fileList = response.data.filelist;
      }).catch(_ => {});
    },
    selectionChanged: function (file) {
      if (this.shiftPressed && this.startElement) {
        let start = this.fileList.indexOf(this.startElement);
        let end = this.fileList.indexOf(file);

        if (end < start) {
          const tmp = end;
          end = start;
          start = tmp;
        }

        this.fileList = this.fileList.map(fileListFile => {
          fileListFile.Selected = false;
          return  fileListFile;
        });

        for (let index = start; index <= end; index++) {
          this.fileList[index].Selected = true;
        }
      } else if (this.ctrlPressed) {
        this.fileList = this.fileList.map(fileListFile => {
          if (fileListFile === file) {
            fileListFile.Selected = true;
            this.startElement = fileListFile;
          }
          return fileListFile;
        });
      } else {
        this.fileList = this.fileList.map(fileListFile => {
          fileListFile.Selected = file === fileListFile;
          if (fileListFile.Selected) {
            this.startElement = fileListFile;
          }
          return fileListFile;
        });
      }
    },
    goUpOneFolder: function () {
      const pathSegments = this.currentPath.split('/');
      pathSegments.pop();
      this.changeDirectory(pathSegments.join('/'));
    },
    onDrag: function(e) {
      e.preventDefault();
    },
    onDrop: function(event) {
      event.preventDefault();
      this.$root.$emit("upload", true);

      let formData = new FormData();
      formData.append("file", event.dataTransfer.files[0]);
      this.$axios.post('/upload', formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
        params: {
          path: this.currentPath,
        },
        onUploadProgress: (event) => {
          const percentage = Math.round((100 * event.loaded) / event.total);
          if (percentage === 100) {
            this.$root.$emit("upload-status", status('moving', percentage));
          } else {
            this.$root.$emit("upload-status", status('uploading', percentage));
          }
        },
      }).then(_ => {
        this.changeDirectory();
        this.$root.$emit("upload-status", status('finished'));
        setTimeout( _ => this.$root.$emit("upload", false), 3000);
      }).catch(_ => {
        this.$root.$emit("upload-status", status('error'));
        setTimeout(_ => this.$root.$emit("upload", false), 3000);
      });
    },
    onKeyup: function (event) {
      if (event.code === 'Delete') {
        this.onDelete();
      } else if (event.key === 'Shift') {
        this.shiftPressed = false;
      } else if (event.key === 'Control') {
        this.ctrlPressed = false;
      }
    },
    onKeydown: function (event) {
      if (event.key === 'Shift') {
        this.shiftPressed = true;
      } else if (event.key === 'Control') {
        this.ctrlPressed = true;
      }
    },
    onDelete: function () {
      const filesToDelete = this.fileList.filter(file => file.Selected);
      axios.all(filesToDelete.map(file => {
        return this.$axios.post('/delete', {
          path: file.Path
        });
      })).then(results => {
        this.$root.$emit('notification', notification('Deleted ' + results.length + ' Files.', 'success'));
        this.changeDirectory();
      });
    }
  },
  mounted: function () {
    this.changeDirectory(this.currentPath);
    document.body.addEventListener('drag', this.onDrag)
    document.body.addEventListener('dragstart', this.onDrag)
    document.body.addEventListener('dragend',this.onDrag)
    document.body.addEventListener('dragover', this.onDrag)
    document.body.addEventListener('dragenter', this.onDrag)
    document.body.addEventListener('dragleave', this.onDrag)
    document.body.addEventListener('drop', this.onDrop);
    window.addEventListener('keyup', this.onKeyup);
    window.addEventListener('keydown', this.onKeydown);
  }
}
</script>

<template>
  <div class="container-md">
    <div class="btn-group" role="group" style="margin-top: 1rem; display: flex">
      <button class="btn btn-outline-danger" v-bind:disabled="!startElement">
        <i class="bi bi-trash3"></i>
        Delete
      </button>
    </div>
    <ul class="list-group" style="border-bottom-left-radius: unset;border-bottom-right-radius: unset;margin-top: 1rem;">
      <li class="list-group-item" style="text-align: left;">
        {{ currentPath }}
      </li>
    </ul>
    <ul class="list-group" id="file-list" style="margin-top: -1px;border-top-left-radius: unset;border-top-right-radius: unset;">
      <ListEntry v-bind:file="{ IsDir: true, Name: '.. /', Path: '' }" @change-directory="goUpOneFolder">
      </ListEntry>
      <ListEntry v-for="file in fileList" v-bind:file="file" v-bind:selected="file.Selected" @change-directory="changeDirectory" @selection-changed="selectionChanged">
      </ListEntry>
    </ul>
  </div>
</template>

<style scoped>

</style>