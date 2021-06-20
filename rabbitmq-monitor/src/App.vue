<template>
  <div id="app">
    <header>
      <h1>App Test monitor</h1>
    </header>
    <user-list class="column a" />
    <message-list class="column b" />
  </div>
</template>

<script>
import Vue from 'vue'
import UserList from './components/UserList.vue'
import MessageList from './components/MessageList.vue'


export default {
  components: { UserList, MessageList },
  name: 'app',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<style lang="scss">
body {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0;
}

a {
  color: #42b983;
}

.view {
  background: #fff;
}


body {
  margin: 0;
  padding: 0 !important;
}

#app {
  display: inline-grid;
  display: -moz-inline-grid;
  display: -ms-inline-grid;
  width: 100%;
  min-height: 100%;
  grid-template-areas: "header header"
                       "cola colb";
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 120px 1fr;
  grid-gap: 8px;

  @media (max-width:780px) {
    grid-template-areas: "header"
                         "cola"
                         "colb";
    grid-template-columns: 1fr;
    grid-template-rows: 120px 1fr 1fr;
  }

  header {
    grid-area: header;
    //height: 120px;
    //max-height: 120px;
  }

  $headerHeight: 48px;

  .view {
    background: #fff;
    border: 1px solid #eee;
    border-radius: 4px;

    header {
      height: $headerHeight;
      display: flex;
      align-items: center;
      justify-content: center;
      h2 {
        margin: 0;
      }
   }
  }

  .column {

    .a {
      grid-area: cola;
    }
    .b {
      grid-area: colb;
    }

    
  }
}



</style>
