<template>
  <div class="about">
    <h1>This is an about page</h1>

    <div v-if="buyers">
      <ul>
      <li v-for="buyer in buyers" :key="buyer.id">
        {{buyer.id}} {{buyer.name}}
      </li>
    </ul>
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
        this.buyers = json;
        this.buyersLoaded = true;
        console.log(this.buyers)
      })
    },
  }
}
</script>