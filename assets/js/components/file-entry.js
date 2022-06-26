import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export const FileEntry = Vue.component('file-entry', {
    props: ['file'],
    methods: {
      onClick: function () {
          if (this.file.IsDir) {
              console.log('emit');
              this.$emit('change-directory', this.file.Path);
          }
      },
      onDelete: function () {
          let formData = new FormData();
          formData.append("path", this.file.Path);

          axios.post('/delete', formData).then(_ => {
              this.$emit('change-directory', this.file.Path);
          });
      },
      onDownload: function () {

      }
    },
    template: `
       <div class="list-group-item d-flex justify-content-between" v-bind:class="{ 'list-group-item-action': file.IsDir }" v-on:click="onClick()">
          <span>
             <i class="bi bi-folder" v-if="file.IsDir"></i>
             <i class="bi bi-file-earmark" v-if="!file.IsDir"></i>
            {{ file.Name }}
          </span>
          
            <button class="btn btn-outline-secondary" v-on:click="onDelete()">
                  <i class="bi bi-trash3"></i>
            </button>
        </div>
    `
})
