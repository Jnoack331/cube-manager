import Vue from 'https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.esm.browser.js';

export function notification(message, background = 'info') {
    return {
        message,
        background,
    }
}

export const Notification = Vue.component('notification', {
    data: function () {
        return {
            message: '',
            background: '',
            show: false,
            animateHide: false,
            timeout: null,
        }
    },
    mounted: function () {
        this.$root.$on('notification', notification => {
                this.error = false;
                this.show = true;
                this.message = notification.message
                this.background = notification.background
                this.animateHide = false;
                if (this.timeout !== null) {
                    clearTimeout(this.timeout);
                }

                this.timeout = setTimeout(_ => {
                    this.show = false;
                    this.animateHide = false;
                }, 2000)
        })
    },
    template: `
        <div class="toast-container position-fixed bottom-0 end-0 p-3">
            <div style="transition: opacity 500ms" class="toast align-items-center border-0" 
            v-bind:class="{'show': show, 'showing': animateHide, 'text-bg-success': background === 'success', 'text-bg-danger': background === 'danger', 'text-bg-info': background === 'info',  'text-bg-success': background === 'success', 'text-bg-warning': background === 'warning' }"
            role="alert" aria-live="assertive" aria-atomic="true">
              <div class="d-flex">
                <div class="toast-body">
                  <div>
                    {{ message }}
                  </div>
                </div>
                <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button>
              </div>
            </div>
        </div>
    `
})
