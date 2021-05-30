<template>
  <div class="top">
    <div class="buyerinfo">
      <h1>Información de comprador</h1>
      <v-divider></v-divider>

      <div v-if="buyerLoaded" class="scrollable-list">
        <h1>Productos Comprados</h1>
        <table class="buyers-table">
          <tr class="buyers-tr">
            <th class="buyers-th">ID</th>
            <th class="buyers-th">Nombre</th>
            <th class="buyers-th">Precio</th>
          </tr>
          <tr v-for="product in buyerInfo.boughtProducts" :key="product.id" class="buyers-tr">
            <td class="buyers-td">{{product.id}}</td>
            <td class="buyers-td">{{product.name}}</td>
            <td class="buyers-td">{{product.price}}</td>
          </tr>
        </table>
        <v-divider></v-divider>
        <h1>Compradores con la misma dirección IP</h1>
        <table class="buyers-table">
          <tr class="buyers-tr">
            <th class="buyers-th">ID</th>
            <th class="buyers-th">Nombre</th>
            <th class="buyers-th">Edad</th>
          </tr>
          <tr v-for="buyer in buyerInfo.buyersSameIP" :key="buyer.id" class="buyers-tr">
            <td class="buyers-td">{{buyer.id}}</td>
            <td class="buyers-td">{{buyer.name}}</td>
            <td class="buyers-td">{{buyer.age}}</td>
          </tr>
        </table>
        <v-divider></v-divider>
        <h1>Productos Recomendados</h1>
        <table class="buyers-table">
          <tr class="buyers-tr">
            <th class="buyers-th">ID</th>
            <th class="buyers-th">Nombre</th>
            <th class="buyers-th">Precio</th>
          </tr>
          <tr v-for="product in buyerInfo.recommendedProducts" :key="product.id" class="buyers-tr">
            <td class="buyers-td">{{product.id}}</td>
            <td class="buyers-td">{{product.name}}</td>
            <td class="buyers-td">{{product.price}}</td>
          </tr>
        </table>
      </div>
      <div vi-else class="consult-buyer">
        <h3>A continuación digite el ID del comprador a consultar</h3>
        <v-text-field
        label="ID del comprador"
        v-model="buyerId"
        outlined
        clearable
        >
        </v-text-field>
        <div>
          <v-row aling="center" @click="fetchBuyer">
            <v-btn dark color="blue">Consultar Comprador</v-btn>
          </v-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BuyerInfo',
  data() {
    return{
      buyerLoaded: false,
      buyerInfo: null,
      buyerId: "",
    }
  },
  methods: {
    async fetchBuyer(){
      if(this.buyerId !=  ""){
        this.buyerLoaded = false
        fetch('http://localhost:5656/buyers'.concat('/', this.buyerId), {method: 'GET'}).then(response => response.json()).then(json => {
          //remove duplicated
          this.buyerInfo = json
          if(Object.keys(this.buyerInfo).length > 0){
            this.buyerLoaded = true;
          }
        })
      }
    },
  }
}
</script>

<style>

.top{
  display: flex;
  justify-content: center;
  align-items: center;
  height: auto;
}

.buyerinfo{
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

.consult-buyer{
  margin: 30px;
}

</style>
