<template>
  <div id="user-list" class="view">
    <header>
      <h2>Rabbitmq</h2>
    </header>
      <div class="message-container">
        <p v-for="(msg, index) in list" :key="index">{{ msg }}</p>
      </div>
  </div>
</template>

<script>
import Vue from 'vue'
import Loader from './loader.vue'

export default {
  name: 'userlist',
  components: {
    loader: Loader,
  },
  beforeCreate() {
    console.log("beforeCreate")
  },
  created() {
    console.log("created")
  },
  beforeMount() {
    console.log("beforeMount")
    if (window.ws) {
      window.ws.close()
    }
  },
  mounted() {
    var self = this
    var url = 'ws://localhost:' + process.env.APP_MONITOR_PORT + "/echo"
    if (window.ws) return

    window.ws = new WebSocket(url)
    window.ws.onopen = evt => {
      console.log('open socket')
    }
    window.ws.onmessage = evt => {
      console.log('message', evt.data)
      self.list.push(evt.data)
      window.setTimeout(function() { window.refresher.updateList() },50)
    }
    window.ws.onclose = evt => {
      window.ws = null
      console.log('close socket')
    }
    window.ws.onerror = evt => {
      console.log('error', evt.data)
    }

  },
  data () {
    return {
      loading: true,
      list: []
    }
  },
  methods: {
    
  }
}
</script>

<style lang="scss">
.message-container {
  p {
    text-align: left;
    margin: 0;
    padding: 0 8px;

    &:nth-child(even) {
      background: #eee;
    }
  }
}
</style>