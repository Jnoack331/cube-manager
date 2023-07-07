<script>
export default {
  props: [
    'file',
    'selected',
  ],
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
      this.$emit('selection-changed', this.file);
    },
    onDoubleClick: function () {
      if (this.file.IsDir) {
        this.$emit('change-directory', this.file.Path);
      } else {
        this.$el.querySelector('a').click();
      }
    },
  },
}
</script>

<template>
  <div class="list-group-item d-flex justify-content-between list-group-item-action" v-bind:class="{ 'active': selected }" v-on:click="onClick()" v-on:dblclick="onDoubleClick()">
          <span>
             <i class="bi bi-folder" v-if="file.IsDir"></i>
             <i class="bi bi-file-earmark" v-if="!file.IsDir"></i>
            {{ file.Name }}
          </span>
    <a v-bind:href="downloadLink" class="btn btn-outline-secondary" v-bind:download="filename"  v-if="!file.IsDir" style="display: none;">
      <i class="bi bi-cloud-download"></i>
    </a>
  </div>
</template>

<style scoped>
</style>
