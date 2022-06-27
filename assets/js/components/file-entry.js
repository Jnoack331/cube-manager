import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';
import {notification} from "./notification.js";

export const FileEntry = Vue.component('file-entry', {
    props: ['file'],
    computed: {
        downloadLink() {
            return '/download?path=' + this.file.Path;
        },
        filename() {
            return this.file.Name;
        }
    },
    methods: {
      onClick: function () {
          if (this.file.IsDir) {
              console.log('emit');
              this.$emit('change-directory', this.file.Path);
          }
      },
      onDelete: function () {
          axios.post('/delete', {
              path: this.file.Path
          }).then(_ => {
              this.$root.$emit('notification', notification('Deleted ' + this.file.Name, 'success'));
              this.$emit('change-directory');
          });
      }
    },
    template: `
       <div class="list-group-item d-flex justify-content-between" v-bind:class="{ 'list-group-item-action': file.IsDir }" v-on:click="onClick()">
          <span>
             <i class="bi bi-folder" v-if="file.IsDir"></i>
             <i class="bi bi-file-earmark" v-if="!file.IsDir"></i>
            {{ file.Name }}
          </span>
          
          <div>
            <a v-bind:href="downloadLink" class="btn btn-outline-secondary" v-bind:download="filename"  v-if="!file.IsDir">
              <i class="bi bi-cloud-download"></i>
            </a>
            <button class="btn btn-outline-secondary" v-on:click="onDelete()">
              <i class="bi bi-trash3"></i>
            </button>
          </div>
        </div>
    `
})
