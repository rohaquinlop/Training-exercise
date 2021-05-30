<template>
  <div class="buyers">
    <h1>Compradores</h1>

    <div v-if="buyers">
      <ul>
      <li v-for="buyer in buyers" :key="buyer.id">
        {{buyer.id}} {{buyer.name}}
      </li>
    </ul>
    </div>
    <div vi-else>
      <p>No se han sincronizado los datos</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return{
      buyersLoaded: false,
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
    fetchBuyers() {
      this.buyers = null
      fetch('http://localhost:5656/buyers', {method: 'GET'}).then(response => response.json()).then(json => {
        //remove duplicated
        this.buyers = json.filter( (json, id, self) =>
          id === self.findIndex((t) => (t.id === json.id))
        );
        if(this.buyers.length > 0){
          this.buyersLoaded = true;
        }
      })
    },
  }
}
</script>