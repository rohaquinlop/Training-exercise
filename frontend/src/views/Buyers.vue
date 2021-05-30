<template>
  <div class="top">
    <div class="buyers">
      <h1>Compradores</h1>
      <v-divider></v-divider>

      <div v-if="buyersloaded" class="scrollable-list">
        <table class="buyers-table">
          <tr class="buyers-tr">
            <th class="buyers-th">ID</th>
            <th class="buyers-th">Nombre</th>
            <th class="buyers-th">Edad</th>
          </tr>
          <tr v-for="buyer in buyers" :key="buyer.id" class="buyers-tr">
            <td class="buyers-td">{{buyer.id}}</td>
            <td class="buyers-td">{{buyer.name}}</td>
            <td class="buyers-td">{{buyer.age}}</td>
          </tr>
        </table>
      </div>
      <div v-else>
        <h2>No se han sincronizado fechas</h2>
        <a href="/">Sincronizar fecha</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return{
      buyersloaded: false,
      buyers: null,
    }
  },
  created () {
    this.fetchBuyers()
  },
  watch: {
    '$route': 'fetchBuyers'
  },
  methods: {
    async fetchBuyers(){
      fetch('http://localhost:5656/buyers', {method: 'GET'}).then(response => response.json()).then(json => {
        //remove duplicated
        this.buyers = json.filter( (json, id, self) =>
          id === self.findIndex((t) => (t.id === json.id))
        );
        if(this.buyers.length > 0){
          this.buyersloaded = true;
        }
      })
    },
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

.buyers{
  margin: auto;
  width: 50vw;
  padding: 10px;
  align-items: center;
}

.scrollable-list{
  margin: 4px, 4px;
  padding: 4px;
  width: 50vw;
  height: 50vh;
  overflow-x: hidden;
  overflow-y: auto;
  text-align: center;
}
.buyers-table{
  margin: auto;
  padding: 10px;
  align-items: center;
  border-collapse: collapse;
  width: 80%;
}

.buyers-tr, .buyers-th{
  border: 1px solid #dddddd;
  text-align: center;
  padding: 8px;
}

.buyers-tr:nth-child(even){
  background-color: #dddddd;
}

</style>
