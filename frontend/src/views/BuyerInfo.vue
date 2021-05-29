<template>
  <div class="buyerInfo">
    <h1>Informacion</h1>

    <div v-if="buyerLoaded">
      <!-- <ul>
        <li v-for="buyer in buyers" :key="buyer.id">
          {{buyer.id}} {{buyer.name}}
        </li>
      </ul> -->
    </div>
    <div vi-else>
      <p>Ingrese el ID del usuario</p>
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
    }
  },
  /*
  created () {
    this.fetchBuyers()
  },
  watch: {
    '$route': 'fetchBuyers'
  },
  methods: {
    fetchBuyers() {
      this.buyers = null
      fetch('http://localhost:5656/buyers', {method: 'GET'}).then(response => response.json()).then(json => {
        this.buyers = json;
        if(this.buyers.length > 0){
          this.buyersLoaded = true;
        }
      })
    },
  }
  */
}
</script>