<template>
  <div id="user-list" class="view">
    <header>
      <h2>Usuários</h2>
    </header>
      <table>
        <thead>
          <tr>
            <th>Nome</th>
            <th>Altura <sub>m</sub></th>
            <th>Peso <sub>Kg</sub></th>
            <th>Sexo</th>
            <th>IMC</th>
          </tr>
        </thead>
        <tbody>
          <tr v-show="loading"><td colspan="4"><loader /></td></tr>
          <tr v-show="!loading" v-for="user in list">
            <td>{{ user.Nome }}</td>
            <td>{{ user.Altura }}</td>
            <td>{{ user.Peso }}</td>
            <td><span>{{ {0:'_', 1:'&#9794;', 2:'&#9792;', 9:'_'}[user.Sexo] }}</span></td>
            <td>{{ user.IMC }}</td>
          </tr>
        </tbody>
      </table>
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
  },
  mounted() {
    console.log("mounted")
    window.setTimeout( this.updateList, 50 )
    window.refresher = this
  },
  beforeUpdate() {
    console.log("beforeUpdate")
  },
  updated() {
    console.log("updated")
  },
  beforeDestroy() {
    console.log("beforeDestroy")
  },
  destroyed() {
    console.log("destroyed")
  },
  data () {
    return {
      loading: true,
      list: ['s']
    }
  },
  methods: {
    updateList: function() {
      const self = this
      self.list.length = 0
      self.loading = true
      var http = new XMLHttpRequest()
      http.onreadystatechange = function() {
        if (this.readyState == 4) {
          console.log(this)
          if (this.status == 200)  {
            JSON.parse(this.responseText).data.map( i => self.list.push(i))
            self.loading = false
          } else {

          }
        }
      }
      http.open('GET', `http://localhost:${process.env.APPLICATION_PORT}/webapi/v1/user`, true)
      http.send()
    }
  }
}
</script>

<style lang="scss">

table {
  width: 100%;
  // height: calc(100% - #{$headerHeight});

  sub {
    position: relative;
    top: -0.3rem;
  }

  thead {
    background: darkkhaki;
  }

  tbody {
    tr:nth-child(odd) {
      background:  #ccc;
    }
    td {
      border-right: 1px solid #eee;
      
      &:last-child {
        border-right: none;
      }

      &:first-child {
        width: 40%;
      }

      &:nth-child(4) {
        position: relative;

        span {
          position: absolute;
          top: -0.2rem;
          font-size: 140%;
          font-weight: bolder;
        }
      }
    }
  }
}
</style>