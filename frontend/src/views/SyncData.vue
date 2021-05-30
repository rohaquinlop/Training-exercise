<template>
  <div class="top">
    <div class="syncdata">
      <h1>Sincronizar datos</h1>
      <v-menu
        ref="menu"
        v-model="menu"
        :close-on-content-click="false"
        :return-value.sync="date"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date"
            label="Selecciona la fecha a sincronizar"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="date"
          no-title
          scrollable
        >
          <v-spacer></v-spacer>
          <v-btn
            text
            color="primary"
            @click="menu = false"
          >
            Cancel
          </v-btn>
          <v-btn
            text
            color="primary"
            @click="$refs.menu.save(date)"
          >
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
      <div v-if="!disabled">
        <v-row aling="center" @click="loadData" :disabled="disabled">
          <v-btn dark color="blue">Sincronizar</v-btn>
        </v-row>
      </div>
      <div v-else class="loading">
        Cargando los datos
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'SyncData',
  data() {
    return{
      date: new Date().toISOString().substr(0, 10),
      menu: false,
      disabled: false,
      timeout: null,
    }
  },
  methods: {
    loadData() {
      fetch('http://localhost:5656/load'.concat('/',this.date), {method: 'POST', mode:'no-cors'})
      this.disabled = true

      this.timeout = setTimeout(() => {
        this.disabled = false
      }, 9000)
    },
  },
  beforeDestroy() {
    clearTimeout(this.timeout)
  }
}
</script>

<style>

.top{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
}

.syncdata{
  margin: auto;
  width: 50vw;
  padding: 10px;
  align-items: center;
}

.loading:after{
  overflow: hidden;
  display: inline-block;
  vertical-align: bottom;
  -webkit-animation: ellipsis steps(4, end) 900ms infinite;
  animation: ellipsis steps(4, end) 900ms infinite;
  content: "\2026";
  width: 0px;
}

@keyframes ellipsis{
  to{
    width: 20px;
  }
}

@-webkit-keyframes ellipsis{
  to {
    width: 20px;
  }
}

</style>
