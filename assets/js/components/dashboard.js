import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export const Dashboard = Vue.component('dashboard', {
    data: function () {
      return {
          currentPath: "",
          fileList: [],
      };
    },
    methods: {
      changeDirectory: function (directory = this.currentPath) {
          axios.get('/filelist', {
              params: {
                  path: directory,
              },
          }).then(response => {
              this.currentPath = response.data.currentPath;
              this.fileList = response.data.filelist;
          }).catch(_ => {});
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

          let formData = new FormData();
          formData.append("file", event.dataTransfer.files[0]);
          axios.post('/upload', formData, {
              headers: {
                  "Content-Type": "multipart/form-data",
              },
              params: {
                  path: this.currentPath,
              },
              onUploadProgress: (event) => {
                  const percentage = Math.round((100 * event.loaded) / event.total);
                  console.log(percentage);
              },
          }).then(_ => {
              this.changeDirectory();
          });
      }
    },
    mounted: function () {
        this.changeDirectory(this.currentPath);
        document.addEventListener('drag', this.onDrag)
        document.addEventListener('dragstart', this.onDrag)
        document.addEventListener('dragend',this.onDrag)
        document.addEventListener('dragover', this.onDrag)
        document.addEventListener('dragenter', this.onDrag)
        document.addEventListener('dragleave', this.onDrag)
        document.addEventListener('drop', this.onDrop);
    },
    template: `
        <div class="container-md">
          <ul class="list-group" style="border-bottom-left-radius: unset;border-bottom-right-radius: unset;margin-top: 2rem;">
            <li class="list-group-item" style="text-align: left;">
                {{ currentPath }}
            </li>
          </ul>
          <ul class="list-group" id="file-list" style="margin-top: -1px;border-top-left-radius: unset;border-top-right-radius: unset;">
            <file-entry v-bind:file="{ IsDir: true, Name: '.. /', Path: '' }" @change-directory="goUpOneFolder">
            </file-entry>
            <file-entry v-for="file in fileList" v-bind:file="file" @change-directory="changeDirectory">
            </file-entry>
          </ul>
        </div>
    `
})
