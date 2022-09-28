<script>
export default {
  name: "ServerConsole",
  data: function () {
    return {
      serverOutput: [],
      command: '',
    };
  },
  methods: {
    refreshServerOutput: function () {
      this.$axios.get('/server/output')
          .then(response => {
            const responseServerOuput = response.data.output.split("\n");
            responseServerOuput.pop();
            if (this.serverOutput.length !== responseServerOuput.length) {
              this.serverOutput = responseServerOuput;
            }
          }).catch(_ => {});

      this.$axios.get('/server/status')
          .then(response => {
            console.log(response.data.running);
          }).catch(_ => {});
    },
    sendCommand: function () {
      this.$axios.post('/server/command', {command: this.command})
          .then(response => {
            this.command = '';
          }).catch(_ => {});
    },
    onType: function (event) {
      if (event.code === 'Enter') {
        this.sendCommand();
      }
    },
  },
  updated: function () {
    const serverOutput = this.$el.querySelector('#server-output');
    serverOutput.scroll(0, serverOutput.scrollHeight)
  },
  mounted: function () {
    setInterval(_ => this.refreshServerOutput(), 1000);
  },
}
</script>

<template>
  <div class="container-md">
    <!-- @todo fix this garbage ass style usages -->
    <ul class="list-group" id="server-output" style="border-bottom-left-radius: unset;border-bottom-right-radius: unset;margin-top: 1rem;height: 80%;display: block;overflow-y: scroll; overflow-x: hidden; word-break: break-all;">
      <li class="list-group-item" style="text-align: left; padding: 0.3rem 1rem; user-select: text" v-for="line in serverOutput">
        {{ line }}
      </li>
    </ul>
    <ul class="list-group" style="border-top-left-radius: unset;border-top-right-radius: unset;margin-top: -1px;display: block;">
      <li class="list-group-item" style="text-align: left;">
        <input type="text" id="command" class="form-control" placeholder="type a command..." v-model="command" v-on:keypress="onType($event)">
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>