<template>
  <div class="bg-gray-100 p-4">
    <input 
      v-model="query" 
      @keyup.enter="onSearch" 
      placeholder="Buscar..." 
      class="w-full py-2 px-4 border-2 border-gray-300 rounded-md focus:border-gray-500 focus:outline-none"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      query: ''
    };
  },
  methods: {
    onSearch() {
      this.$emit('search', this.query);
    }
  }
};
</script>
