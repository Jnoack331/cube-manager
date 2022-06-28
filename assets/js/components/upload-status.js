import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export function status(status, percent = 0) {
    return {
        status,
        percent,
    }
}

export const UploadStatus = Vue.component('upload-status', {
    data: function () {
      return {
          status: '',
          percent: 0,
          show: false,
          animateHide: false,
          timeout: null,
      }
    },
    mounted: function () {
        this.$root.$on('upload', upload => {
            if (upload) {
                this.error = false;
                this.show = true;
                this.animateHide = false;
                if (this.timeout !== null) {
                    clearTimeout(this.timeout);
                }
            } else {
                this.animateHide = true;
                this.timeout = setTimeout(_ => {
                    this.show = false;
                    this.animateHide = false;
                }, 2000)
            }
        })
        this.$root.$on('upload-status', (status) => {
            this.status = status.status;
            this.percent = status.percent;
        })
    },
    template: `
        <div class="toast-container position-fixed bottom-0 end-0 p-3">
            <div style="transition: opacity 500ms" class="toast align-items-center border-0" 
            v-bind:style="{ 'background': status === 'uploading' ? 'linear-gradient(to right, var(--bs-primary) '+percent+'%, white 0)': '' }" 
            v-bind:class="{ 'text-bg-primary': status === 'moving', 'text-bg-danger': status === 'error', 'text-bg-success': status === 'finished', 
            'show': show, 'showing': animateHide }" 
            role="alert" aria-live="assertive" aria-atomic="true">
              <div class="d-flex">
                <div class="toast-body">
                  <div v-if="status === 'uploading'">
                    Upload at {{ percent }}%.
                  </div>
                  <div v-if="status === 'moving'">
                    Moving Files.
                  </div>
                  <div v-if="status === 'finished'">
                    Upload finished.
                  </div>
                  <div v-if="status === 'error'">
                    Error during Upload.
                  </div>
                </div>
                <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button>
              </div>
            </div>
        </div>
    `
})
