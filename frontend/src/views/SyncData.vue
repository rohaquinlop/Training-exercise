<template>
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
      <v-row aling="center" @click="loadData">
        <v-btn dark color="blue">Sincronizar</v-btn>
      </v-row>
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
    }
  },
  methods: {
    loadData(){
      fetch('http://localhost:5656/load'.concat('/',this.date), {method: 'POST', mode:'no-cors'})
    }
  },
}
</script>
